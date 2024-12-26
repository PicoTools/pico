package pools

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// pool to handle operators'es subscriptions
var Pool pool

type pool struct {
	Hello       PoolHello
	Chat        PoolChat
	Listeners   PoolListeners
	Ants        PoolAnts
	Operators   PoolOperators
	Credentials PoolCredentials
	Tasks       PoolTasks
}

// DisconnectAll disconnects operator by its cookie from all topics
func (p *pool) DisconnectAll(ctx context.Context, cookie string) {
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		p.Chat.Disconnect(cookie)
		return nil
	})
	g.Go(func() error {
		p.Listeners.Disconnect(cookie)
		return nil
	})
	g.Go(func() error {
		p.Ants.Disconnect(cookie)
		return nil
	})
	g.Go(func() error {
		p.Operators.Disconnect(cookie)
		return nil
	})
	g.Go(func() error {
		p.Credentials.Disconnect(cookie)
		return nil
	})
	g.Go(func() error {
		p.Tasks.Disconnect(cookie)
		return nil
	})
	_ = g.Wait()
}
