package pools

import (
	"strings"

	operatorv1 "github.com/PicoTools/pico/pkg/proto/operator/v1"
	"github.com/lrita/cmap"
)

type SubscribeCredentials struct {
	username   string
	disconnect chan struct{}
	data       chan any
	err        chan error
	stream     operatorv1.OperatorService_SubscribeCredentialsServer
}

type PoolCredentials struct {
	pool cmap.Map[string, *SubscribeCredentials]
}

// Error returns error appeared in process of stream handling
func (s *SubscribeCredentials) Error() chan error {
	return s.err
}

// IsDisconnect notify if operator must be disconnected
func (s *SubscribeCredentials) IsDisconnect() chan struct{} {
	return s.disconnect
}

// Add adds operator's stream to pool
func (p *PoolCredentials) Add(cookie, username string, stream operatorv1.OperatorService_SubscribeCredentialsServer) {
	p.pool.Store(cookie, &SubscribeCredentials{
		username:   username,
		disconnect: make(chan struct{}, 1),
		data:       make(chan any, 1),
		err:        make(chan error),
		stream:     stream,
	})
}

// Disconnect notify about necessity of operator's disconnection
func (p *PoolCredentials) Disconnect(cookie string) {
	v, ok := p.pool.Load(cookie)
	if ok {
		v.disconnect <- struct{}{}
	}
}

// Remove removes operator's stream from pool
func (p *PoolCredentials) Remove(cookie string) {
	p.pool.Delete(cookie)
}

// Get returns operator's subscription object
func (p *PoolCredentials) Get(cookie string) *SubscribeCredentials {
	v, _ := p.pool.Load(cookie)
	return v
}

// Send sends proto message to subscribed operators
func (p *PoolCredentials) Send(val any) {
	msg := &operatorv1.SubscribeCredentialsResponse{}

	switch val := val.(type) {
	case *operatorv1.CredentialResponse:
		msg.Response = &operatorv1.SubscribeCredentialsResponse_Credential{
			Credential: val,
		}
	case *operatorv1.CredentialColorResponse:
		msg.Response = &operatorv1.SubscribeCredentialsResponse_Color{
			Color: val,
		}
	case *operatorv1.CredentialNoteResponse:
		msg.Response = &operatorv1.SubscribeCredentialsResponse_Note{
			Note: val,
		}
	case *operatorv1.CredentialsResponse:
		msg.Response = &operatorv1.SubscribeCredentialsResponse_Credentials{
			Credentials: val,
		}
	default:
		return
	}

	p.pool.Range(func(_ string, v *SubscribeCredentials) bool {
		if err := v.stream.Send(msg); err != nil {
			v.err <- err
		}
		return true
	})
}

// Exists check if operator exists in pool by its username
func (p *PoolCredentials) Exists(username string) bool {
	f := false
	p.pool.Range(func(_ string, v *SubscribeCredentials) bool {
		if strings.Compare(v.username, username) == 0 {
			f = true
		}
		return !f
	})
	return f
}
