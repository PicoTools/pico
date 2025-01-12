package cfg

import (
	"context"
)

// configCtxKey is the context's key to store config object in
type configCtxKey struct{}

// SetConfigCtx saves config object to context
func SetConfigCtx(ctx context.Context, c Config) context.Context {
	return context.WithValue(ctx, configCtxKey{}, c)
}

// GetConfigCtx gathers config object from context
func GetConfigCtx(ctx context.Context) Config {
	return ctx.Value(configCtxKey{}).(Config)
}
