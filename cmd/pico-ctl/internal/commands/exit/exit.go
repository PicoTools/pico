package exit

import (
	"os"

	"github.com/PicoTools/pico/cmd/pico-ctl/internal/constants"
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/service"
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/utils"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

// Cmd returns command "exit"
func Cmd(c *console.Console) *cobra.Command {
	return &cobra.Command{
		Use:     "exit",
		Short:   "Exit management's CLI",
		GroupID: constants.BaseGroupId,
		Run: func(*cobra.Command, []string) {
			if utils.ExitConsole(c) {
				service.Close()
				os.Exit(0)
			}
		},
	}
}
