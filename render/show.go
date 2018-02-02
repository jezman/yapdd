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
