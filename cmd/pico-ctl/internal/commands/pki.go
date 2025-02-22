package commands

import (
	"fmt"

	"github.com/PicoTools/pico/cmd/pico-ctl/internal/service"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

func pkiCertCACommand(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:                   "ca",
		Short:                 "get CA certificate",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			cert, err := service.GetCertCA()
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("get CA cert: %s", err.Error())
				}
				return
			}
			fmt.Println(cert.GetCertificate().GetData())
		},
	}
}

func pkiCertOperatorCommand(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:                   "operator",
		Short:                 "get operator's GRPC certificate",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			cert, err := service.GetCertOperator()
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("get operator cert: %s", err.Error())
				}
				return
			}
			fmt.Println(cert.GetCertificate().GetData())
		},
	}
}

func pkiCertListenerCommand(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:                   "listener",
		Short:                 "get listener's GRPC certificate",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			cert, err := service.GetCertListener()
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("get listener cert: %s", err.Error())
				}
				return
			}
			fmt.Println(cert.GetCertificate().GetData())
		},
	}
}

func pkiCommand(c *console.Console) *cobra.Command {
	pkiCmd := &cobra.Command{
		Use:                   "pki",
		Short:                 "manage pki",
		GroupID:               pkiGroupId,
		DisableFlagsInUseLine: true,
	}

	pkiCmd.AddCommand(
		pkiCertCACommand(c),
		pkiCertOperatorCommand(c),
		pkiCertListenerCommand(c),
	)

	return pkiCmd
}
