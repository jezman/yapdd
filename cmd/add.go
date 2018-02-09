package cmd

import (
	"github.com/jezman/yapdd/render"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Added domains, accounts to Yandex PDD.",
	Run: func(cmd *cobra.Command, args []string) {
		render.AddDomain(domain)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&domain, "domain", "d", "", "add domain to Yandex PDD.")
	addCmd.Flags().StringVarP(&account, "account", "a", "", "add account to domain.")
}
