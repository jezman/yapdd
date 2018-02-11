package render

import (
	"fmt"

	"github.com/apcera/termtables"
	"github.com/jezman/yapdd/models"
)

// Domains print domains list table
func Domains(verbose bool) {
	domains := &models.Domains{}
	list, err := domains.List(verbose)

	if err != nil {
		fmt.Println(err)
	} else if list.Success != "ok" {
		fmt.Printf("Status: %s\nError: %s\n", list.Success, list.Error)
	} else {
		table := termtables.CreateTable()
		table.AddTitle("List of user domains.")

		if verbose {
			table.AddHeaders("Domains",
				"Aliases",
				"Status",
				"NS Delegated",
				"No DKIM",
				"Accounts",
				"Max accounts",
			)

			for _, d := range list.Domains {
				table.AddRow(d.Name,
					d.Aliases,
					d.Status,
					d.NSDelegated,
					d.NoDKIM,
					d.EmailsCount,
					d.EmailsMaxCount,
				)
			}

		} else {
			table.AddHeaders("Domains", "Accounts")

			for _, d := range list.Domains {
				table.AddRow(d.Name, d.EmailsCount)
			}
		}
		fmt.Println(table.Render())
		fmt.Println("Total domains:", list.Total)
	}
}

// Accounts list in domain
func Accounts(domain string, verbose bool) {
	dmn := &models.Domain{}
	list, err := dmn.List(domain, verbose)

	if err != nil {
		fmt.Println(err)
	} else if list.Success != "ok" {
		fmt.Printf("Status: %s\nError: %s\n", list.Success, list.Error)
	} else {
		table := termtables.CreateTable()
		table.AddTitle("A list of accounts in the domain.")

		if verbose {
			table.AddHeaders("Account", "Active/Ready", "User name/Birthday", "Question hint")

			for _, a := range list.Accounts {
				table.AddRow(
					a.Login,
					a.Enabled+"/"+a.Ready,
					a.User+" "+a.Birthday,
					a.Question,
				)
			}
		} else {
			table.AddHeaders("Account")

			for _, a := range list.Accounts {
				table.AddRow(
					a.Login,
				)
			}
		}
		fmt.Println(table.Render())
		fmt.Printf("Total accounts in domain %s: %d\n", domain, list.Total)
	}
}

// CountOfUnreadMail rendered
func CountOfUnreadMail(account string) {
	a := &models.Account{}
	emailsCount, err := a.UnreadMail(account)

	if err != nil {
		fmt.Println(err)
	} else if emailsCount.Success != "ok" {
		fmt.Printf("Status: %s\nError: %s\n", emailsCount.Success, emailsCount.Error)
	} else {
		fmt.Println("Count of unread emails:", emailsCount.Counters.New)
		fmt.Println("Count of letters received since the last mailbox test:", emailsCount.Counters.Unread)
	}
}
