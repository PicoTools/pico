package run

import (
	"github.com/PicoTools/pico/internal/cfg"
	"github.com/PicoTools/pico/internal/ent"
	"github.com/PicoTools/pico/internal/listener"
	"github.com/PicoTools/pico/internal/management"
	"github.com/PicoTools/pico/internal/operator"
	"golang.org/x/sync/errgroup"

	"github.com/go-faster/sdk/zctx"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// holds objects needing for C2 server's bootstrap
type run struct {
	lg         *zap.Logger
	db         *ent.Client
	operator   operator.Config
	listener   listener.Config
	management management.Config
}

// Command provides cobra's command to run C2 server
func Command() *cobra.Command {
	c := run{}
	return &cobra.Command{
		Use:               "run",
		Short:             "run pico server",
		ValidArgsFunction: cobra.NoFileCompletions,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			var err error

			ctx := cmd.Context()
			c.lg = zctx.From(ctx).Named("run")
			ctxCfg := cfg.GetConfigCtx(ctx)

			// initialize DB
			if c.db, err = ctxCfg.Db.Init(ctx); err != nil {
				return err
			}
			// initialize PKI for GRPC servers
			if err = ctxCfg.Pki.Init(ctx, c.db, ctxCfg.Listener.IP, ctxCfg.Operator.IP, ctxCfg.Management.IP); err != nil {
				return err
			}
			// save configs
			c.operator = ctxCfg.Operator
			c.listener = ctxCfg.Listener
			c.management = ctxCfg.Management

			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			g, ctx := errgroup.WithContext(ctx)

			// management GRPC
			g.Go(func() error {
				return management.Serve(ctx, c.management, c.db)
			})

			// listener GRPC
			g.Go(func() error {
				return listener.Serve(ctx, c.listener, c.db)
			})

			// operator GRPC
			g.Go(func() error {
				return operator.Serve(ctx, c.operator, c.db)
			})

			return g.Wait()
		},
		PostRun: func(cmd *cobra.Command, _ []string) {
			if err := c.db.Close(); err != nil {
				c.lg.Error("closing storage", zap.Error(err))
			}
			c.lg.Debug("storage closed")
		},
	}
}
