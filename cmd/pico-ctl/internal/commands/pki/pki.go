package pki

import (
	"github.com/PicoTools/pico/cmd/pico-ctl/internal/constants"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

// Cmd returns command "pki"
func Cmd(c *console.Console) *cobra.Command {
	pkiCmd := &cobra.Command{
		Use:     "pki",
		Short:   "Manage PKI for gRPC TLS",
		GroupID: constants.BaseGroupId,
	}

	pkiCmd.AddCommand(
		// manage CA for gRPC
		caCmd(c),
		// get operator's TLS certificate
		operatorCmd(c),
		// get listener's TLS certificate
		listenerCmd(c),
	)

	return pkiCmd
}
