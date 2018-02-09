package cmd

import (
	"github.com/jezman/yapdd/render"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Domains and accounts list. Count of unread emails in account.",
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case domain != "":
			render.Accounts(domain, verbose)
		case account != "":
			render.CountOfUnreadMail(account)
		default:
			render.Domains(verbose)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().StringVarP(&domain, "domain", "d", "", "list of accounts by domain.")
	showCmd.Flags().StringVarP(&account, "account", "a", "", "count of unread emails in account.")
}
