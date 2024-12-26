package pools

import (
	"strings"

	operatorv1 "github.com/PicoTools/pico-shared/proto/gen/operator/v1"

	"github.com/lrita/cmap"
)

type SubscribeAnts struct {
	username   string
	disconnect chan struct{}
	data       chan any
	err        chan error
	stream     operatorv1.OperatorService_SubscribeAntsServer
}

type PoolAnts struct {
	pool cmap.Map[string, *SubscribeAnts]
}

// Error returns error appeared in process of stream handling
func (s *SubscribeAnts) Error() chan error {
	return s.err
}

// IsDisconnect notify if operator must be disconnected
func (s *SubscribeAnts) IsDisconnect() chan struct{} {
	return s.disconnect
}

// Add adds operator's stream to pool
func (p *PoolAnts) Add(cookie, username string, stream operatorv1.OperatorService_SubscribeAntsServer) {
	p.pool.Store(cookie, &SubscribeAnts{
		username:   username,
		disconnect: make(chan struct{}, 1),
		data:       make(chan any, 1),
		err:        make(chan error),
		stream:     stream,
	})
}

// Disconnect notify about necessity of operator's disconnection
func (p *PoolAnts) Disconnect(cookie string) {
	v, ok := p.pool.Load(cookie)
	if ok {
		v.disconnect <- struct{}{}
	}
}

// Remove removes operator's stream from pool
func (p *PoolAnts) Remove(cookie string) {
	p.pool.Delete(cookie)
}

// Get returns operator's subscription object
func (p *PoolAnts) Get(cookie string) *SubscribeAnts {
	v, _ := p.pool.Load(cookie)
	return v
}

// Send sends proto message to subscribed operators
func (p *PoolAnts) Send(val any) {
	msg := &operatorv1.SubscribeAntsResponse{}

	switch val := val.(type) {
	case *operatorv1.AntResponse:
		msg.Response = &operatorv1.SubscribeAntsResponse_Ant{
			Ant: val,
		}
	case *operatorv1.AntColorResponse:
		msg.Response = &operatorv1.SubscribeAntsResponse_Color{
			Color: val,
		}
	case *operatorv1.AntNoteResponse:
		msg.Response = &operatorv1.SubscribeAntsResponse_Note{
			Note: val,
		}
	case *operatorv1.AntLastResponse:
		msg.Response = &operatorv1.SubscribeAntsResponse_Last{
			Last: val,
		}
	case *operatorv1.AntSleepResponse:
		msg.Response = &operatorv1.SubscribeAntsResponse_Sleep{
			Sleep: val,
		}
	case *operatorv1.AntsResponse:
		msg.Response = &operatorv1.SubscribeAntsResponse_Ants{
			Ants: val,
		}
	default:
		return
	}

	p.pool.Range(func(_ string, v *SubscribeAnts) bool {
		if err := v.stream.Send(msg); err != nil {
			v.err <- err
		}
		return true
	})
}

// Exists check if operator exists in pool by its username
func (p *PoolAnts) Exists(username string) bool {
	f := false
	p.pool.Range(func(_ string, v *SubscribeAnts) bool {
		if strings.Compare(v.username, username) == 0 {
			f = true
		}
		return !f
	})
	return f
}
