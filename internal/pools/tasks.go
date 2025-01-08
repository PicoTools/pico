package pools

import (
	"strings"

	operatorv1 "github.com/PicoTools/pico-shared/proto/gen/operator/v1"
	"github.com/lrita/cmap"
)

type SubscribeTasks struct {
	username   string
	disconnect chan struct{}
	// map of agent's ids
	ids    map[uint32]struct{}
	data   chan any
	err    chan error
	stream operatorv1.OperatorService_SubscribeTasksServer
}

type PoolTasks struct {
	pool cmap.Map[string, *SubscribeTasks]
}

// Error returns error appeared in process of stream handling
func (s *SubscribeTasks) Error() chan error {
	return s.err
}

// IsDisconnect notify if operator must be disconnected
func (s *SubscribeTasks) IsDisconnect() chan struct{} {
	return s.disconnect
}

// Add adds operator's stream to pool
func (p *PoolTasks) Add(cookie, username string, stream operatorv1.OperatorService_SubscribeTasksServer) {
	p.pool.Store(cookie, &SubscribeTasks{
		username:   username,
		disconnect: make(chan struct{}, 1),
		data:       make(chan any, 1),
		err:        make(chan error),
		ids:        make(map[uint32]struct{}),
		stream:     stream,
	})
}

// AddAgent saves ID of agent for operator's polling
func (p *PoolTasks) AddAgent(cookie string, id uint32) {
	if s, ok := p.pool.Load(cookie); ok {
		s.ids[id] = struct{}{}
	}
}

// DeleteAgent removes ID of agent from operator's polling
func (p *PoolTasks) DeleteAgent(cookie string, id uint32) {
	if s, ok := p.pool.Load(cookie); ok {
		delete(s.ids, id)
	}
}

// Disconnect notify about necessity of operator's disconnection
func (p *PoolTasks) Disconnect(cookie string) {
	v, ok := p.pool.Load(cookie)
	if ok {
		v.disconnect <- struct{}{}
	}
}

// Remove removes operator's stream from pool
func (p *PoolTasks) Remove(cookie string) {
	p.pool.Delete(cookie)
}

// Get returns operator's subscription object
func (p *PoolTasks) Get(cookie string) *SubscribeTasks {
	t, _ := p.pool.Load(cookie)
	return t
}

// Send sends proto message to subscribed operators
func (p *PoolTasks) Send(val any) {
	msg := &operatorv1.SubscribeTasksResponse{}
	var id uint32
	var author string

	switch val := val.(type) {
	case *operatorv1.CommandResponse:
		id = val.GetAid()
		author = val.GetAuthor()
		msg.Type = &operatorv1.SubscribeTasksResponse_Command{
			Command: val,
		}
	case *operatorv1.MessageResponse:
		id = val.GetAid()
		msg.Type = &operatorv1.SubscribeTasksResponse_Message{
			Message: val,
		}
	case *operatorv1.TaskResponse:
		id = val.GetAid()
		msg.Type = &operatorv1.SubscribeTasksResponse_Task{
			Task: val,
		}
	case *operatorv1.TaskStatusResponse:
		id = val.GetAid()
		msg.Type = &operatorv1.SubscribeTasksResponse_TaskStatus{
			TaskStatus: val,
		}
	case *operatorv1.TaskDoneResponse:
		id = val.GetAid()
		msg.Type = &operatorv1.SubscribeTasksResponse_TaskDone{
			TaskDone: val,
		}
	default:
		return
	}

	p.pool.Range(func(_ string, v *SubscribeTasks) bool {
		if _, ok := v.ids[id]; ok {
			// if command is invisible - skip sendings
			if x, ok := val.(*operatorv1.CommandResponse); ok {
				if !x.Visible && author != v.username {
					return true
				}
			}
			if err := v.stream.Send(msg); err != nil {
				v.err <- err
			}
		}
		return true
	})
}

// Exists check if operator exists in pool by its username
func (p *PoolTasks) Exists(username string) bool {
	f := false
	p.pool.Range(func(c string, v *SubscribeTasks) bool {
		if strings.Compare(v.username, username) == 0 {
			f = true
		}
		return !f
	})
	return f
}
