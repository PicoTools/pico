package pki

import (
	"fmt"

	"github.com/PicoTools/pico/cmd/pico-ctl/internal/service"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

// caCmd returns command "ca" for "pki"
func caCmd(c *console.Console) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ca",
		Short: "Manage CA's TLS for gRPC",
	}

	cmd.AddCommand(
		// get CA's TLS certificate for gRPC
		getCaCertCmd(c),
	)

	return cmd
}

// getCaCertCmd returns comand "get" for "ca"
func getCaCertCmd(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get CA's TLS certificate for gRPC",
		Run: func(cmd *cobra.Command, args []string) {
			// get CA certificate
			cert, err := service.GetCertCA()
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("Get CA certificate: %s", err.Error())
				}
				return
			}
			fmt.Print(cert.GetCertificate().GetData())
		},
	}
}
