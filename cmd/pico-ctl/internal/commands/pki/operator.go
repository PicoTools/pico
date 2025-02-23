package pki

import (
	"fmt"

	"github.com/PicoTools/pico/cmd/pico-ctl/internal/service"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

// operatorCmd returns command "operator" for "pki"
func operatorCmd(c *console.Console) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "operator",
		Short: "Manage operator's TLS for gRPC",
	}

	cmd.AddCommand(
		// get operator's TLS certificate for gRPC
		getOperatorCertCmd(c),
	)

	return cmd
}

// getOperatorCertCmd returns command "get" for "operator"
func getOperatorCertCmd(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get operator's TLS certificate for gRPC",
		Run: func(cmd *cobra.Command, args []string) {
			// get operators' certificate
			cert, err := service.GetCertOperator()
			if err != nil {
				switch status.Code(err) {
				default:
					color.Red("Get operator certificate: %s", err.Error())
				}
				return
			}
			fmt.Print(cert.GetCertificate().GetData())
		},
	}
}
