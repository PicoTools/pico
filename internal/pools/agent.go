package pools

import (
	"strings"

	operatorv1 "github.com/PicoTools/pico/pkg/proto/operator/v1"

	"github.com/lrita/cmap"
)

type SubscribeAgents struct {
	username   string
	disconnect chan struct{}
	data       chan any
	err        chan error
	stream     operatorv1.OperatorService_SubscribeAgentsServer
}

type PoolAgents struct {
	pool cmap.Map[string, *SubscribeAgents]
}

// Error returns error appeared in process of stream handling
func (s *SubscribeAgents) Error() chan error {
	return s.err
}

// IsDisconnect notify if operator must be disconnected
func (s *SubscribeAgents) IsDisconnect() chan struct{} {
	return s.disconnect
}

// Add adds operator's stream to pool
func (p *PoolAgents) Add(cookie, username string, stream operatorv1.OperatorService_SubscribeAgentsServer) {
	p.pool.Store(cookie, &SubscribeAgents{
		username:   username,
		disconnect: make(chan struct{}, 1),
		data:       make(chan any, 1),
		err:        make(chan error),
		stream:     stream,
	})
}

// Disconnect notify about necessity of operator's disconnection
func (p *PoolAgents) Disconnect(cookie string) {
	v, ok := p.pool.Load(cookie)
	if ok {
		v.disconnect <- struct{}{}
	}
}

// Remove removes operator's stream from pool
func (p *PoolAgents) Remove(cookie string) {
	p.pool.Delete(cookie)
}

// Get returns operator's subscription object
func (p *PoolAgents) Get(cookie string) *SubscribeAgents {
	v, _ := p.pool.Load(cookie)
	return v
}

// Send sends proto message to subscribed operators
func (p *PoolAgents) Send(val any) {
	msg := &operatorv1.SubscribeAgentsResponse{}

	switch val := val.(type) {
	case *operatorv1.AgentResponse:
		msg.Response = &operatorv1.SubscribeAgentsResponse_Agent{
			Agent: val,
		}
	case *operatorv1.AgentColorResponse:
		msg.Response = &operatorv1.SubscribeAgentsResponse_Color{
			Color: val,
		}
	case *operatorv1.AgentNoteResponse:
		msg.Response = &operatorv1.SubscribeAgentsResponse_Note{
			Note: val,
		}
	case *operatorv1.AgentLastResponse:
		msg.Response = &operatorv1.SubscribeAgentsResponse_Last{
			Last: val,
		}
	case *operatorv1.AgentSleepResponse:
		msg.Response = &operatorv1.SubscribeAgentsResponse_Sleep{
			Sleep: val,
		}
	case *operatorv1.AgentsResponse:
		msg.Response = &operatorv1.SubscribeAgentsResponse_Agents{
			Agents: val,
		}
	default:
		return
	}

	p.pool.Range(func(_ string, v *SubscribeAgents) bool {
		if err := v.stream.Send(msg); err != nil {
			v.err <- err
		}
		return true
	})
}

// Exists check if operator exists in pool by its username
func (p *PoolAgents) Exists(username string) bool {
	f := false
	p.pool.Range(func(_ string, v *SubscribeAgents) bool {
		if strings.Compare(v.username, username) == 0 {
			f = true
		}
		return !f
	})
	return f
}
