package version

import (
	"fmt"

	"github.com/PicoTools/pico/internal/version"
	"github.com/spf13/cobra"
)

// Command provides cobra's command to print version of software
func Command() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Show version information",
		Run: func(*cobra.Command, []string) {
			fmt.Println(version.Get().PrettyColorful())
		},
	}
}
