package cmd

import (
	"github.com/jezman/yapdd/render"
	"github.com/jezman/yapdd/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add domains, accounts to Yandex PDD.",
	Long: `Command "add" - Add domains, accounts to Yandex PDD.

Example:
  yapdd add example.com        Add domain "example.com".
  yapdd add acc@example.com    Adding account into domain "example.com".`,
	Args: cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Add accounts.
		if utils.IsAccount(args[0]) {
			render.AddAccount(args[0])
		} else {
			render.AddDomain(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
