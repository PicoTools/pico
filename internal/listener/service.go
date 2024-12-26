package listener

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"

	commonv1 "github.com/PicoTools/pico-shared/proto/gen/common/v1"
	listenerv1 "github.com/PicoTools/pico-shared/proto/gen/listener/v1"
	operatorv1 "github.com/PicoTools/pico-shared/proto/gen/operator/v1"
	"github.com/PicoTools/pico-shared/shared"
	"github.com/PicoTools/pico/internal/constants"
	"github.com/PicoTools/pico/internal/ent"
	"github.com/PicoTools/pico/internal/ent/blobber"
	"github.com/PicoTools/pico/internal/ent/task"
	"github.com/PicoTools/pico/internal/errors"
	picoErrors "github.com/PicoTools/pico/internal/errors"
	"github.com/PicoTools/pico/internal/middleware/grpcauth"
	"github.com/PicoTools/pico/internal/pools"
	"github.com/PicoTools/pico/internal/types"
	"github.com/PicoTools/pico/internal/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// server implements UnimplementedListenerServiceServer
type server struct {
	listenerv1.UnimplementedListenerServiceServer
	db *ent.Client
	lg *zap.Logger
}

// UpdateListener updates information about listener
func (s *server) UpdateListener(ctx context.Context, req *listenerv1.UpdateListenerRequest) (*listenerv1.UpdateListenerResponse, error) {
	listenerId := grpcauth.ListenerFromCtx(ctx)
	lg := s.lg.Named("UpdateListener").With(zap.Int64("listener-id", listenerId))

	v := s.db.Listener.UpdateOneID(listenerId)

	// name
	if req.GetName() != nil {
		if len(req.GetName().GetValue()) > shared.ListenerNameMaxLength {
			v.SetName(req.GetName().GetValue()[:shared.ListenerNameMaxLength])
		} else {
			v.SetName(req.GetName().GetValue())
		}
	}
	// ip
	if req.GetIp() != nil {
		ip := types.Inet{}
		if err := ip.Scan(req.GetIp().GetValue()); err != nil {
			// ignore parsing error and continue
			lg.Warn(errors.ParseIP, zap.Error(err))
		} else {
			v.SetIP(ip)
		}
	}
	// port
	if req.GetPort() != nil {
		v.SetPort(uint16(req.GetPort().GetValue()))
	}

	l, err := v.Save(ctx)
	if err != nil {
		lg.Error(errors.UpdateListener, zap.Error(err))
		return nil, status.Errorf(codes.Internal, errors.Internal)
	}

	// notify operators with listener updates
	go pools.Pool.Listeners.Send(pools.ToListenerInfoResponse(l))

	return &listenerv1.UpdateListenerResponse{}, nil
}

// Register new ant
func (s *server) RegisterAnt(ctx context.Context, req *listenerv1.RegisterAntRequest) (*listenerv1.RegisterAntResponse, error) {
	listenerId := grpcauth.ListenerFromCtx(ctx)
	lg := s.lg.Named("RegisterAnt").With(zap.Int64("listener-id", listenerId))

	// check if ant with such ID already exists
	if _, err := s.db.Ant.Get(ctx, req.GetId()); err != nil {
		if !ent.IsNotFound(err) {
			lg.Error(errors.QueryAnt, zap.Error(err))
			return nil, status.Error(codes.Internal, errors.Internal)
		}
	} else {
		lg.Warn(errors.AntAlreadyExists)
		return nil, status.Error(codes.AlreadyExists, errors.AntAlreadyExists)
	}

	// get listener
	listener, err := s.db.Listener.Get(ctx, listenerId)
	if err != nil {
		lg.Error(errors.QueryListener, zap.Error(err))
		return nil, status.Errorf(codes.Internal, errors.Internal)
	}

	ant := s.db.Ant.Create()

	// id
	ant.SetID(req.GetId())
	// listener_id
	ant.SetListener(listener)
	// ext_ip
	if req.GetExtIp() != nil {
		ip := types.Inet{}
		if err = ip.Scan(req.GetExtIp().GetValue()); err != nil {
			// skip if IP is not scaned
			lg.Warn(errors.ParseExtIP, zap.Error(err))
		} else {
			ant.SetExtIP(ip)
		}
	}
	// int_ip
	if req.GetIntIp() != nil {
		ip := types.Inet{}
		if err = ip.Scan(req.GetIntIp().GetValue()); err != nil {
			// skip if IP is not scaned
			lg.Warn(errors.ParseIntIP, zap.Error(err))
		} else {
			ant.SetIntIP(ip)
		}
	}
	// os
	ant.SetOs(shared.AntOs(req.GetOs()))
	// os_meta
	if req.GetOsMeta() != nil {
		if len(req.GetOsMeta().GetValue()) > shared.AntOsMetaMaxLength {
			ant.SetOsMeta(req.GetOsMeta().GetValue()[:shared.AntOsMetaMaxLength])
		} else {
			ant.SetOsMeta(req.GetOsMeta().GetValue())
		}
	}
	// hostname
	if req.GetHostname() != nil {
		if len(req.GetHostname().GetValue()) > shared.AntHostnameMaxLength {
			ant.SetHostname(req.GetHostname().GetValue()[:shared.AntHostnameMaxLength])
		} else {
			ant.SetHostname(req.GetHostname().GetValue())
		}
	}
	// username
	if req.GetUsername() != nil {
		if len(req.GetUsername().GetValue()) > shared.AntUsernameMaxLength {
			ant.SetUsername(req.GetUsername().GetValue()[:shared.AntUsernameMaxLength])
		} else {
			ant.SetUsername(req.GetUsername().GetValue())
		}
	}
	// domain
	if req.GetDomain() != nil {
		if len(req.GetDomain().GetValue()) > shared.AntDomainMaxLength {
			ant.SetDomain(req.GetDomain().GetValue()[:shared.AntDomainMaxLength])
		} else {
			ant.SetDomain(req.GetDomain().GetValue())
		}
	}
	// privileged
	if req.GetPrivileged() != nil {
		ant.SetPrivileged(req.GetPrivileged().GetValue())
	}
	// proc_name
	if req.GetProcName() != nil {
		if len(req.GetProcName().GetValue()) > shared.AntProcessNameMaxLength {
			ant.SetProcessName(req.GetProcName().GetValue()[:shared.AntProcessNameMaxLength])
		} else {
			ant.SetProcessName(req.GetProcName().GetValue())
		}
	}
	// pid
	if req.GetPid() != nil {
		ant.SetPid(int64(req.GetPid().GetValue()))
	}
	// arch
	ant.SetArch(shared.AntArch(req.GetArch()))
	// sleep
	ant.SetSleep(req.GetSleep())
	// jitter
	ant.SetJitter(uint8(req.GetJitter()))
	// caps
	ant.SetCaps(req.GetCaps())

	// save ant
	var antObj *ent.Ant
	if antObj, err = ant.Save(ctx); err != nil {
		lg.Error(errors.SaveAnt, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.Internal)
	}

	// prepare message to chat with notification
	hostname := antObj.Hostname
	if hostname == "" {
		hostname = hex.EncodeToString([]byte(strconv.Itoa(int(antObj.ID))))
	}
	username := antObj.Username
	if username == "" {
		username = "[unknown]"
	}
	intIp := antObj.IntIP.String()
	if intIp == "" {
		intIp = "[unknown]"
	}
	ch, err := s.db.Chat.
		Create().
		SetIsServer(true).
		SetMessage(fmt.Sprintf("new ant [%s] %s@%s (%s)", fmt.Sprintf("%08x", antObj.ID), username, intIp, hostname)).
		Save(ctx)
	if err != nil {
		lg.Error(errors.SaveChatMessage, zap.Error(err))
		return nil, status.Error(codes.Internal, errors.Internal)
	}

	// notify operators with new chat message
	go pools.Pool.Chat.Send(&operatorv1.ChatResponse{
		CreatedAt: timestamppb.New(ch.CreatedAt),
		Message:   ch.Message,
		IsServer:  true,
	})

	// notify operators with new ant
	go pools.Pool.Ants.Send(&operatorv1.AntResponse{
		Id:         antObj.ID,
		Lid:        listener.ID,
		ExtIp:      wrapperspb.String(antObj.ExtIP.String()),
		IntIp:      wrapperspb.String(antObj.IntIP.String()),
		Os:         uint32(antObj.Os),
		OsMeta:     wrapperspb.String(antObj.OsMeta),
		Hostname:   wrapperspb.String(antObj.Hostname),
		Username:   wrapperspb.String(antObj.Username),
		Domain:     wrapperspb.String(antObj.Domain),
		Privileged: wrapperspb.Bool(antObj.Privileged),
		ProcName:   wrapperspb.String(antObj.ProcessName),
		Pid:        wrapperspb.UInt64(uint64(antObj.Pid)),
		Arch:       uint32(antObj.Arch),
		Sleep:      antObj.Sleep,
		Jitter:     uint32(antObj.Jitter),
		Caps:       antObj.Caps,
		Color:      wrapperspb.UInt32(antObj.Color),
		Note:       wrapperspb.String(antObj.Note),
		First:      timestamppb.New(antObj.First),
		Last:       timestamppb.New(antObj.Last),
	})

	return &listenerv1.RegisterAntResponse{}, nil
}

// GetTask returns task for ant from queue
func (s *server) GetTask(ctx context.Context, req *listenerv1.GetTaskRequest) (*listenerv1.GetTaskResponse, error) {
	listenerId := grpcauth.ListenerFromCtx(ctx)
	lg := s.lg.Named("GetTask").With(zap.Int64("listener-id", listenerId), zap.Uint32("ant-id", req.GetId()))

	// get ant by ID
	ant, err := s.db.Ant.Get(ctx, req.GetId())
	if err != nil {
		if !ent.IsNotFound(err) {
			lg.Error(picoErrors.QueryAnt, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			lg.Error("attempt to fetch task for unknown ant")
			return nil, status.Error(codes.InvalidArgument, picoErrors.UnknownAnt)
		}
	}

	// update ant's last checkout
	ant, err = ant.Update().
		SetLast(time.Now()).
		Save(ctx)
	if err != nil {
		lg.Error(picoErrors.UpdateLastAnt, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// notify operators with ant's last checkout update
	go pools.Pool.Ants.Send(pools.ToAntLastResponse(ant))

	// get first in queue task for ant
	task, err := s.db.Task.
		Query().
		WithBlobberArgs().
		Where(task.AntIDEQ(ant.ID)).
		Where(task.StatusEQ(shared.StatusNew)).
		Order(task.ByCreatedAt()).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return &listenerv1.GetTaskResponse{}, nil
		}
		lg.Error(picoErrors.QueryCommandTask, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// get task blob
	taskBlob, err := task.Edges.BlobberArgsOrErr()
	if err != nil {
		lg.Error(picoErrors.QueryBlob, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// prepare response
	response := &listenerv1.GetTaskResponse{
		Id:  task.ID,
		Cap: uint32(task.Cap),
	}
	switch task.Cap {
	case shared.CapSleep:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Sleep{
				Sleep: v.(*commonv1.CapSleep),
			}
		}
	case shared.CapLs:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Ls{
				Ls: v.(*commonv1.CapLs),
			}
		}
	case shared.CapPwd:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Pwd{
				Pwd: v.(*commonv1.CapPwd),
			}
		}
	case shared.CapCd:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Cd{
				Cd: v.(*commonv1.CapCd),
			}
		}
	case shared.CapWhoami:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Whoami{
				Whoami: v.(*commonv1.CapWhoami),
			}
		}
	case shared.CapPs:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Ps{
				Ps: v.(*commonv1.CapPs),
			}
		}
	case shared.CapCat:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Cat{
				Cat: v.(*commonv1.CapCat),
			}
		}
	case shared.CapExec:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Exec{
				Exec: v.(*commonv1.CapExec),
			}
		}
	case shared.CapCp:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Cp{
				Cp: v.(*commonv1.CapCp),
			}
		}
	case shared.CapJobs:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Jobs{
				Jobs: v.(*commonv1.CapJobs),
			}
		}
	case shared.CapJobkill:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Jobkill{
				Jobkill: v.(*commonv1.CapJobkill),
			}
		}
	case shared.CapKill:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Kill{
				Kill: v.(*commonv1.CapKill),
			}
		}
	case shared.CapMv:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Mv{
				Mv: v.(*commonv1.CapMv),
			}
		}
	case shared.CapMkdir:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Mkdir{
				Mkdir: v.(*commonv1.CapMkdir),
			}
		}
	case shared.CapRm:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Rm{
				Rm: v.(*commonv1.CapRm),
			}
		}
	case shared.CapExecAssembly:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_ExecAssembly{
				ExecAssembly: v.(*commonv1.CapExecAssembly),
			}
		}
	case shared.CapShellcodeInjection:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_ShellcodeInjection{
				ShellcodeInjection: v.(*commonv1.CapShellcodeInjection),
			}
		}
	case shared.CapDownload:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Download{
				Download: v.(*commonv1.CapDownload),
			}
		}
	case shared.CapUpload:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Upload{
				Upload: v.(*commonv1.CapUpload),
			}
		}
	case shared.CapPause:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Pause{
				Pause: v.(*commonv1.CapPause),
			}
		}
	case shared.CapDestruct:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Destruct{
				Destruct: v.(*commonv1.CapDestruct),
			}
		}
	case shared.CapExecDetach:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_ExecDetach{
				ExecDetach: v.(*commonv1.CapExecDetach),
			}
		}
	case shared.CapShell:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Shell{
				Shell: v.(*commonv1.CapShell),
			}
		}
	case shared.CapPpid:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Ppid{
				Ppid: v.(*commonv1.CapPpid),
			}
		}
	case shared.CapExit:
		if v, err := task.Cap.Unmarshal(taskBlob.Blob); err != nil {
			lg.Error(picoErrors.UnmarshalCapability, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			response.Body = &listenerv1.GetTaskResponse_Exit{
				Exit: v.(*commonv1.CapExit),
			}
		}
	default:
		lg.Error("unknown capability to unmarshal", zap.String("cap", task.Cap.String()))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// update task's push timestamp and status
	task, err = task.Update().
		SetPushedAt(time.Now()).
		SetStatus(shared.StatusInProgress).
		Save(ctx)
	if err != nil {
		lg.Error("unable update task info for push", zap.Error(err))
		return nil, status.Error(codes.Internal, "unable process task from DB")
	}

	// notify operators with task's status update
	go pools.Pool.Tasks.Send(&operatorv1.TaskStatusResponse{
		Id:     task.CommandID,
		Tid:    task.ID,
		Aid:    ant.ID,
		Status: uint32(task.Status),
	})

	return response, nil
}

// Save task's result from ant
func (s *server) PutResult(ctx context.Context, req *listenerv1.PutResultRequest) (*listenerv1.PutResultResponse, error) {
	listenerId := grpcauth.ListenerFromCtx(ctx)
	lg := s.lg.Named("PutResult").With(zap.Int64("listener-id", listenerId), zap.Uint32("ant-id", req.GetId()))

	// get ant by ID
	ant, err := s.db.Ant.Get(ctx, req.GetId())
	if err != nil {
		if !ent.IsNotFound(err) {
			lg.Error(picoErrors.QueryAnt, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		} else {
			lg.Error(picoErrors.UnknownAnt)
			return nil, status.Error(codes.InvalidArgument, picoErrors.UnknownAnt)
		}
	}

	// update ant's last checkout timestamp
	ant, err = ant.Update().
		SetLast(time.Now()).
		Save(ctx)
	if err != nil {
		lg.Error(picoErrors.UpdateLastAnt, zap.Error(err))
		return nil, status.Error(codes.Internal, picoErrors.Internal)
	}

	// notify operators with ant's last checkout update
	go pools.Pool.Ants.Send(pools.ToAntLastResponse(ant))

	// get task for ant by ID
	task, err := s.db.Task.Get(ctx, req.GetTid())
	if err != nil {
		if !ent.IsNotFound(err) {
			lg.Error(picoErrors.QueryCommandTask, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}
		lg.Error(picoErrors.UnknownTask)
		return nil, status.Error(codes.InvalidArgument, picoErrors.UnknownTask)
	}

	if task.Status != shared.StatusInProgress {
		// if status not "IN PROGRESS" -> drop saving (logical error)
		lg.Warn("attempt to save output for task with invalid status")
		return nil, status.Error(codes.InvalidArgument, "invalid task id")
	}

	if shared.TaskStatus(req.GetStatus()) == shared.StatusNew {
		// if ant's task status is "NEW" -> drop saving (logical error)
		lg.Warn("attempt to update task with invalid status")
		return nil, status.Error(codes.InvalidArgument, "invalid task status")
	}

	n := constants.TempTaskOutputPrefix + fmt.Sprintf("%d", task.ID)

	if shared.TaskStatus(req.GetStatus()) == shared.StatusInProgress {
		// if ant's task status is "IN PROGRESS" -> save output in temporary file
		if req.GetOutput() != nil {
			if err := os.WriteFile(n, req.GetOutput().GetValue(), os.FileMode(os.O_CREATE|os.O_APPEND|os.O_WRONLY)); err != nil {
				lg.Error("unable write output in temp file", zap.Error(err))
				return nil, status.Error(codes.Internal, picoErrors.Internal)
			}
			// TODO: notify operators with temporary output update
		}
	} else {
		// save output in DB
		data, err := os.ReadFile(n)
		if err != nil {
			if !os.IsNotExist(err) {
				lg.Error("unable read file with temp output", zap.Error(err))
				return nil, status.Error(codes.Internal, picoErrors.Internal)
			}
			if req.GetOutput() != nil {
				data = req.GetOutput().GetValue()
			}
		} else {
			if req.GetOutput() != nil {
				data = append(data, req.GetOutput().GetValue()...)
			}
		}
		// avoid "NOT NULL constraint" error
		if data == nil {
			data = []byte{}
		}
		// save output blob
		h := utils.CalcHash(data)
		blob, err := s.db.Blobber.
			Query().
			Where(blobber.HashEQ(h)).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				// new blob
				blob, err = s.db.Blobber.
					Create().
					SetBlob(data).
					SetHash(h).
					SetSize(len(data)).
					Save(ctx)
				if err != nil {
					lg.Error(picoErrors.SaveBlob, zap.Error(err))
					return nil, status.Error(codes.Internal, picoErrors.Internal)
				}
			} else {
				lg.Error(picoErrors.QueryBlob, zap.Error(err))
				return nil, status.Error(codes.Internal, picoErrors.Internal)
			}
		}
		// update task
		task, err = task.Update().
			SetBlobberOutput(blob).
			SetOutputBig(blob.Size > shared.TaskOutputMaxShowSize).
			SetDoneAt(time.Now()).
			SetStatus(shared.TaskStatus(req.GetStatus())).
			Save(ctx)
		if err != nil {
			lg.Error(picoErrors.SaveCommandTask, zap.Error(err))
			return nil, status.Error(codes.Internal, picoErrors.Internal)
		}
		// notify operators about task completed
		x := &operatorv1.TaskDoneResponse{
			Id:        task.CommandID,
			Tid:       task.ID,
			Aid:       ant.ID,
			Status:    uint32(task.Status),
			OutputBig: task.OutputBig,
			OutputLen: uint64(blob.Size),
		}
		if !task.OutputBig {
			x.Output = wrapperspb.Bytes(blob.Blob)
		}
		go pools.Pool.Tasks.Send(x)

		// if task was CapSleep -> update ant data and notify operators
		if task.Cap == shared.CapSleep {
			sleepBlob, err := s.db.Blobber.Get(ctx, task.ArgsID)
			if err != nil {
				lg.Error("unable query sleep blob from DB", zap.Error(err))
			} else {
				if v, err := task.Cap.Unmarshal(sleepBlob.Blob); err != nil {
					lg.Error("unable unmarshal task sleep arguments to proto", zap.Error(err))
				} else {
					x := v.(*commonv1.CapSleep)
					antUpdated, err := ant.Update().
						SetSleep(x.GetSleep()).
						SetJitter(uint8(x.GetJitter())).
						Save(ctx)
					if err != nil {
						lg.Error("unable update ant to save new sleep/jitter values", zap.Error(err))
					} else {
						go pools.Pool.Ants.Send(&operatorv1.AntSleepResponse{
							Id:     antUpdated.ID,
							Sleep:  antUpdated.Sleep,
							Jitter: uint32(antUpdated.Jitter),
						})
					}
				}
			}
		}
	}

	return &listenerv1.PutResultResponse{}, nil
}
