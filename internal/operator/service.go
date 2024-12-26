package operator

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	operatorv1 "github.com/PicoTools/pico-shared/proto/gen/operator/v1"
	"github.com/PicoTools/pico-shared/shared"
	"github.com/PicoTools/pico/internal/constants"
	"github.com/PicoTools/pico/internal/ent"
	"github.com/PicoTools/pico/internal/ent/blobber"
	"github.com/PicoTools/pico/internal/ent/command"
	"github.com/PicoTools/pico/internal/ent/message"
	"github.com/PicoTools/pico/internal/ent/operator"
	"github.com/PicoTools/pico/internal/ent/task"
	picoErrors "github.com/PicoTools/pico/internal/errors"
	"github.com/PicoTools/pico/internal/events"
	"github.com/PicoTools/pico/internal/middleware/grpcauth"
	"github.com/PicoTools/pico/internal/pools"
	"github.com/PicoTools/pico/internal/utils"
	"github.com/PicoTools/pico/internal/version"
	errs "github.com/go-faster/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	operatorv1.UnimplementedOperatorServiceServer
	ctx context.Context
	db  *ent.Client
	lg  *zap.Logger
}

// Hello register operator's session and manages subscriptions on all topics
func (s *server) Hello(req *operatorv1.HelloRequest, stream operatorv1.OperatorService_HelloServer) error {
	ctx := stream.Context()
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("Hello").With(zap.String("username", username))

	// validate operator's client verrsion
	if strings.Compare(req.GetVersion(), version.Version()) != 0 {
		lg.Warn(picoErrors.VersionMismatched, zap.String("version", req.GetVersion()))
		return status.Error(codes.InvalidArgument, picoErrors.VersionMismatched)
	}

	// check if operator's session already exists
	if pools.Pool.Hello.Exists(username) {
		lg.Warn(picoErrors.OperatorAlreadyConnected)
		return status.Error(codes.AlreadyExists, picoErrors.OperatorAlreadyConnected)
	}

	// generate cookie for session
	cookie := utils.RandString(32)

	// save operator's session
	pools.Pool.Hello.Add(cookie, username, stream)

	defer func() {
		// remove operator's session
		pools.Pool.Hello.Remove(cookie)
		// disconnect operator from topics
		pools.Pool.DisconnectAll(s.ctx, cookie)

		// create message with logout event
		if c, err := s.db.Chat.
			Create().
			SetIsServer(true).
			SetMessage(username + " logged out").
			Save(s.ctx); err != nil {
			lg.Warn(picoErrors.SaveChatMessage, zap.Error(err))
		} else {
			// notify operators with chat update
			go pools.Pool.Chat.Send(&operatorv1.ChatResponse{
				CreatedAt: timestamppb.New(c.CreatedAt),
				Message:   c.Message,
				IsServer:  c.IsServer,
			})
		}
	}()

	// send server data to operator with new cookie value
	if err := stream.Send(&operatorv1.HelloResponse{
		Response: &operatorv1.HelloResponse_Handshake{
			Handshake: &operatorv1.HandshakeResponse{
				Username: username,
				Time:     timestamppb.Now(),
				Cookie: &operatorv1.SessionCookie{
					Value: cookie,
				},
			},
		},
	}); err != nil {
		return err
	}

	lg.Info(events.OperatorLoggedIn)

	// create new message with login event
	if c, err := s.db.Chat.
		Create().
		SetIsServer(true).
		SetMessage(username + " logged in").
		Save(ctx); err != nil {
		lg.Warn(picoErrors.SaveChatMessage, zap.Error(err))
	} else {
		// notify operators with chat update
		go pools.Pool.Chat.Send(&operatorv1.ChatResponse{
			CreatedAt: timestamppb.New(c.CreatedAt),
			Message:   c.Message,
			IsServer:  c.IsServer,
		})
	}

	// ticker to send empty message every N secods
	ticker := time.NewTicker(constants.GrpcOperatorHealthCheckTimeout)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := stream.Send(&operatorv1.HelloResponse{
				Response: &operatorv1.HelloResponse_Empty{},
			}); err != nil {
				lg.Warn(events.OperatorLoggedOutWithError, zap.Error(err))
				return status.Error(codes.Canceled, picoErrors.Send)
			}
		case <-ctx.Done():
			lg.Info(events.OperatorLoggedOut)
			return nil
		}
	}
}

// SubscribeListeners subscribes operator for gathering events associated with listeners
func (s *server) SubscribeListeners(req *operatorv1.SubscribeListenersRequest, stream operatorv1.OperatorService_SubscribeListenersServer) error {
	ctx := stream.Context()
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SubscribeListeners").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// check if operator already exists in subscription topic
	if pools.Pool.Listeners.Exists(username) {
		lg.Warn(picoErrors.OperatorAlreadyConnected, zap.String("username", username))
		return status.Error(codes.AlreadyExists, picoErrors.OperatorAlreadyConnected)
	}

	// save operator's session for subscription
	pools.Pool.Listeners.Add(cookie, username, stream)

	defer func() {
		// remove operator's session from subscription
		pools.Pool.Listeners.Remove(cookie)
	}()

	lg.Info(events.OperatorSubscribed)

	// query all listeners from database
	ls, err := s.db.Listener.Query().All(ctx)
	if err != nil {
		lg.Error(picoErrors.QueryListeners, zap.Error(err))
		return status.Error(codes.Internal, picoErrors.Internal)
	}
	// split list by chunks and send to operator
	for _, chunk := range utils.ChunkBy(ls, constants.MaxObjectChunks) {
		if err = stream.Send(&operatorv1.SubscribeListenersResponse{
			Response: &operatorv1.SubscribeListenersResponse_Listeners{
				Listeners: pools.ToListenersResponse(chunk),
			},
		}); err != nil {
			lg.Error(picoErrors.SendListener, zap.Error(err))
			return status.Error(codes.Internal, picoErrors.Send)
		}
	}

	// get subscription's data
	val := pools.Pool.Listeners.Get(cookie)
	if val == nil {
		lg.Error(picoErrors.GetSubscriptionData)
		return status.Error(codes.Internal, picoErrors.GetSubscriptionData)
	}

	for {
		select {
		case <-val.IsDisconnect():
			lg.Info(events.OperatorUnsubscribedLoggedOut)
			return nil
		case err = <-val.Error():
			lg.Error(picoErrors.DuringSubscription, zap.Error(err))
			return status.Error(codes.Internal, picoErrors.DuringSubscription)
		case <-ctx.Done():
			lg.Info(events.OperatorUnsubscribed)
			return nil
		}
	}
}

// SubscribeListeners subscribes operator for gathering events associated with ants
func (s *server) SubscribeAnts(req *operatorv1.SubscribeAntsRequest, stream operatorv1.OperatorService_SubscribeAntsServer) error {
	ctx := stream.Context()
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SubscribeAnts").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// check if operator already exists in subscription topic
	if pools.Pool.Ants.Exists(username) {
		lg.Warn(picoErrors.OperatorAlreadyConnected)
		return status.Error(codes.AlreadyExists, picoErrors.OperatorAlreadyConnected)
	}

	// save operator's session for subscription
	pools.Pool.Ants.Add(cookie, username, stream)

	defer func() {
		// remove operator's session from subscription
		pools.Pool.Ants.Remove(cookie)
	}()

	lg.Info(events.OperatorSubscribed)

	// query all ants from database
	ants, err := s.db.Ant.Query().WithListener().All(ctx)
	if err != nil {
		lg.Error(picoErrors.QueryAnts, zap.Error(err))
		return status.Error(codes.Internal, picoErrors.Internal)
	}
	// split list by chunks and send to operator
	for _, chunk := range utils.ChunkBy(ants, constants.MaxObjectChunks) {
		if err = stream.Send(&operatorv1.SubscribeAntsResponse{
			Response: &operatorv1.SubscribeAntsResponse_Ants{
				Ants: pools.ToAntsResponse(chunk),
			},
		}); err != nil {
			lg.Error(picoErrors.SendListener, zap.Error(err))
			return status.Error(codes.Internal, picoErrors.Send)
		}
	}

	// get subscription's data
	val := pools.Pool.Ants.Get(cookie)
	if val == nil {
		lg.Error(picoErrors.GetSubscriptionData)
		return status.Error(codes.Internal, picoErrors.GetSubscriptionData)
	}

	for {
		select {
		case <-val.IsDisconnect():
			lg.Info(events.OperatorUnsubscribedLoggedOut)
			return nil
		case err = <-val.Error():
			lg.Error(picoErrors.DuringSubscription, zap.Error(err))
			return status.Error(codes.Internal, picoErrors.DuringSubscription)
		case <-ctx.Done():
			lg.Info(events.OperatorUnsubscribed)
			return nil
		}
	}
}

// SubscribeOperators subscribes operator for gathering events associated with operators
func (s *server) SubscribeOperators(req *operatorv1.SubscribeOperatorsRequest, stream operatorv1.OperatorService_SubscribeOperatorsServer) error {
	ctx := stream.Context()
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SubscribeOperators").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// check if operator already exists in subscription topic
	if pools.Pool.Operators.Exists(username) {
		lg.Warn(picoErrors.OperatorAlreadyConnected)
		return status.Error(codes.AlreadyExists, picoErrors.OperatorAlreadyConnected)
	}

	// save operator's session for subscription
	pools.Pool.Operators.Add(cookie, username, stream)

	defer func() {
		// remove operator's session from subscription
		pools.Pool.Operators.Remove(cookie)
	}()

	lg.Info(events.OperatorSubscribed)

	// query all operators from database
	operators, err := s.db.Operator.Query().All(ctx)
	if err != nil {
		lg.Error(picoErrors.QueryOperators, zap.Error(err))
		return status.Error(codes.Internal, picoErrors.Internal)
	}
	// split list by chunks and send to operator
	for _, chunk := range utils.ChunkBy(operators, constants.MaxObjectChunks) {
		if err = stream.Send(&operatorv1.SubscribeOperatorsResponse{
			Response: &operatorv1.SubscribeOperatorsResponse_Operators{
				Operators: pools.ToOperatorsResponse(chunk),
			},
		}); err != nil {
			lg.Error(picoErrors.SendListener, zap.Error(err))
			return status.Error(codes.Internal, picoErrors.Send)
		}
	}

	// get subscription's data
	val := pools.Pool.Operators.Get(cookie)
	if val == nil {
		lg.Error(picoErrors.GetSubscriptionData)
		return status.Error(codes.Internal, picoErrors.GetSubscriptionData)
	}

	for {
		select {
		case <-val.IsDisconnect():
			lg.Info(events.OperatorUnsubscribedLoggedOut)
			return nil
		case err = <-val.Error():
			lg.Error(picoErrors.DuringSubscription, zap.Error(err))
			return status.Error(codes.Internal, picoErrors.DuringSubscription)
		case <-ctx.Done():
			lg.Info(events.OperatorUnsubscribed)
			return nil
		}
	}
}

// SubscribeChat subscribes operator for gathering events associated with chat
func (s *server) SubscribeChat(req *operatorv1.SubscribeChatRequest, stream operatorv1.OperatorService_SubscribeChatServer) error {
	ctx := stream.Context()
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SubscribeChat").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// check if operator already exists in subscription topic
	if pools.Pool.Chat.Exists(username) {
		lg.Warn(picoErrors.OperatorAlreadyConnected)
		return status.Error(codes.AlreadyExists, picoErrors.OperatorAlreadyConnected)
	}

	// save operator's session for subscription
	pools.Pool.Chat.Add(cookie, username, stream)

	defer func() {
		// remove operator's session from subscription
		pools.Pool.Chat.Remove(cookie)
	}()

	lg.Info(events.OperatorSubscribed)

	// query all chat messages from database
	messages, err := s.db.Chat.
		Query().
		WithOperator().
		All(ctx)
	if err != nil {
		lg.Error(picoErrors.QueryChatMessages, zap.Error(err))
		return status.Error(codes.Internal, picoErrors.Internal)
	}
	// split list by chunks and send to operator
	for _, chunk := range utils.ChunkBy(messages, constants.MaxObjectChunks) {
		if err = stream.Send(&operatorv1.SubscribeChatResponse{
			Response: &operatorv1.SubscribeChatResponse_Messages{
				Messages: pools.ToChatMessagesResponse(chunk),
			},
		}); err != nil {
			lg.Error(picoErrors.SendListener, zap.Error(err))
			return status.Error(codes.Internal, picoErrors.Send)
		}
	}

	// get subscription's data
	val := pools.Pool.Chat.Get(cookie)
	if val == nil {
		lg.Error(picoErrors.GetSubscriptionData)
		return status.Error(codes.Internal, picoErrors.GetSubscriptionData)
	}

	for {
		select {
		case <-val.IsDisconnect():
			lg.Info(events.OperatorUnsubscribedLoggedOut)
			return nil
		case err = <-val.Error():
			lg.Error(picoErrors.DuringSubscription, zap.Error(err))
			return status.Error(codes.Internal, picoErrors.DuringSubscription)
		case <-ctx.Done():
			lg.Info(events.OperatorUnsubscribed)
			return nil
		}
	}
}

// SubscribeTasks subscribes operator for gathering events associated with credentials
func (s *server) SubscribeTasks(ss operatorv1.OperatorService_SubscribeTasksServer) error {
	ctx := ss.Context()
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SubscribeTasks").With(zap.String("username", username))

	// first request must be hello
	helloReq, err := ss.Recv()
	if err != nil {
		lg.Error(picoErrors.GetFirstHelloMsg, zap.Error(err))
		return status.Error(codes.Internal, picoErrors.GetFirstHelloMsg)
	}
	if helloReq.GetHello() == nil {
		lg.Error(picoErrors.NotHelloMsg)
		return status.Error(codes.InvalidArgument, picoErrors.NotHelloMsg)
	}

	// get cookie
	cookie := helloReq.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, helloReq.GetCookie().GetValue()) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// check if operator already exists in subscription topic
	if pools.Pool.Tasks.Exists(username) {
		lg.Warn(picoErrors.OperatorAlreadyConnected)
		return status.Error(codes.AlreadyExists, picoErrors.OperatorAlreadyConnected)
	}

	// save operator's session for subscription
	pools.Pool.Tasks.Add(cookie, username, ss)

	defer func() {
		// remove operator's session from subscription
		pools.Pool.Tasks.Remove(cookie)
	}()

	lg.Info(events.OperatorSubscribed)

	// get subscription's data
	val := pools.Pool.Tasks.Get(cookie)
	if val == nil {
		lg.Error(picoErrors.GetSubscriptionData)
		return status.Error(codes.Internal, picoErrors.GetSubscriptionData)
	}

	// query operator
	operator, err := s.db.Operator.
		Query().
		Where(operator.UsernameEQ(username)).
		Only(ctx)
	if err != nil {
		lg.Error(picoErrors.QueryOperator)
		return status.Error(codes.Internal, picoErrors.QueryOperator)
	}

	gw, subCtx := errgroup.WithContext(ctx)

	gw.Go(func() error {
		for {
			// recieve message
			msg, err := ss.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return nil
				}
				return err
			}

			// check if username and cookie are valid session identifiers
			if !pools.Pool.Hello.Validate(username, msg.GetCookie().GetValue()) {
				lg.Warn(picoErrors.InvalidSessionCookie)
				return errors.New(picoErrors.InvalidSessionCookie)
			}

			// add ant for operator's polling
			if msg.GetStart() != nil {
				// query ant
				ant, err := s.db.Ant.Get(ctx, msg.GetStart().GetId())
				if err != nil {
					if ent.IsNotFound(err) {
						lg.Warn(picoErrors.UnknownAnt, zap.Error(err))
						return status.Error(codes.InvalidArgument, picoErrors.UnknownAnt)
					}
					lg.Error(picoErrors.QueryAnt, zap.Error(err))
					return status.Error(codes.Internal, picoErrors.Internal)
				}
				// query all commands for ant
				commands, err := s.db.Command.
					Query().
					WithOperator().
					Where(command.AntIDEQ(ant.ID)).
					// skip invisible for operator commands
					Where(command.Not(command.And(
						command.AuthorIDNEQ(operator.ID),
						command.VisibleEQ(false),
					))).
					All(ctx)
				if err != nil {
					return errs.Wrap(err, picoErrors.QueryCommands)
				}
				for _, command := range commands {
					// query operator
					commandOperator, err := command.Edges.OperatorOrErr()
					if err != nil {
						lg.Warn(picoErrors.LoadOperator, zap.Int64("command-id", command.ID))
						continue
					}
					// send comand to operator
					if err = ss.Send(&operatorv1.SubscribeTasksResponse{
						Type: &operatorv1.SubscribeTasksResponse_Command{
							Command: &operatorv1.CommandResponse{
								Id:      command.ID,
								Aid:     ant.ID,
								Cmd:     command.Cmd,
								Author:  commandOperator.Username,
								Created: timestamppb.New(command.CreatedAt),
								Visible: command.Visible,
							},
						},
					}); err != nil {
						return errs.Wrap(err, picoErrors.SendCommand)
					}

					// query messages associated with command
					commandMessages, err := s.db.Message.
						Query().
						Where(message.CommandIDEQ(command.ID)).
						All(ctx)
					if err != nil {
						return errs.Wrap(err, picoErrors.QueryCommandMessages)
					}
					for _, commandMessage := range commandMessages {
						// send command's message to operator
						if err = ss.Send(&operatorv1.SubscribeTasksResponse{
							Type: &operatorv1.SubscribeTasksResponse_Message{
								Message: &operatorv1.MessageResponse{
									Id:      command.ID,
									Mid:     int64(commandMessage.ID),
									Aid:     ant.ID,
									Type:    uint32(commandMessage.Type),
									Message: commandMessage.Message,
									Created: timestamppb.New(commandMessage.CreatedAt),
								},
							},
						}); err != nil {
							return errs.Wrap(err, picoErrors.SendCommandMessage)
						}
					}

					// query tasks associated with command
					commandTasks, err := s.db.Task.
						Query().
						WithBlobberOutput().
						Where(task.CommandIDEQ(command.ID)).
						All(ctx)
					if err != nil {
						return errs.Wrap(err, picoErrors.QueryCommandTasks)
					}
					for _, commandTask := range commandTasks {
						taskRep := &operatorv1.TaskResponse{
							Id:        command.ID,
							Tid:       commandTask.ID,
							Aid:       ant.ID,
							Status:    uint32(commandTask.Status),
							OutputBig: commandTask.OutputBig,
							Created:   timestamppb.New(commandTask.CreatedAt),
						}
						// if output length > maximum -> add to protobuf message
						blob, err := commandTask.Edges.BlobberOutputOrErr()
						if err != nil {
							// if blob not found -> it doesn't exist for task
							if !ent.IsNotFound(err) {
								return errs.Wrap(err, picoErrors.QueryBlob)
							}
							taskRep.OutputLen = 0
						} else {
							taskRep.OutputLen = uint64(blob.Size)
							if !commandTask.OutputBig {
								taskRep.Output = wrapperspb.Bytes(blob.Blob)
							}
						}
						// send task to operator
						if err = ss.Send(&operatorv1.SubscribeTasksResponse{
							Type: &operatorv1.SubscribeTasksResponse_Task{
								Task: taskRep,
							},
						}); err != nil {
							return errs.Wrap(err, picoErrors.SendCommandTask)
						}
					}
				}
				// add ant to operator's polling
				pools.Pool.Tasks.AddAnt(cookie, msg.GetStart().GetId())
			}

			// remove ant from operator's polling
			if msg.GetStop() != nil {
				pools.Pool.Tasks.DeleteAnt(cookie, msg.GetStop().GetId())
			}
		}
	})

	for {
		select {
		case <-val.IsDisconnect():
			lg.Info(events.OperatorUnsubscribedLoggedOut)
			return nil
		case err = <-val.Error():
			lg.Error(picoErrors.DuringSubscription, zap.Error(err))
			return status.Error(codes.Internal, picoErrors.DuringSubscription)
		case <-subCtx.Done():
			if err = gw.Wait(); err != nil {
				lg.Error("error during receiving", zap.Error(err))
				return status.Error(codes.Internal, "something went wrong")
			}
			return nil
		case <-ctx.Done():
			lg.Info(events.OperatorUnsubscribed)
			return nil
		}
	}
}

// SubscribeCredentials subscribes operator for gathering events associated with credentials
func (s *server) SubscribeCredentials(req *operatorv1.SubscribeCredentialsRequest, stream operatorv1.OperatorService_SubscribeCredentialsServer) error {
	ctx := stream.Context()
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SubscribeCredentials").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// check if operator already exists in subscription topic
	if pools.Pool.Credentials.Exists(username) {
		lg.Warn(picoErrors.OperatorAlreadyConnected)
		return status.Error(codes.AlreadyExists, picoErrors.OperatorAlreadyConnected)
	}

	// save operator's session for subscription
	pools.Pool.Credentials.Add(cookie, username, stream)

	defer func() {
		// remove operator's session from subscription
		pools.Pool.Credentials.Remove(cookie)
	}()

	lg.Info(events.OperatorSubscribed)

	// query all credentials from database
	cs, err := s.db.Credential.
		Query().
		All(ctx)
	if err != nil {
		lg.Error(picoErrors.QueryCredentials, zap.Error(err))
		return status.Error(codes.Internal, picoErrors.Internal)
	}
	// split list by chunks and send to operator
	for _, chunk := range utils.ChunkBy(cs, constants.MaxObjectChunks) {
		if err = stream.Send(&operatorv1.SubscribeCredentialsResponse{
			Response: &operatorv1.SubscribeCredentialsResponse_Credentials{
				Credentials: pools.ToCredentialsResponse(chunk),
			},
		}); err != nil {
			lg.Error(picoErrors.SendListener, zap.Error(err))
			return status.Error(codes.Internal, picoErrors.Send)
		}
	}

	// query subscription's data
	val := pools.Pool.Credentials.Get(cookie)
	if val == nil {
		lg.Error(picoErrors.GetSubscriptionData)
		return status.Error(codes.Internal, picoErrors.GetSubscriptionData)
	}

	for {
		select {
		case <-val.IsDisconnect():
			lg.Info(events.OperatorUnsubscribedLoggedOut)
			return nil
		case err = <-val.Error():
			lg.Error(picoErrors.DuringSubscription, zap.Error(err))
			return status.Error(codes.Internal, picoErrors.DuringSubscription)
		case <-stream.Context().Done():
			lg.Info(events.OperatorUnsubscribed)
			return nil
		}
	}
}

// SetCredentialColor sets color to credential
func (s *server) SetCredentialColor(ctx context.Context, req *operatorv1.SetCredentialColorRequest) (*operatorv1.SetCredentialColorResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SetCredentialColor").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// query credential
	credential, err := s.db.Credential.Get(ctx, req.GetId())
	if err != nil {
		lg.Error(picoErrors.QueryCredential, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}
	// update color
	if credential, err = credential.Update().SetColor(req.GetColor().GetValue()).Save(ctx); err != nil {
		lg.Error(picoErrors.UpdateCredential, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// notify operators about credential's color update
	go pools.Pool.Credentials.Send(pools.ToCredentialColorResponse(credential))

	return &operatorv1.SetCredentialColorResponse{}, nil
}

// SetCredentialsColor sets color to list of credentials
func (s *server) SetCredentialsColor(ctx context.Context, req *operatorv1.SetCredentialsColorRequest) (*operatorv1.SetCredentialsColorResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SetCredentialsColor").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	for _, id := range req.GetIds() {
		// query credential
		credential, err := s.db.Credential.Get(ctx, id)
		if err != nil {
			lg.Error(picoErrors.QueryCredential, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}
		// update color
		if credential, err = credential.Update().SetColor(req.GetColor().GetValue()).Save(ctx); err != nil {
			lg.Error(picoErrors.UpdateCredential, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}

		// notify operators about credential's color update
		go pools.Pool.Credentials.Send(pools.ToCredentialColorResponse(credential))
	}

	return &operatorv1.SetCredentialsColorResponse{}, nil
}

// SetCredentialNote sets note to credential
func (s *server) SetCredentialNote(ctx context.Context, req *operatorv1.SetCredentialNoteRequest) (*operatorv1.SetCredentialNoteResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SetCredentialNote").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// query credential
	credential, err := s.db.Credential.Get(ctx, req.GetId())
	if err != nil {
		lg.Error(picoErrors.QueryCredential, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}
	// update note
	if credential, err = credential.Update().SetNote(req.GetNote()).Save(ctx); err != nil {
		lg.Error(picoErrors.UpdateCredential, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// notify operators about credential's note update
	go pools.Pool.Credentials.Send(pools.ToCredentialNoteResponse(credential))

	return &operatorv1.SetCredentialNoteResponse{}, nil
}

// SetCredentialNote sets note to list of credentials
func (s *server) SetCredentialsNote(ctx context.Context, req *operatorv1.SetCredentialsNoteRequest) (*operatorv1.SetCredentialsNoteResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SetCredentialsNote").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	for _, id := range req.GetIds() {
		// query credential
		credential, err := s.db.Credential.Get(ctx, id)
		if err != nil {
			lg.Error(picoErrors.QueryCredential, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}
		// update note
		if credential, err = credential.Update().SetNote(req.GetNote()).Save(ctx); err != nil {
			lg.Error(picoErrors.UpdateCredential, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}

		// notify operators about credential's note update
		go pools.Pool.Credentials.Send(pools.ToCredentialNoteResponse(credential))
	}

	return &operatorv1.SetCredentialsNoteResponse{}, nil
}

// SetListenerColor sets color to listener
func (s *server) SetListenerColor(ctx context.Context, req *operatorv1.SetListenerColorRequest) (*operatorv1.SetListenerColorResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SetListenerColor").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// query listener
	listener, err := s.db.Listener.Get(ctx, req.GetId())
	if err != nil {
		lg.Error(picoErrors.QueryListener, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}
	// update color
	if listener, err = listener.Update().SetColor(req.GetColor().GetValue()).Save(ctx); err != nil {
		lg.Error(picoErrors.UpdateListener, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// notify operators about listener's color update
	go pools.Pool.Listeners.Send(pools.ToListenerColorResponse(listener))

	return &operatorv1.SetListenerColorResponse{}, nil
}

// SetListenersColor sets color to list of listeners
func (s *server) SetListenersColor(ctx context.Context, req *operatorv1.SetListenersColorRequest) (*operatorv1.SetListenersColorResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SetListenersColor").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	for _, id := range req.GetIds() {
		// query listener
		listener, err := s.db.Listener.Get(ctx, id)
		if err != nil {
			lg.Error(picoErrors.QueryListener, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}
		// update color
		if listener, err = listener.Update().SetColor(req.GetColor().GetValue()).Save(ctx); err != nil {
			lg.Error(picoErrors.UpdateListener, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}

		// notify operators about listener's color update
		go pools.Pool.Listeners.Send(pools.ToListenerColorResponse(listener))
	}

	return &operatorv1.SetListenersColorResponse{}, nil
}

// SetListenerNote sets note to listener
func (s *server) SetListenerNote(ctx context.Context, req *operatorv1.SetListenerNoteRequest) (*operatorv1.SetListenerNoteResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SetListenerNote").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// query listener
	listener, err := s.db.Listener.Get(ctx, req.GetId())
	if err != nil {
		lg.Error(picoErrors.QueryListener, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}
	// update note
	if listener, err = listener.Update().SetNote(req.GetNote()).Save(ctx); err != nil {
		lg.Error(picoErrors.UpdateListener, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// notify operators about listener's note update
	go pools.Pool.Listeners.Send(pools.ToListenerNoteResponse(listener))

	return &operatorv1.SetListenerNoteResponse{}, nil
}

// SetListenerNote sets note to list of listeners
func (s *server) SetListenersNote(ctx context.Context, req *operatorv1.SetListenersNoteRequest) (*operatorv1.SetListenersNoteResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SetListenersNote").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	for _, id := range req.GetIds() {
		// query listener
		l, err := s.db.Listener.Get(ctx, id)
		if err != nil {
			lg.Error(picoErrors.QueryListener, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}
		// update note
		if l, err = l.Update().SetNote(req.GetNote()).Save(ctx); err != nil {
			lg.Error(picoErrors.UpdateListener, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}

		// notify operators about listener's note update
		go pools.Pool.Listeners.Send(pools.ToListenerNoteResponse(l))
	}

	return &operatorv1.SetListenersNoteResponse{}, nil
}

// SetAntColor sets color to ant
func (s *server) SetAntColor(ctx context.Context, req *operatorv1.SetAntColorRequest) (*operatorv1.SetAntColorResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SetAntColor").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// query ant
	ant, err := s.db.Ant.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			lg.Warn(picoErrors.UnknownAnt, zap.Error(err))
			return nil, status.Error(codes.InvalidArgument, picoErrors.UnknownAnt)
		}
		lg.Error(picoErrors.QueryAnt, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}
	// update color
	if ant, err = ant.Update().SetColor(req.GetColor().GetValue()).Save(ctx); err != nil {
		lg.Error(picoErrors.UpdateAnt, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// notify operators about ant's color update
	go pools.Pool.Ants.Send(pools.ToAntColorResponse(ant))

	return &operatorv1.SetAntColorResponse{}, nil
}

// SetAntsColor sets color to list of ants
func (s *server) SetAntsColor(ctx context.Context, req *operatorv1.SetAntsColorRequest) (*operatorv1.SetAntsColorResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SetAntsColor").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	for _, id := range req.GetIds() {
		// query ant
		ant, err := s.db.Ant.Get(ctx, id)
		if err != nil {
			if ent.IsNotFound(err) {
				lg.Warn(picoErrors.UnknownAnt, zap.Error(err))
				return nil, status.Error(codes.InvalidArgument, picoErrors.UnknownAnt)
			}
			lg.Error(picoErrors.QueryAnt, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}
		// update color
		if ant, err = ant.Update().SetColor(req.GetColor().GetValue()).Save(ctx); err != nil {
			lg.Error(picoErrors.UpdateAnt, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}

		// notify operators about ant's color update
		go pools.Pool.Ants.Send(pools.ToAntColorResponse(ant))
	}

	return &operatorv1.SetAntsColorResponse{}, nil
}

// SetAntNote set note to ant
func (s *server) SetAntNote(ctx context.Context, req *operatorv1.SetAntNoteRequest) (*operatorv1.SetAntNoteResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SetAntNote").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// query ant
	ant, err := s.db.Ant.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			lg.Warn(picoErrors.UnknownAnt, zap.Error(err))
			return nil, status.Error(codes.InvalidArgument, picoErrors.UnknownAnt)
		}
		lg.Error(picoErrors.QueryAnt, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}
	// update note
	if ant, err = ant.Update().SetNote(req.GetNote()).Save(ctx); err != nil {
		lg.Error(picoErrors.UpdateAnt, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// notify operators about ant's note update
	go pools.Pool.Ants.Send(pools.ToAntNoteResponse(ant))

	return &operatorv1.SetAntNoteResponse{}, nil
}

// SetAntNote set note to list of ants
func (s *server) SetAntsNote(ctx context.Context, req *operatorv1.SetAntsNoteRequest) (*operatorv1.SetAntsNoteResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SetAntsNote").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	for _, id := range req.GetIds() {
		// query ant
		ant, err := s.db.Ant.Get(ctx, id)
		if err != nil {
			if ent.IsNotFound(err) {
				lg.Warn(picoErrors.UnknownAnt, zap.Error(err))
				return nil, status.Error(codes.InvalidArgument, picoErrors.UnknownAnt)
			}
			lg.Error(picoErrors.QueryAnt, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}
		// update note
		if ant, err = ant.Update().SetNote(req.GetNote()).Save(ctx); err != nil {
			lg.Error(picoErrors.UpdateAnt, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}

		// notify operators about ant's note update
		go pools.Pool.Ants.Send(pools.ToAntNoteResponse(ant))
	}

	return &operatorv1.SetAntsNoteResponse{}, nil
}

// SetOperatorColor sets color to operator
func (s *server) SetOperatorColor(ctx context.Context, req *operatorv1.SetOperatorColorRequest) (*operatorv1.SetOperatorColorResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SetOperatorColor").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// query operator
	operator, err := s.db.Operator.
		Query().
		Where(operator.Username(req.GetUsername())).
		Only(ctx)
	if err != nil {
		lg.Error(picoErrors.QueryOperator, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}
	// update color
	if operator, err = operator.Update().SetColor(req.GetColor().GetValue()).Save(ctx); err != nil {
		lg.Error(picoErrors.UpdateOperator, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// notify operators about ant's color update
	go pools.Pool.Operators.Send(pools.ToOperatorColorResponse(operator))

	return &operatorv1.SetOperatorColorResponse{}, nil
}

// SetOperatorsColor sets color to list of operators
func (s *server) SetOperatorsColor(ctx context.Context, req *operatorv1.SetOperatorsColorRequest) (*operatorv1.SetOperatorsColorResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("SetOperatorColor").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	for _, username := range req.GetUsernames() {
		// query operator
		operator, err := s.db.Operator.
			Query().
			Where(operator.Username(username)).
			Only(ctx)
		if err != nil {
			lg.Error(picoErrors.QueryOperator, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}
		// update color
		if operator, err = operator.Update().SetColor(req.GetColor().GetValue()).Save(ctx); err != nil {
			lg.Error(picoErrors.UpdateOperator, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}

		// notify operators about ant's color update
		go pools.Pool.Operators.Send(pools.ToOperatorColorResponse(operator))
	}

	return &operatorv1.SetOperatorsColorResponse{}, nil
}

// NewChatMessage saves new message in chat created by operator
func (s *server) NewChatMessage(ctx context.Context, req *operatorv1.NewChatMessageRequest) (*operatorv1.NewChatMessageResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("NewChatMessage").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// query operator
	operator, err := s.db.Operator.
		Query().
		Where(operator.Username(username)).
		Only(ctx)
	if err != nil {
		lg.Error(picoErrors.QueryOperator, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}
	// save message
	message, err := s.db.Chat.
		Create().
		SetOperator(operator).
		SetMessage(req.GetMessage()).
		Save(ctx)
	if err != nil {
		lg.Error(picoErrors.SaveChatMessage, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// notify operators with new chat message
	go pools.Pool.Chat.Send(&operatorv1.ChatResponse{
		CreatedAt: timestamppb.New(message.CreatedAt),
		From:      wrapperspb.String(operator.Username),
		Message:   message.Message,
	})

	return &operatorv1.NewChatMessageResponse{}, nil
}

// NewCredential saves new credential
func (s *server) NewCredential(ctx context.Context, req *operatorv1.NewCredentialRequest) (*operatorv1.NewCredentialResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("NewCredential").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	c := s.db.Credential.Create()

	// username
	if req.GetUsername() != nil {
		if len(req.GetUsername().GetValue()) > shared.CredentialUsernameMaxLength {
			c.SetUsername(req.GetUsername().GetValue()[:shared.CredentialUsernameMaxLength])
		} else {
			c.SetUsername(req.GetUsername().GetValue())
		}
	}
	// password
	if req.GetPassword() != nil {
		if len(req.GetPassword().GetValue()) > shared.CredentialSecretMaxLength {
			c.SetSecret(req.GetPassword().GetValue()[:shared.CredentialSecretMaxLength])
		} else {
			c.SetSecret(req.GetPassword().GetValue())
		}
	}
	// realm
	if req.GetRealm() != nil {
		if len(req.GetRealm().GetValue()) > shared.CredentialRealmMaxLength {
			c.SetRealm(req.GetRealm().GetValue()[:shared.CredentialRealmMaxLength])
		} else {
			c.SetRealm(req.GetRealm().GetValue())
		}
	}
	// host
	if req.GetHost() != nil {
		if len(req.GetHost().GetValue()) > shared.CredentialHostMaxLength {
			c.SetHost(req.GetHost().GetValue()[:shared.CredentialHostMaxLength])
		} else {
			c.SetHost(req.GetHost().GetValue())
		}
	}

	var credential *ent.Credential
	var err error
	// save credential
	if credential, err = c.Save(ctx); err != nil {
		lg.Error(picoErrors.SaveCredential, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// notify operators with credentials update
	go pools.Pool.Credentials.Send(pools.ToCredentialResponse(credential))

	return &operatorv1.NewCredentialResponse{}, nil
}

// NewCommand creates new command to handle tasks and messages in
func (s *server) NewCommand(ss operatorv1.OperatorService_NewCommandServer) error {
	ctx := ss.Context()
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("NewCommand").With(zap.String("username", username))

	// first request - creation of command itself
	val, err := ss.Recv()
	if err != nil {
		lg.Error("receive command request", zap.Error(err))
		return status.Error(codes.Internal, "receive command request failed")
	}

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, val.GetCookie().GetValue()) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// if first request is not command -> drop
	if val.GetCommand() == nil {
		lg.Warn("first message is not request for command creation")
		return status.Error(codes.InvalidArgument, "invalid initial message")
	}

	// query operator
	operator, err := s.db.Operator.
		Query().
		Where(operator.Username(username)).
		Only(ctx)
	if err != nil {
		lg.Error(picoErrors.QueryOperator, zap.Error(err))
		return status.Error(codes.Internal, picoErrors.Internal)
	}
	// query ant
	ant, err := s.db.Ant.Get(ctx, val.GetCommand().GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			lg.Warn(picoErrors.UnknownAnt, zap.Error(err))
			return status.Error(codes.InvalidArgument, picoErrors.UnknownAnt)
		}
		lg.Error(picoErrors.QueryAnt, zap.Error(err))
		return status.Error(codes.Internal, picoErrors.Internal)
	}
	// create new command
	command, err := s.db.Command.
		Create().
		SetAnt(ant).
		SetOperator(operator).
		SetCmd(val.GetCommand().GetCmd()).
		SetVisible(val.GetCommand().GetVisible()).
		Save(ctx)
	if err != nil {
		lg.Error(picoErrors.SaveCommand, zap.Error(err))
		return status.Error(codes.Internal, picoErrors.Internal)
	}

	// notify operators with commands update
	go pools.Pool.Tasks.Send(&operatorv1.CommandResponse{
		Id:      command.ID,
		Aid:     ant.ID,
		Cmd:     command.Cmd,
		Author:  operator.Username,
		Created: timestamppb.New(command.CreatedAt),
		Visible: command.Visible,
	})

	defer func() {
		// close command on GRPC stream closing
		if _, err := command.Update().SetClosedAt(time.Now()).Save(ctx); err != nil {
			lg.Error(picoErrors.CloseCommand, zap.Error(err))
		}
	}()

	for {
		// recieve message
		msg, err := ss.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			lg.Error(picoErrors.Receive, zap.Error(err))
			return status.Error(codes.Internal, picoErrors.Receive)
		}

		// check if username and cookie are valid session identifiers
		if !pools.Pool.Hello.Validate(username, msg.GetCookie().GetValue()) {
			lg.Warn(picoErrors.InvalidSessionCookie)
			return status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
		}

		// create message in command
		if msg.GetMessage() != nil {
			// save command's message
			m, err := s.db.Message.
				Create().
				SetCommand(command).
				SetMessage(msg.GetMessage().GetMsg()).
				SetType(shared.TaskMessage(msg.GetMessage().GetType())).
				Save(ctx)
			if err != nil {
				lg.Error(picoErrors.SaveCommandMessage, zap.Error(err))
				return status.Error(codes.Internal, picoErrors.Internal)
			}

			// notify operators with command's updates
			go pools.Pool.Tasks.Send(&operatorv1.MessageResponse{
				Id:      int64(command.ID),
				Aid:     ant.ID,
				Mid:     int64(m.ID),
				Type:    uint32(m.Type),
				Message: m.Message,
				Created: timestamppb.New(m.CreatedAt),
			})

			continue
		}

		// create task in command
		if msg.GetTask() != nil {
			var raw []byte
			cap := shared.Capability(msg.GetTask().GetCap())
			switch cap {
			case shared.CapSleep:
				raw, err = cap.Marshal(msg.GetTask().GetSleep())
			case shared.CapLs:
				raw, err = cap.Marshal(msg.GetTask().GetLs())
			case shared.CapPwd:
				raw, err = cap.Marshal(msg.GetTask().GetPwd())
			case shared.CapCd:
				raw, err = cap.Marshal(msg.GetTask().GetCd())
			case shared.CapWhoami:
				raw, err = cap.Marshal(msg.GetTask().GetWhoami())
			case shared.CapPs:
				raw, err = cap.Marshal(msg.GetTask().GetPs())
			case shared.CapCat:
				raw, err = cap.Marshal(msg.GetTask().GetCat())
			case shared.CapExec:
				raw, err = cap.Marshal(msg.GetTask().GetExec())
			case shared.CapCp:
				raw, err = cap.Marshal(msg.GetTask().GetCp())
			case shared.CapJobs:
				raw, err = cap.Marshal(msg.GetTask().GetJobs())
			case shared.CapJobkill:
				raw, err = cap.Marshal(msg.GetTask().GetJobkill())
			case shared.CapKill:
				raw, err = cap.Marshal(msg.GetTask().GetKill())
			case shared.CapMv:
				raw, err = cap.Marshal(msg.GetTask().GetMv())
			case shared.CapMkdir:
				raw, err = cap.Marshal(msg.GetTask().GetMkdir())
			case shared.CapRm:
				raw, err = cap.Marshal(msg.GetTask().GetRm())
			case shared.CapExecAssembly:
				raw, err = cap.Marshal(msg.GetTask().GetExecAssembly())
			case shared.CapShellcodeInjection:
				raw, err = cap.Marshal(msg.GetTask().GetShellcodeInjection())
			case shared.CapDownload:
				raw, err = cap.Marshal(msg.GetTask().GetDownload())
			case shared.CapUpload:
				raw, err = cap.Marshal(msg.GetTask().GetUpload())
			case shared.CapPause:
				raw, err = cap.Marshal(msg.GetTask().GetPause())
			case shared.CapDestruct:
				raw, err = cap.Marshal(msg.GetTask().GetDestruct())
			case shared.CapExecDetach:
				raw, err = cap.Marshal(msg.GetTask().GetExecDetach())
			case shared.CapShell:
				raw, err = cap.Marshal(msg.GetTask().GetShell())
			case shared.CapPpid:
				raw, err = cap.Marshal(msg.GetTask().GetPpid())
			case shared.CapExit:
				raw, err = cap.Marshal(msg.GetTask().GetExit())
			default:
				err = fmt.Errorf("unknown capability %d", cap)
			}

			if err != nil {
				lg.Warn(picoErrors.MarshalCapability, zap.Error(err))
				return status.Error(codes.InvalidArgument, picoErrors.MarshalCapability)
			}

			// save task arguments in blob
			var blob *ent.Blobber
			h := utils.CalcHash(raw)
			if blob, err = s.db.Blobber.Query().Where(blobber.Hash(h)).Only(ctx); err != nil {
				if ent.IsNotFound(err) {
					// if not such blob - create new one
					if blob, err = s.db.Blobber.
						Create().
						SetBlob(raw).
						SetHash(h).
						SetSize(len(raw)).
						Save(ctx); err != nil {
						lg.Error(picoErrors.SaveBlob, zap.Error(err))
						return status.Error(codes.Internal, picoErrors.Internal)
					}
				} else {
					lg.Error(picoErrors.QueryBlob, zap.Error(err))
					return status.Error(codes.Internal, picoErrors.Internal)
				}
			}
			// create taks
			task, err := s.db.Task.
				Create().
				SetAnt(ant).
				SetCommand(command).
				SetBlobberArgs(blob).
				SetCap(cap).
				SetStatus(shared.StatusNew).
				Save(ctx)
			if err != nil {
				lg.Error(picoErrors.SaveCommandTask, zap.Error(err))
				return status.Error(codes.Internal, picoErrors.Internal)
			}

			// notify operators with command's updates
			go pools.Pool.Tasks.Send(&operatorv1.TaskResponse{
				Id:        int64(command.ID),
				Tid:       int64(task.ID),
				Aid:       ant.ID,
				Status:    uint32(task.Status),
				Created:   timestamppb.New(task.CreatedAt),
				OutputBig: false,
			})

			continue
		}
	}
	return nil
}

// CancelTasks cancels tasks created by operator which are in status NEW
func (s *server) CancelTasks(ctx context.Context, req *operatorv1.CancelTasksRequest) (*operatorv1.CancelTasksResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("CancelTasks").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// query ant
	ant, err := s.db.Ant.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			lg.Warn(picoErrors.UnknownAnt, zap.Error(err))
			return nil, status.Error(codes.InvalidArgument, picoErrors.UnknownAnt)
		}
		lg.Error(picoErrors.QueryAnt, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// query operator
	operator, err := s.db.Operator.
		Query().
		Where(operator.UsernameEQ(username)).
		Only(ctx)
	if err != nil {
		lg.Error(picoErrors.QueryOperator, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// query tasks
	tasks, err := s.db.Task.
		Query().
		WithCommand(func(q *ent.CommandQuery) {
			q.Where(command.AuthorIDEQ(operator.ID))
		}).
		Order(task.ByCreatedAt()).
		Where(task.StatusEQ(shared.StatusNew)).
		Where(task.AntIDEQ(ant.ID)).
		All(ctx)
	if err != nil {
		lg.Error(picoErrors.QueryCommandTasks, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	for _, task := range tasks {
		// set task's status CANCELLED
		task, err = task.Update().
			SetStatus(shared.StatusCancelled).
			SetDoneAt(time.Now()).
			Save(ctx)
		if err != nil {
			lg.Warn(picoErrors.UpdateCommandTask, zap.Error(err))
			continue
		}

		// notify operators with task updates
		go pools.Pool.Tasks.Send(&operatorv1.TaskStatusResponse{
			Id:     task.CommandID,
			Aid:    ant.ID,
			Tid:    task.ID,
			Status: uint32(task.Status),
		})
	}
	return &operatorv1.CancelTasksResponse{}, nil
}

// GetTaskOutput returns tasks's output directly. Used to gather blob in case of IsOutputBig is true
func (s *server) GetTaskOutput(ctx context.Context, req *operatorv1.GetTaskOutputRequest) (*operatorv1.GetTaskOutputResponse, error) {
	username := grpcauth.OperatorFromCtx(ctx)
	lg := s.lg.Named("GetTaskOutput").With(zap.String("username", username))
	cookie := req.GetCookie().GetValue()

	// check if username and cookie are valid session identifiers
	if !pools.Pool.Hello.Validate(username, cookie) {
		lg.Warn(picoErrors.InvalidSessionCookie)
		return nil, status.Error(codes.PermissionDenied, picoErrors.InvalidSessionCookie)
	}

	// query task
	task, err := s.db.Task.
		Query().
		WithBlobberOutput().
		Where(task.IDEQ(req.GetId())).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, picoErrors.UnknownTask)
		}
		lg.Error(picoErrors.QueryCommandTask, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	data := make([]byte, 0)

	// query blob
	blob, err := task.Edges.BlobberOutputOrErr()
	if err != nil {
		// if blob not found -> it doesn't exist for task
		if !ent.IsNotFound(err) {
			lg.Error(picoErrors.QueryBlob, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			data = blob.Blob
		}
	}

	return &operatorv1.GetTaskOutputResponse{
		Output: wrapperspb.Bytes(data),
	}, nil
}
