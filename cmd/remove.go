package cmd

import (
	"github.com/jezman/yapdd/render"
	"github.com/jezman/yapdd/utils"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove domains from YandexPDD and accounts from domains.",
	Long: `Command "remove" - Remove domains from YandexPDD and accounts from domains.
Need confirmation for removal.

Example:
  yapdd remove example.com        Removes domain "example.com" from YandexPDD.
  yapdd remove toremove@example.com    Remove account "toremove" from domain "example.com".`,
	Args: cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
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
