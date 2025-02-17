package management

import (
	"context"

	"github.com/PicoTools/pico/internal/ent"
	"github.com/PicoTools/pico/internal/ent/listener"
	"github.com/PicoTools/pico/internal/ent/operator"
	"github.com/PicoTools/pico/internal/ent/pki"
	"github.com/PicoTools/pico/internal/errors"
	"github.com/PicoTools/pico/internal/pools"
	"github.com/PicoTools/pico/internal/utils"
	managementv1 "github.com/PicoTools/pico/pkg/proto/management/v1"
	operatorv1 "github.com/PicoTools/pico/pkg/proto/operator/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type server struct {
	managementv1.UnimplementedManagementServiceServer
	db *ent.Client
	lg *zap.Logger
}

// GetOperators returns list of registred operators
func (s *server) GetOperators(ctx context.Context, req *managementv1.GetOperatorsRequest) (*managementv1.GetOperatorsResponse, error) {
	lg := s.lg.Named("GetOperators")

	operators, err := s.db.Operator.
		Query().
		All(ctx)
	if err != nil {
		lg.Error(errors.QueryOperators, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.Internal)
	}

	var result []*managementv1.Operator
	for _, operator := range operators {
		result = append(result, &managementv1.Operator{
			Username: operator.Username,
			Token:    wrapperspb.String(operator.Token),
			Last:     timestamppb.New(operator.Last),
		})
	}

	return &managementv1.GetOperatorsResponse{Operators: result}, nil
}

// NewOperator register new operator
func (s *server) NewOperator(ctx context.Context, req *managementv1.NewOperatorRequest) (*managementv1.NewOperatorResponse, error) {
	lg := s.lg.Named("NewOperator").With(zap.String("username", req.GetUsername()))

	// check if operator already exists in database
	_, err := s.db.Operator.Query().Where(operator.Username(req.GetUsername())).Only(ctx)
	if err == nil {
		lg.Warn(errors.OperatorAlreadyExists)
		return nil, status.Error(codes.AlreadyExists, errors.OperatorAlreadyExists)
	} else if !ent.IsNotFound(err) {
		lg.Error(errors.QueryOperator, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.Internal)
	}

	operator, err := s.db.Operator.
		Create().
		SetUsername(req.GetUsername()).
		SetToken(utils.RandString(32)).
		Save(ctx)
	if err != nil {
		lg.Error(errors.SaveOperator, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.Internal)
	}

	// notify operators with new registration of new operator
	go pools.Pool.Operators.Send(&operatorv1.OperatorResponse{
		Username: operator.Username,
		Last:     timestamppb.New(operator.Last),
		Color:    wrapperspb.UInt32(operator.Color),
	})

	return &managementv1.NewOperatorResponse{Operator: &managementv1.Operator{
		Username: operator.Username,
		Token:    wrapperspb.String(operator.Token),
		Last:     timestamppb.New(operator.Last),
	}}, nil
}

// RevokeOperator revokes access token for operator
func (s *server) RevokeOperator(ctx context.Context, req *managementv1.RevokeOperatorRequest) (*managementv1.RevokeOperatorResponse, error) {
	lg := s.lg.Named("RevokeOperator").With(zap.String("username", req.GetUsername()))

	// check if operator already exists in database
	if _, err := s.db.Operator.Query().Where(operator.Username(req.GetUsername())).Only(ctx); err != nil {
		if ent.IsNotFound(err) {
			lg.Warn(errors.UnknownOperator)
			return nil, status.Error(codes.InvalidArgument, errors.UnknownOperator)
		}
		lg.Error(errors.QueryOperator, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.Internal)
	}

	if _, err := s.db.Operator.
		Update().
		Where(operator.Username(req.GetUsername())).
		ClearToken().
		Save(ctx); err != nil {
		lg.Error(errors.UpdateOperator, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.Internal)
	}
	return &managementv1.RevokeOperatorResponse{}, nil
}

// RegenerateOperator regenerates access token for operator
func (s *server) RegenerateOperator(ctx context.Context, req *managementv1.RegenerateOperatorRequest) (*managementv1.RegenerateOperatorResponse, error) {
	lg := s.lg.Named("RegenerateOperator")

	// check if operator already exists in database
	o, err := s.db.Operator.Query().Where(operator.Username(req.GetUsername())).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			lg.Warn(errors.UnknownOperator)
			return nil, status.Error(codes.InvalidArgument, errors.UnknownOperator)
		}
		lg.Error(errors.QueryOperator, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.Internal)
	}

	token := utils.RandString(32)

	if _, err = s.db.Operator.
		Update().
		Where(operator.Username(req.GetUsername())).
		SetToken(token).
		Save(ctx); err != nil {
		lg.Error(errors.UpdateOperator, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.Internal)
	}

	return &managementv1.RegenerateOperatorResponse{Operator: &managementv1.Operator{
		Username: o.Username,
		Token:    wrapperspb.String(token),
		Last:     timestamppb.New(o.Last),
	}}, nil
}

// GetListeners returns list of listeners
func (s *server) GetListeners(ctx context.Context, req *managementv1.GetListenersRequest) (*managementv1.GetListenersResponse, error) {
	lg := s.lg.Named("GetListeners")

	listeners, err := s.db.Listener.Query().All(ctx)
	if err != nil {
		lg.Error(errors.QueryListeners, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.Internal)
	}

	var result []*managementv1.Listener
	for _, listener := range listeners {
		result = append(result, &managementv1.Listener{
			Id:    listener.ID,
			Name:  wrapperspb.String(listener.Name),
			Ip:    wrapperspb.String(listener.IP.String()),
			Port:  wrapperspb.UInt32(uint32(listener.Port)),
			Token: wrapperspb.String(listener.Token),
			Last:  timestamppb.New(listener.Last),
		})
	}

	return &managementv1.GetListenersResponse{Listeners: result}, nil
}

// NewListener registers new listener
func (s *server) NewListener(ctx context.Context, req *managementv1.NewListenerRequest) (*managementv1.NewListenerResponse, error) {
	lg := s.lg.Named("NewListener")

	listener, err := s.db.Listener.
		Create().
		SetToken(utils.RandString(32)).
		Save(ctx)
	if err != nil {
		lg.Error(errors.SaveListener, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.Internal)
	}

	// notify operators with registration of new listener
	go pools.Pool.Listeners.Send(&operatorv1.ListenerResponse{
		Id:    listener.ID,
		Color: wrapperspb.UInt32(listener.Color),
	})

	return &managementv1.NewListenerResponse{Listener: &managementv1.Listener{
		Id:    listener.ID,
		Token: wrapperspb.String(listener.Token),
		Last:  timestamppb.New(listener.Last),
	}}, nil
}

// RevokeListener revokes access token for listener
func (s *server) RevokeListener(ctx context.Context, req *managementv1.RevokeListenerRequest) (*managementv1.RevokeListenerResponse, error) {
	lg := s.lg.Named("RevokeListener").With(zap.Int64("listener-id", req.GetId()))

	// check if listener already exists in database
	if _, err := s.db.Listener.Get(ctx, req.GetId()); err != nil {
		if ent.IsNotFound(err) {
			lg.Warn(errors.UnknownListener)
			return nil, status.Error(codes.InvalidArgument, errors.UnknownListener)
		}
		lg.Error(errors.QueryListener, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.Internal)
	}

	if _, err := s.db.Listener.
		Update().
		Where(listener.ID(req.GetId())).
		ClearToken().
		Save(ctx); err != nil {
		lg.Error(errors.UpdateListener, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.Internal)
	}

	return &managementv1.RevokeListenerResponse{}, nil
}

// RegenerateListener regenerates access token for listener
func (s *server) RegenerateListener(ctx context.Context, req *managementv1.RegenerateListenerRequest) (*managementv1.RegenerateListenerResponse, error) {
	lg := s.lg.Named("RegenerateListener").With(zap.Int64("listener-id", req.GetId()))

	// check if listener already exists in database
	l, err := s.db.Listener.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			lg.Warn(errors.UnknownListener)
			return nil, status.Error(codes.InvalidArgument, errors.UnknownListener)
		}
		lg.Error(errors.QueryListener, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.Internal)
	}

	token := utils.RandString(32)

	if _, err := s.db.Listener.
		Update().
		Where(listener.ID(req.GetId())).
		SetToken(token).
		Save(ctx); err != nil {
		lg.Error(errors.UpdateListener, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.Internal)
	}

	return &managementv1.RegenerateListenerResponse{Listener: &managementv1.Listener{
		Id:    l.ID,
		Name:  wrapperspb.String(l.Name),
		Ip:    wrapperspb.String(l.IP.String()),
		Port:  wrapperspb.UInt32(uint32(l.Port)),
		Token: wrapperspb.String(token),
		Last:  timestamppb.New(l.Last),
	}}, nil
}

// GetCertCA returns CA certificate from PKI
func (s *server) GetCertCA(ctx context.Context, req *managementv1.GetCertCARequest) (*managementv1.GetCertCAResponse, error) {
	lg := s.lg.Named("GetCertCA")

	pki, err := s.db.Pki.
		Query().
		Where(pki.TypeEQ(pki.TypeCa)).
		Only(ctx)
	if err != nil {
		lg.Error(errors.QueryPkiCa, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.QueryPkiCa)
	}

	return &managementv1.GetCertCAResponse{
		Certificate: &managementv1.Certificate{
			Data: string(pki.Cert),
		},
	}, nil
}

// GetCertOperator returns operator's certificate from PKI
func (s *server) GetCertOperator(ctx context.Context, req *managementv1.GetCertOperatorRequest) (*managementv1.GetCertOperatorResponse, error) {
	lg := s.lg.Named("GetCertOperator")

	pki, err := s.db.Pki.
		Query().
		Where(pki.TypeEQ(pki.TypeOperator)).
		Only(ctx)
	if err != nil {
		lg.Error(errors.QueryPkiOperator, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.QueryPkiOperator)
	}

	return &managementv1.GetCertOperatorResponse{
		Certificate: &managementv1.Certificate{
			Data: string(pki.Cert),
		},
	}, nil
}

// GetCertListener returns listener's certificate from PKI
func (s *server) GetCertListener(ctx context.Context, req *managementv1.GetCertListenerRequest) (*managementv1.GetCertListenerResponse, error) {
	lg := s.lg.Named("GetCertListener")

	pki, err := s.db.Pki.
		Query().
		Where(pki.TypeEQ(pki.TypeListener)).
		Only(ctx)
	if err != nil {
		lg.Error(errors.QueryPkiListener, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.QueryPkiListener)
	}

	return &managementv1.GetCertListenerResponse{
		Certificate: &managementv1.Certificate{
			Data: string(pki.Cert),
		},
	}, nil
}
