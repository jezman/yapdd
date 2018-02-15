package cmd

import (
	"github.com/jezman/yapdd/render"
	"github.com/spf13/cobra"
)

var (
	enable  bool
	disable bool
)

// dkimCmd represents the dkim command
var dkimCmd = &cobra.Command{
	Use:   "dkim",
	Short: "DKIM managements.",
	Long: `Command "dkim" - managements DKIM.

Example:
  yapdd dkim example.com        DKIM information for domain "example.com".
  yapdd dkim -e example.com     Enable DKIM for domain "example.com".
  yapdd dkim -d example.com     Disable DKIM for domain "example.com".`,
	Args: cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case enable:
			render.DKIMEnable(args[0])
		case disable:
			render.DKIMDisable(args[0])
		default:
			render.DKIMStatus(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(dkimCmd)

	dkimCmd.Flags().BoolVarP(&enable, "enable", "e", false, "enable dkim for domain")
	dkimCmd.Flags().BoolVarP(&disable, "disable", "d", false, "disable dkim for domain")
}
