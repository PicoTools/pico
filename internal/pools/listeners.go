package pools

import (
	"strings"

	operatorv1 "github.com/PicoTools/pico-shared/proto/gen/operator/v1"

	"github.com/lrita/cmap"
)

type SubscribeListeners struct {
	username   string
	disconnect chan struct{}
	data       chan any
	err        chan error
	stream     operatorv1.OperatorService_SubscribeListenersServer
}

type PoolListeners struct {
	pool cmap.Map[string, *SubscribeListeners]
}

// Error returns error appeared in process of stream handling
func (s *SubscribeListeners) Error() chan error {
	return s.err
}

// IsDisconnect notify if operator must be disconnected
func (s *SubscribeListeners) IsDisconnect() chan struct{} {
	return s.disconnect
}

// Add adds operator's stream to pool
func (p *PoolListeners) Add(cookie, username string, stream operatorv1.OperatorService_SubscribeListenersServer) {
	p.pool.Store(cookie, &SubscribeListeners{
		username:   username,
		disconnect: make(chan struct{}, 1),
		data:       make(chan any, 1),
		err:        make(chan error),
		stream:     stream,
	})
}

// Disconnect notify about necessity of operator's disconnection
func (p *PoolListeners) Disconnect(cookie string) {
	v, ok := p.pool.Load(cookie)
	if ok {
		v.disconnect <- struct{}{}
	}
}

// Remove removes operator's stream from pool
func (p *PoolListeners) Remove(cookie string) {
	p.pool.Delete(cookie)
}

// Get returns operator's subscription object
func (p *PoolListeners) Get(cookie string) *SubscribeListeners {
	l, _ := p.pool.Load(cookie)
	return l
}

// Send sends proto message to subscribed operators
func (p *PoolListeners) Send(val any) {
	msg := &operatorv1.SubscribeListenersResponse{}

	switch val := val.(type) {
	case *operatorv1.ListenerResponse:
		msg.Response = &operatorv1.SubscribeListenersResponse_Listener{
			Listener: val,
		}
	case *operatorv1.ListenerColorResponse:
		msg.Response = &operatorv1.SubscribeListenersResponse_Color{
			Color: val,
		}
	case *operatorv1.ListenerNoteResponse:
		msg.Response = &operatorv1.SubscribeListenersResponse_Note{
			Note: val,
		}
	case *operatorv1.ListenerInfoResponse:
		msg.Response = &operatorv1.SubscribeListenersResponse_Info{
			Info: val,
		}
	case *operatorv1.ListenerLastResponse:
		msg.Response = &operatorv1.SubscribeListenersResponse_Last{
			Last: val,
		}
	case *operatorv1.ListenersResponse:
		msg.Response = &operatorv1.SubscribeListenersResponse_Listeners{
			Listeners: val,
		}
	default:
		return
	}

	p.pool.Range(func(_ string, v *SubscribeListeners) bool {
		if err := v.stream.Send(msg); err != nil {
			v.err <- err
		}
		return true
	})
}

// Exists check if operator exists in pool by its username
func (p *PoolListeners) Exists(username string) bool {
	f := false
	p.pool.Range(func(_ string, v *SubscribeListeners) bool {
		if strings.Compare(v.username, username) == 0 {
			f = true
		}
		return !f
	})
	return f
}
