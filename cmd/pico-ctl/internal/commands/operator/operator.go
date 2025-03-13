package operator

import (
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/constants"
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/service"
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/utils"
	"github.com/PicoTools/pico/pkg/shared"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Cmd returns command "operators"
func Cmd(c *console.Console) *cobra.Command {
	operatorCmd := &cobra.Command{
		Use:     "operators",
		Short:   "Manage operators",
		GroupID: constants.BaseGroupId,
	}

	operatorCmd.AddCommand(
		// list registered operators
		listCmd(c),
		// register new operator
		addCmd(c),
		// revoke access token for operator
		revokeCmd(c),
		// regenerate access token for operator
		regenerateCmd(c),
	)

	return operatorCmd
}

// listCmd returns command "list" for "operators"
func listCmd(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List registered operators",
		Run: func(cmd *cobra.Command, _ []string) {
			// get list of operators
			operators, err := service.ListOperators()
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("List operators: %s", err.Error())
				}
				return
			}
			if len(operators) == 0 {
				color.Yellow("No operators registered yet")
				return
			}
			utils.TableOperators(operators)
		},
	}
}

// addCmd returns command "add" for "operators"
func addCmd(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Register new operator",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// parse input
			username := args[0]
			if len(username) < shared.OperatorUsernameMinLength || len(username) > shared.OperatorUsernameMaxLength {
				color.Red("Invalid username's length")
				return
			}
			// register operator
			operator, err := service.AddOperator(username)
			if err != nil {
				switch status.Code(err) {
				case codes.AlreadyExists:
					color.Red("Operator with such username already exists")
				default:
					color.Red("Register new operator: %s", err.Error())
				}
				return
			}
			utils.TableOperator(operator)
		},
	}
}

// revokeCmd returns command "revoke" for "operators"
func revokeCmd(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "revoke",
		Short: "Revoke access token for operator",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// parse input
			username := args[0]
			if len(username) < shared.OperatorUsernameMinLength || len(username) > shared.OperatorUsernameMaxLength {
				color.Red("Invalid username's length")
				return
			}
			// revoke access token
			if err := service.RevokeTokenOperator(username); err != nil {
				switch status.Code(err) {
				default:
					color.Red("Revoke access token: %s", err.Error())
				}
			}
			color.Green("Access token revoked for '%s'", username)
		},
	}
}

// regenerateCmd returns command "regenerate" for "operators"
func regenerateCmd(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "regenerate",
		Short: "Regenerate access token for operator",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// parse input
			username := args[0]
			if len(username) < shared.OperatorUsernameMinLength || len(username) > shared.OperatorUsernameMaxLength {
				color.Red("Invalid username's length")
				return
			}
			// regenerate access token
			operator, err := service.RegenerateTokenOperator(username)
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("Regenerate access token: %s", err.Error())
				}
				return
			}
			utils.TableOperator(operator)
		},
	}
}
