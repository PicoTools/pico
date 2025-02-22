package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PicoTools/pico/cmd/pico-ctl/internal/service"
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/utils"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

func listenerListCommand(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:                   "list",
		Short:                 "list listeners",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, _ []string) {
			listeners, err := service.ListListeners()
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("list listeners: %s", err.Error())
				}
				return
			}
			if len(listeners) == 0 {
				color.Yellow("no listeners exists")
				return
			}
			for _, l := range listeners {
				fmt.Println(strings.Repeat("-", 42))
				fmt.Println(utils.PrettyListener(l))
			}
			fmt.Println(strings.Repeat("-", 42))
		},
	}
}

func listenerAddCommand(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:                   "add",
		Short:                 "add new listener",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			listener, err := service.AddListener()
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("add listener: %s", err.Error())
				}
				return
			}
			fmt.Println(utils.PrettyListener(listener))
		},
	}
}

func listenerRevokeCommand(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:                   "revoke",
		Short:                 "revoke listener's token",
		DisableFlagsInUseLine: true,
		Args:                  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				color.Red("invalid listener id")
				return
			}
			if err := service.RevokeListener(int64(id)); err != nil {
				switch status.Code(err) {
				default:
					color.Red("revoke listener: %s", err.Error())
				}
			}
			color.Green("token revoked for listener '%d'", id)
		},
	}
}

func listenerRegenCommand(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:                   "regen",
		Short:                 "regen listener's token",
		DisableFlagsInUseLine: true,
		Args:                  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				color.Red("invalid listener id")
				return
			}
			listener, err := service.RegenListener(int64(id))
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("regen listener: %s", err.Error())
				}
				return
			}
			fmt.Println(utils.PrettyListener(listener))
		},
	}
}

func listenerCommand(c *console.Console) *cobra.Command {
	listenerCmd := &cobra.Command{
		Use:                   "listener",
		Short:                 "manage listeners",
		GroupID:               listenerGroupId,
		DisableFlagsInUseLine: true,
	}

	listenerCmd.AddCommand(
		listenerListCommand(c),
		listenerAddCommand(c),
		listenerRevokeCommand(c),
		listenerRegenCommand(c),
	)

	return listenerCmd
}
