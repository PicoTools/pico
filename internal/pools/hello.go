package pools

import (
	"strings"

	operatorv1 "github.com/PicoTools/pico-shared/proto/gen/operator/v1"

	"github.com/lrita/cmap"
)

type subscribeHello struct {
	username string
	stream   operatorv1.OperatorService_HelloServer
}

type PoolHello struct {
	pool cmap.Map[string, *subscribeHello]
}

// Add adds operator's stream to pool
func (p *PoolHello) Add(cookie, username string, stream operatorv1.OperatorService_HelloServer) {
	p.pool.Store(cookie, &subscribeHello{
		username: username,
		stream:   stream,
	})
}

// Remove removes operator's stream from pool
func (p *PoolHello) Remove(cookie string) {
	p.pool.Delete(cookie)
}

// Exists check if operator exists in pool by its username
func (p *PoolHello) Exists(username string) bool {
	f := false
	p.pool.Range(func(c string, v *subscribeHello) bool {
		if strings.Compare(v.username, username) == 0 {
			f = true
			return false
		}
		return true
	})
	return f
}

// Validate checks if supplied by operator cookie is valid and stores in pool
func (p *PoolHello) Validate(username, cookie string) bool {
	if v, ok := p.pool.Load(cookie); ok {
		if strings.Compare(v.username, username) == 0 {
			return true
		}
	}
	return false
}
