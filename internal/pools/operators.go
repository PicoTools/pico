package pools

import (
	"strings"

	operatorv1 "github.com/PicoTools/pico/pkg/proto/operator/v1"

	"github.com/lrita/cmap"
)

type SubscribeOperators struct {
	username   string
	disconnect chan struct{}
	data       chan any
	err        chan error
	stream     operatorv1.OperatorService_SubscribeOperatorsServer
}

type PoolOperators struct {
	pool cmap.Map[string, *SubscribeOperators]
}

// Error returns error appeared in process of stream handling
func (s *SubscribeOperators) Error() chan error {
	return s.err
}

// IsDisconnect notify if operator must be disconnected
func (s *SubscribeOperators) IsDisconnect() chan struct{} {
	return s.disconnect
}

// Add adds operator's stream to pool
func (p *PoolOperators) Add(cookie, username string, stream operatorv1.OperatorService_SubscribeOperatorsServer) {
	p.pool.Store(cookie, &SubscribeOperators{
		username:   username,
		disconnect: make(chan struct{}, 1),
		data:       make(chan any, 1),
		err:        make(chan error),
		stream:     stream,
	})
}

// Disconnect notify about necessity of operator's disconnection
func (p *PoolOperators) Disconnect(cookie string) {
	v, ok := p.pool.Load(cookie)
	if ok {
		v.disconnect <- struct{}{}
	}
}

// Remove removes operator's stream from pool
func (p *PoolOperators) Remove(cookie string) {
	p.pool.Delete(cookie)
}

// Get returns operator's subscription object
func (p *PoolOperators) Get(cookie string) *SubscribeOperators {
	v, _ := p.pool.Load(cookie)
	return v
}

// Send sends proto message to subscribed operators
func (p *PoolOperators) Send(val any) {
	msg := &operatorv1.SubscribeOperatorsResponse{}

	switch val := val.(type) {
	case *operatorv1.OperatorLastResponse:
		msg.Response = &operatorv1.SubscribeOperatorsResponse_Last{
			Last: val,
		}
	case *operatorv1.OperatorResponse:
		msg.Response = &operatorv1.SubscribeOperatorsResponse_Operator{
			Operator: val,
		}
	case *operatorv1.OperatorColorResponse:
		msg.Response = &operatorv1.SubscribeOperatorsResponse_Color{
			Color: val,
		}
	case *operatorv1.OperatorsResponse:
		msg.Response = &operatorv1.SubscribeOperatorsResponse_Operators{
			Operators: val,
		}
	default:
		return
	}

	p.pool.Range(func(_ string, v *SubscribeOperators) bool {
		if err := v.stream.Send(msg); err != nil {
			v.err <- err
		}
		return true
	})
}

// Exists check if operator exists in pool by its username
func (p *PoolOperators) Exists(username string) bool {
	f := false
	p.pool.Range(func(_ string, v *SubscribeOperators) bool {
		if strings.Compare(v.username, username) == 0 {
			f = true
		}
		return !f
	})
	return f
}
