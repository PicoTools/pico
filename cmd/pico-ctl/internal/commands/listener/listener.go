package listener

import (
	"strconv"

	"github.com/PicoTools/pico/cmd/pico-ctl/internal/constants"
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/service"
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/utils"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

// Cmd returns command "listeners"
func Cmd(c *console.Console) *cobra.Command {
	listenerCmd := &cobra.Command{
		Use:     "listeners",
		Short:   "Manage listeners",
		GroupID: constants.BaseGroupId,
	}

	listenerCmd.AddCommand(
		// list registered listener
		listCmd(c),
		// add new listener
		addCmd(c),
		// revoke access token for listener
		revokeCmd(c),
		// regenerate access token for listener
		regenerateCmd(c),
	)

	return listenerCmd
}

// listCmd returns command "list" for "listeners"
func listCmd(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List registered listeners",
		Run: func(cmd *cobra.Command, _ []string) {
			// get list of listeners
			listeners, err := service.ListListeners()
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("List listeners: %s", err.Error())
				}
				return
			}
			if len(listeners) == 0 {
				color.Yellow("No listeners registered yet")
				return
			}
			utils.TableListeners(listeners)
		},
	}
}

// addCmd returns command "add" for "listeners"
func addCmd(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Register new listener",
		Run: func(cmd *cobra.Command, args []string) {
			// add listener
			listener, err := service.AddListener()
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("Register new listener: %s", err.Error())
				}
				return
			}
			utils.TableListener(listener)
		},
	}
}

// revokeCmd returns command "revoke" for "listeners"
func revokeCmd(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "revoke",
		Short: "Revoke access token for listener",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// parse input
			id, err := strconv.Atoi(args[0])
			if err != nil {
				color.Red("Invalid listener's ID")
				return
			}
			// revoke access token for listener
			if err := service.RevokeTokenListener(int64(id)); err != nil {
				switch status.Code(err) {
				default:
					color.Red("Revoke access token: %s", err.Error())
				}
			}
			color.Green("Access token revoked for listener '%d'", id)
		},
	}
}

// regenerateCmd returns command "regenerate" for "listeners"
func regenerateCmd(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "regenerate",
		Short: "Regenerate access token for listener",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// parse input
			id, err := strconv.Atoi(args[0])
			if err != nil {
				color.Red("Invalid listener's ID")
				return
			}
			// regenerate token
			listener, err := service.RegenerateTokenListener(int64(id))
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("Regenerate access token: %s", err.Error())
				}
				return
			}
			utils.TableListener(listener)
		},
	}
}
