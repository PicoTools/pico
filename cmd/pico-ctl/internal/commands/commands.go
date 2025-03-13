package commands

import (
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/commands/exit"
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/commands/listener"
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/commands/operator"
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/commands/pki"
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/constants"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

func Commands(app *console.Console) console.Commands {
	return func() *cobra.Command {
		cmd := &cobra.Command{
			DisableFlagsInUseLine: true,
			SilenceErrors:         true,
			SilenceUsage:          true,
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
		}

		cmd.AddGroup(
			&cobra.Group{ID: constants.BaseGroupId, Title: constants.BaseGroupId},
		)

		// exit from shell
		cmd.AddCommand(exit.Cmd(app))
		// manage operators
		cmd.AddCommand(operator.Cmd(app))
		// manage listeners
		cmd.AddCommand(listener.Cmd(app))
		// manage pki
		cmd.AddCommand(pki.Cmd(app))

		// initialize help (will be interactive)
		cmd.InitDefaultHelpCmd()

		return cmd
	}
}
