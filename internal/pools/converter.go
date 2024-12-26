package pools

import (
	"context"

	operatorv1 "github.com/PicoTools/pico-shared/proto/gen/operator/v1"
	"github.com/PicoTools/pico/internal/ent"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ToListenerResponse converts DB model Listener to protobuf ListenerResponse
func ToListenerResponse(listener *ent.Listener) *operatorv1.ListenerResponse {
	return &operatorv1.ListenerResponse{
		Id:    listener.ID,
		Name:  wrapperspb.String(listener.Name),
		Ip:    wrapperspb.String(listener.IP.String()),
		Port:  wrapperspb.UInt32(uint32(listener.Port)),
		Note:  wrapperspb.String(listener.Note),
		Last:  timestamppb.New(listener.Last),
		Color: wrapperspb.UInt32(listener.Color),
	}
}

// ToListenersResponse converts list of DB models Listener to protobuf ListenersResponse
func ToListenersResponse(listeners []*ent.Listener) *operatorv1.ListenersResponse {
	result := make([]*operatorv1.ListenerResponse, 0)
	for _, listener := range listeners {
		result = append(result, ToListenerResponse(listener))
	}
	return &operatorv1.ListenersResponse{
		Listeners: result,
	}
}

// ToListenerNoteResponse converts DB model Listener to protobuf ListenerNoteResponse
func ToListenerNoteResponse(listener *ent.Listener) *operatorv1.ListenerNoteResponse {
	return &operatorv1.ListenerNoteResponse{
		Id:   listener.ID,
		Note: wrapperspb.String(listener.Note),
	}
}

// ToCredentialNoteResponse converts DB model Credential to protobuf CredentialNoteResponse
func ToCredentialNoteResponse(credential *ent.Credential) *operatorv1.CredentialNoteResponse {
	return &operatorv1.CredentialNoteResponse{
		Id:   credential.ID,
		Note: wrapperspb.String(credential.Note),
	}
}

// ToCredentialColorResponse converts DB model Credential to protobuf CredentialColorResponse
func ToCredentialColorResponse(credential *ent.Credential) *operatorv1.CredentialColorResponse {
	return &operatorv1.CredentialColorResponse{
		Id:    credential.ID,
		Color: wrapperspb.UInt32(credential.Color),
	}
}

// ToListenerColorResponse converts DB model Listener to protobuf ListenerColorResponse
func ToListenerColorResponse(listener *ent.Listener) *operatorv1.ListenerColorResponse {
	return &operatorv1.ListenerColorResponse{
		Id:    listener.ID,
		Color: wrapperspb.UInt32(listener.Color),
	}
}

// ToListenerInfoResponse converts DB model Listener to protobuf ListenerInfoResponse
func ToListenerInfoResponse(listener *ent.Listener) *operatorv1.ListenerInfoResponse {
	return &operatorv1.ListenerInfoResponse{
		Id:   listener.ID,
		Name: wrapperspb.String(listener.Name),
		Ip:   wrapperspb.String(listener.IP.String()),
		Port: wrapperspb.UInt32(uint32(listener.Port)),
	}
}

// ToAntColorResponse converts DB model Ant to protobuf AntColorResponse
func ToAntColorResponse(ant *ent.Ant) *operatorv1.AntColorResponse {
	return &operatorv1.AntColorResponse{
		Id:    ant.ID,
		Color: wrapperspb.UInt32(ant.Color),
	}
}

// ToAntNoteResponse converts DB model Ant to protobuf AntColorResponse
func ToAntNoteResponse(ant *ent.Ant) *operatorv1.AntNoteResponse {
	return &operatorv1.AntNoteResponse{
		Id:   ant.ID,
		Note: wrapperspb.String(ant.Note),
	}
}

// ToAntResponse converts DB model Ant to protobuf AntResponse
func ToAntResponse(ant *ent.Ant) (*operatorv1.AntResponse, error) {
	listener, err := ant.Edges.ListenerOrErr()
	if err != nil {
		return nil, err
	}
	return &operatorv1.AntResponse{
		Id:         ant.ID,
		Lid:        listener.ID,
		ExtIp:      wrapperspb.String(ant.ExtIP.String()),
		IntIp:      wrapperspb.String(ant.IntIP.String()),
		Os:         uint32(ant.Os),
		OsMeta:     wrapperspb.String(ant.OsMeta),
		Hostname:   wrapperspb.String(ant.Hostname),
		Username:   wrapperspb.String(ant.Username),
		Domain:     wrapperspb.String(ant.Domain),
		Privileged: wrapperspb.Bool(ant.Privileged),
		ProcName:   wrapperspb.String(ant.ProcessName),
		Pid:        wrapperspb.UInt64(uint64(ant.Pid)),
		Arch:       uint32(ant.Arch),
		Sleep:      ant.Sleep,
		Jitter:     uint32(ant.Jitter),
		Caps:       ant.Caps,
		Color:      wrapperspb.UInt32(ant.Color),
		Note:       wrapperspb.String(ant.Note),
		First:      timestamppb.New(ant.First),
		Last:       timestamppb.New(ant.Last),
	}, nil
}

// ToAntsResponse converts list of DB model Ant to protobuf AntsResponse
func ToAntsResponse(ants []*ent.Ant) *operatorv1.AntsResponse {
	result := make([]*operatorv1.AntResponse, 0)
	for _, ant := range ants {
		antResponse, err := ToAntResponse(ant)
		if err != nil {
			continue
		}
		result = append(result, antResponse)
	}
	return &operatorv1.AntsResponse{
		Ants: result,
	}
}

// ToAntLastResponse converts DB model Ant to protobuf AntLastResponse
func ToAntLastResponse(ant *ent.Ant) *operatorv1.AntLastResponse {
	return &operatorv1.AntLastResponse{
		Id:   ant.ID,
		Last: timestamppb.New(ant.Last),
	}
}

// ToOperatorResponse converts DB model Operator to protobuf OperatorResponse
func ToOperatorResponse(operator *ent.Operator) *operatorv1.OperatorResponse {
	return &operatorv1.OperatorResponse{
		Username: operator.Username,
		Color:    wrapperspb.UInt32(operator.Color),
		Last:     timestamppb.New(operator.Last),
	}
}

// ToOperatorColorResponse converts DB model Operator to protobuf OperatorColorResponse
func ToOperatorColorResponse(operator *ent.Operator) *operatorv1.OperatorColorResponse {
	return &operatorv1.OperatorColorResponse{
		Username: operator.Username,
		Color:    wrapperspb.UInt32(operator.Color),
	}
}

// ToOperatorsResponse converts list of DB model Operator to protobuf OperatorsResponse
func ToOperatorsResponse(operators []*ent.Operator) *operatorv1.OperatorsResponse {
	result := make([]*operatorv1.OperatorResponse, 0)
	for _, operator := range operators {
		result = append(result, &operatorv1.OperatorResponse{
			Username: operator.Username,
			Color:    wrapperspb.UInt32(operator.Color),
			Last:     timestamppb.New(operator.Last),
		})
	}
	return &operatorv1.OperatorsResponse{
		Operators: result,
	}
}

// ToChatMessageResponse converts DB model Chat to protobuf ChatResponse
func ToChatMessageResponse(chatMessage *ent.Chat) (*operatorv1.ChatResponse, error) {
	response := &operatorv1.ChatResponse{
		CreatedAt: timestamppb.New(chatMessage.CreatedAt),
		Message:   chatMessage.Message,
	}
	if chatMessage.IsServer {
		response.IsServer = chatMessage.IsServer
	} else {
		if operator, err := chatMessage.Edges.OperatorOrErr(); err != nil {
			if ent.IsNotLoaded(err) {
				operator, err = chatMessage.QueryOperator().Only(context.Background())
				if err != nil {
					return nil, err
				} else {
					response.From = wrapperspb.String(operator.Username)
				}
			} else {
				return nil, err
			}
		} else {
			response.From = wrapperspb.String(operator.Username)
		}
	}
	return response, nil
}

// ToChatMessagesResponse converts list of DB model Chat to ChatMessagesResponse
func ToChatMessagesResponse(chatMessages []*ent.Chat) *operatorv1.ChatMessagesResponse {
	result := make([]*operatorv1.ChatResponse, 0)
	for _, chatMessage := range chatMessages {
		chatResponse, err := ToChatMessageResponse(chatMessage)
		if err != nil {
			continue
		}
		result = append(result, chatResponse)
	}
	return &operatorv1.ChatMessagesResponse{
		Messages: result,
	}
}

// ToCredentialResponse converts DB model Credential to CredentialResponse
func ToCredentialResponse(credential *ent.Credential) *operatorv1.CredentialResponse {
	return &operatorv1.CredentialResponse{
		Id:        credential.ID,
		Username:  wrapperspb.String(credential.Username),
		Password:  wrapperspb.String(credential.Secret),
		Realm:     wrapperspb.String(credential.Realm),
		Host:      wrapperspb.String(credential.Host),
		CreatedAt: timestamppb.New(credential.CreatedAt),
		Note:      wrapperspb.String(credential.Note),
		Color:     wrapperspb.UInt32(credential.Color),
	}
}

// ToCredentialsResponse converts list of DB model Credential to CredentialsResponse
func ToCredentialsResponse(credentials []*ent.Credential) *operatorv1.CredentialsResponse {
	results := make([]*operatorv1.CredentialResponse, 0)
	for _, credential := range credentials {
		results = append(results, ToCredentialResponse(credential))
	}
	return &operatorv1.CredentialsResponse{
		Credentials: results,
	}
}
