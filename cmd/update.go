package cmd

import (
	"github.com/jezman/yapdd/render"
	"github.com/spf13/cobra"
)

var (
	pass         string
	firstName    string
	lastName     string
	accoutStatus string
	birthday     string
	sex          string
	hintQuestion string
	hintAnswer   string
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update account information.",
	Args: cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		params := make(map[string]string, 8)
		params["password"] = pass
		params["iname"] = firstName
		params["fname"] = lastName
		params["enabled"] = accoutStatus
		params["birth_date"] = birthday
		params["sex"] = sex
		params["hintq"] = hintQuestion
		params["hinta"] = hintAnswer

		for k, v := range params {
			if v == "" {
				delete(params, k)
			}
		}

		render.UpdateAccount(args[0], params)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&pass, "pass", "p", "", "set user password")
	updateCmd.Flags().StringVarP(&firstName, "firstname", "f", "", "set user first name")
	updateCmd.Flags().StringVarP(&lastName, "lastname", "l", "", "set user last name")
	updateCmd.Flags().StringVarP(&accoutStatus, "enabled", "e", "", "yes or no to change account status")
	updateCmd.Flags().StringVarP(&birthday, "birth", "b", "", "set birthday YYYY-MM-DD")
	updateCmd.Flags().StringVarP(&sex, "sex", "s", "", "sex 0: undefined, 1: male, 2: female")
	updateCmd.Flags().StringVarP(&hintQuestion, "question", "q", "", "secret question")
	updateCmd.Flags().StringVarP(&hintAnswer, "answer", "a", "", "secret answer")
}
