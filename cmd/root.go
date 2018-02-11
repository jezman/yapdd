package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jezman/yapdd/render"
	"github.com/spf13/cobra"
)

var (
	verbose bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yapdd",
	Short: "Command line application for administration Yandex PDD.",
	Long: `YaPDD - Command line application for administration Yandex PDD.

Example:
  yapdd                     Domains list.
  yapdd example.com         List of accounts in domain.
  yapdd acc@example.com     Count of unread emails in account.`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			render.Domains(verbose)
		} else {
			if strings.Contains(args[0], "@") {
				render.CountOfUnreadMail(args[0])
			} else {
				render.Accounts(args[0], verbose)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	// rootCmd.Flags().BoolVarP(&status, "status", "s", false, "show status")
}
