package pools

import (
	"strings"

	operatorv1 "github.com/PicoTools/pico/pkg/proto/operator/v1"

	"github.com/lrita/cmap"
)

type SubscribeChat struct {
	username   string
	disconnect chan struct{}
	data       chan any
	err        chan error
	stream     operatorv1.OperatorService_SubscribeChatServer
}

type PoolChat struct {
	pool cmap.Map[string, *SubscribeChat]
}

// Error returns error appeared in process of stream handling
func (s *SubscribeChat) Error() chan error {
	return s.err
}

// IsDisconnect notify if operator must be disconnected
func (s *SubscribeChat) IsDisconnect() chan struct{} {
	return s.disconnect
}

// Add adds operator's stream to pool
func (p *PoolChat) Add(cookie, username string, stream operatorv1.OperatorService_SubscribeChatServer) {
	p.pool.Store(cookie, &SubscribeChat{
		username:   username,
		disconnect: make(chan struct{}, 1),
		data:       make(chan any, 1),
		err:        make(chan error),
		stream:     stream,
	})
}

// Disconnect notify about necessity of operator's disconnection
func (p *PoolChat) Disconnect(cookie string) {
	v, ok := p.pool.Load(cookie)
	if ok {
		v.disconnect <- struct{}{}
	}
}

// Remove removes operator's stream from pool
func (p *PoolChat) Remove(cookie string) {
	p.pool.Delete(cookie)
}

// Get returns operator's subscription object
func (p *PoolChat) Get(cookie string) *SubscribeChat {
	v, _ := p.pool.Load(cookie)
	return v
}

// Send sends proto message to subscribed operators
func (p *PoolChat) Send(val any) {
	msg := &operatorv1.SubscribeChatResponse{}

	switch val := val.(type) {
	case *operatorv1.ChatResponse:
		msg.Response = &operatorv1.SubscribeChatResponse_Message{
			Message: val,
		}
	case *operatorv1.ChatMessagesResponse:
		msg.Response = &operatorv1.SubscribeChatResponse_Messages{
			Messages: val,
		}
	default:
		return
	}

	p.pool.Range(func(_ string, v *SubscribeChat) bool {
		if err := v.stream.Send(msg); err != nil {
			v.err <- err
		}
		return true
	})
}

// Exists check if operator exists in pool by its username
func (p *PoolChat) Exists(username string) bool {
	f := false
	p.pool.Range(func(_ string, v *SubscribeChat) bool {
		if strings.Compare(v.username, username) == 0 {
			f = true
		}
		return !f
	})
	return f
}
