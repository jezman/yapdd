package cmd

import (
	"github.com/jezman/yapdd/render"
	"github.com/jezman/yapdd/utils"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Add accounts.
		if utils.IsAccount(args[0]) {
			render.RemoveAccount(args[0])
		} else {
			render.RemoveDomain(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
