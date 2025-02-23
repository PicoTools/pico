package pki

import (
	"fmt"

	"github.com/PicoTools/pico/cmd/pico-ctl/internal/service"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

// listenerCmd returns command "listener" for "pki"
func listenerCmd(c *console.Console) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listener",
		Short: "Manage listener's TLS for gRPC",
	}

	cmd.AddCommand(
		// get listener's TLS certificate for gRPC
		getListenerCertCmd(c),
	)

	return cmd
}

// getListenerCertCmd returns command "get" for "listener"
func getListenerCertCmd(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get listener's TLS certificate for gRPC",
		Run: func(cmd *cobra.Command, args []string) {
			// get listener's certificate
			cert, err := service.GetCertListener()
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("Get listener certificate: %s", err.Error())
				}
				return
			}
			fmt.Print(cert.GetCertificate().GetData())
		},
	}
}
