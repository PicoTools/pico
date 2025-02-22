package commands

import (
	"os"

	"github.com/PicoTools/pico/cmd/pico-ctl/internal/service"
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/utils"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

func exitCommand(c *console.Console) *cobra.Command {
	return &cobra.Command{
		Use:     "exit",
		Short:   "exit management cli",
		GroupID: globalGroupId,
		Run: func(*cobra.Command, []string) {
			if utils.ExitConsole(c) {
				service.Close()
				os.Exit(0)
			}
		},
	}
}
