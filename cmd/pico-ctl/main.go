package main

import (
	"context"
	"os"
	"os/signal"
	"slices"

	"github.com/PicoTools/pico/cmd/pico-ctl/internal/cmd"
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/service"
	"github.com/PicoTools/pico/internal/zapcfg"
	"github.com/fatih/color"
	"github.com/go-faster/sdk/zctx"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func main() {
	lg, err := zapcfg.New().Build()
	if err != nil {
		panic(err)
	}

	flush := func() {
		_ = lg.Sync()
	}
	defer flush()

	exit := func(code int) {
		flush()
		os.Exit(code)
	}

	defer func() {
		if r := recover(); r != nil {
			lg.Fatal("recovered from panic", zap.Any("panic", r))
			exit(2)
		}
	}()

	app := cmd.App{}
	ctx, cancel := signal.NotifyContext(zctx.Base(context.Background(), lg), os.Interrupt)
	defer cancel()

	root := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
		Use:           "pico-ctl",
		Short:         "pico management cli",
		Long:          "pico management cli",
		Args:          cobra.NoArgs,
		RunE:          app.Run,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if !slices.Contains([]string{
				"help",
			}, cmd.Name()) {
				if err = app.Validate(); err != nil {
					return err
				}
				if err = service.Init(cmd.Context(), app.Host, app.Token); err != nil {
					return err
				}
			}
			return nil
		},
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			flush()
		},
	}

	root.CompletionOptions.DisableDefaultCmd = true
	app.RegisterFlags(root.PersistentFlags())

	if err = root.ExecuteContext(ctx); err != nil {
		color.Red("%v", err)
		exit(2)
	}
}
