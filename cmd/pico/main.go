package main

import (
	"context"
	"os"
	"os/signal"
	"slices"

	"github.com/PicoTools/pico/cmd/pico/internal/cmd"
	"github.com/PicoTools/pico/cmd/pico/internal/cmd/run"
	"github.com/PicoTools/pico/cmd/pico/internal/cmd/version"

	_ "github.com/PicoTools/pico/internal/ent/runtime"
	"github.com/PicoTools/pico/internal/zapcfg"

	"github.com/fatih/color"
	"github.com/go-faster/errors"
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
		// ignore: /dev/stderr: invalid argument
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

	a := cmd.App{}
	ctx, cancel := signal.NotifyContext(zctx.Base(context.Background(), lg), os.Interrupt)
	defer cancel()

	root := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,

		Use:   "pico",
		Short: "pico C2",
		Long:  "pico C2 server",
		Args:  cobra.NoArgs,

		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			cmdCtx := cmd.Context()

			if !slices.Contains([]string{
				"version",
				"help",
			}, cmd.Name()) {
				cmdCtx, err = a.Globals.Validate(cmdCtx)
				if err != nil {
					return err
				}
			}
			if a.Globals.Debug {
				zapcfg.AtomLvl.SetLevel(zap.DebugLevel)
			}
			cmd.SetContext(cmdCtx)
			return nil
		},
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			flush()
		},
	}

	root.CompletionOptions.DisableDefaultCmd = true
	a.Globals.RegisterFlags(root.PersistentFlags())
	root.AddCommand(
		version.Command(),
		run.Command(),
	)

	if err := root.ExecuteContext(ctx); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			err = context.Canceled
		}
		color.Red("%v", err)
		exit(2)
	}
}
