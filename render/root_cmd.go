package render

import (
	"fmt"
	"strings"

	"github.com/apcera/termtables"
	"github.com/jezman/yapdd/models"
)

// Domains render domains list table.
func Domains(verbose bool) {
	domains := &models.Domains{}
	json, err := domains.List(verbose)

	if err != nil {
		fmt.Println(err)
	} else if json.Success != "ok" {
		fmt.Println("Error:", json.Error)
	} else {
		table := termtables.CreateTable()
		table.AddTitle("List of user domains.")

		// verbose output
		if verbose {
			table.AddHeaders(
				"#",
				"Domains",
				"Aliases",
				"Status",
				"NS Delegated",
				"No DKIM",
				"Accounts",
				"Max accounts",
			)

			for i, d := range json.Domains {
				table.AddRow(
					i+1,
					d.Name,
					d.Aliases,
					d.Status,
					d.NSDelegated,
					d.NoDKIM,
					d.EmailsCount,
					d.EmailsMaxCount,
				)
			}

			// quiet output
		} else {
			table.AddHeaders("#", "Domains", "Accounts")

			for i, d := range json.Domains {
				table.AddRow(i+1, d.Name, d.EmailsCount)
			}
		}
		fmt.Println(table.Render())
	}
}

// Accounts render accounts list in domain.
func Accounts(domainName string, verbose bool) {
	domain := &models.Domain{}
	json, err := domain.List(domainName, verbose)

	if err != nil {
		fmt.Println(err)
	} else if json.Success != "ok" {
		fmt.Println("Error:", json.Error)
	} else {
		table := termtables.CreateTable()
		table.AddTitle("A list of accounts in the domain.")

		// vebose output
		if verbose {
			table.AddHeaders("#", "Account", "Active/Ready", "Username/Birthday", "Question hint")

			for i, a := range json.Accounts {
				table.AddRow(
					i+1,
					a.Login,
					a.Enabled+"/"+a.Ready,
					a.User+" "+a.Birthday,
					a.Question,
				)
			}
			// quiet output
		} else {
			table.AddHeaders("#", "Account")

			for i, a := range json.Accounts {
				table.AddRow(
					i+1,
					a.Login,
				)
			}
		}
		fmt.Println(table.Render())
	}
}

// CountOfUnreadMail in account.
func CountOfUnreadMail(accountName string) {
	account := &models.Account{}
	json, err := account.UnreadMail(accountName)

	if err != nil {
		fmt.Println(err)
	} else if json.Success != "ok" {
		fmt.Println("Error:", json.Error)
	} else {
		fmt.Println("Count of unread emails:", json.Counters.New)
		fmt.Println("Count of letters received since the last mailbox test:", json.Counters.Unread)
	}
}

// DomainStatus render connection status.
func DomainStatus(domainName string) {
	domain := &models.Domain{}
	json, err := domain.ConnectionStatus(domainName)

	if err != nil {
		fmt.Println(err)
	} else if json.Success != "ok" {
		fmt.Println("Error:", json.Error)
	} else {
		table := termtables.CreateTable()
		table.AddTitle("Domain connection status.")

		table.AddHeaders("Domain", "Status", "Check results", "Last check", "Next check")

		table.AddRow(
			json.Domain,
			json.Status,
			json.CheckResults,
			json.LastCheck,
			json.NextCheck,
		)
		fmt.Println(table.Render())
	}
}

// DomainConfig render domain settings.
func DomainConfig(domainName string) {
	domain := &models.Domain{}
	json, err := domain.Config(domainName)

	if err != nil {
		fmt.Println(err)
	} else if json.Success != "ok" {
		fmt.Println("Error:", json.Error)
	} else {
		table := termtables.CreateTable()
		table.AddTitle("Domain settings.")

		table.AddHeaders(
			"Domain",
			"Status",
			"Delegated",
			"Country",
			"IMAP",
			"POP",
		)

		table.AddRow(
			json.Domain,
			json.Status,
			json.Delegated,
			json.Country,
			json.ImapEnabled,
			json.PopEnabled,
		)
		fmt.Println(table.Render())
	}
}

// DomainDNSRecords print DNS records in domain.
func DomainDNSRecords(domainName string) {
	domain := &models.DNSRecords{}
	json, err := domain.DNSRecords(domainName)

	if err != nil {
		fmt.Println(err)
	} else if json.Success != "ok" {
		fmt.Println("Error:", json.Error)
	} else {
		for _, d := range json.Records {
			fmt.Printf("%s\n", strings.Repeat("-", 50))
			fmt.Printf("%s\t| ", d.Type)
			fmt.Println("Record ID:", d.RecordID)
			fmt.Println("\t| Full domain name:", d.FQDN)
			fmt.Println("\t| TTL:", d.TTL)
			fmt.Println("\t| Subdomain:", d.Subdomain)
			fmt.Println("\t| Content:", d.Content)
		}
	}
}
