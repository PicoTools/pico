package console

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/PicoTools/pico/cmd/pico-ctl/internal/commands"
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/service"
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/utils"
	"github.com/fatih/color"
	"github.com/reeflective/console"
)

func Run(ctx context.Context) error {
	app := console.New("pico-ctl")
	main := app.ActiveMenu()
	main.Short = "management commands"
	main.Prompt().Primary = func() string {
		return fmt.Sprintf("%s > ", color.New(color.FgHiCyan).Add(color.Underline).Sprint("pico-ctl"))
	}
	main.SetCommands(commands.Commands(app))
	main.AddInterrupt(io.EOF, func(c *console.Console) {
		if utils.ExitConsole(c) {
			service.Close()
			os.Exit(0)
		}
	})
	return app.StartContext(ctx)
}
