package cmd

import (
	"context"
	"fmt"

	"github.com/PicoTools/pico/internal/cfg"
	"github.com/PicoTools/pico/internal/constants"
	"github.com/PicoTools/pico/internal/utils"

	"github.com/creasty/defaults"
	"github.com/go-faster/errors"
	"github.com/spf13/pflag"
)

type App struct {
	Globals Globals
}

type Globals struct {
	Config string
	Debug  bool
}

// RegisterFlags registers global flags
func (g *Globals) RegisterFlags(f *pflag.FlagSet) {
	f.StringVarP(&g.Config, "config", "c", utils.EnvOr(constants.ConfigPathEnvKey, ""),
		fmt.Sprintf("path to config file [env:%s]", constants.ConfigPathEnvKey))
	f.BoolVarP(&g.Debug, "debug", "d", false, "enable debug logging")
}

// Validate validates config specified by path
func (g *Globals) Validate(ctx context.Context) (context.Context, error) {
	// read config file
	c, err := cfg.Read(g.Config)
	if err != nil {
		return nil, errors.Wrap(err, "read config file")
	}
	// set up defaults
	if err = defaults.Set(&c); err != nil {
		return nil, errors.Wrap(err, "set defaults")
	}
	// validate config
	if err = c.Validate(ctx); err != nil {
		return nil, errors.Wrap(err, "validate config")
	}
	// inject config object in context
	return cfg.SetConfigCtx(ctx, c), nil
}
