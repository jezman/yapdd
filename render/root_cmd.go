package render

import (
	"fmt"
	"strings"

	"github.com/apcera/termtables"
	"github.com/jezman/yapdd/models"
	"github.com/jezman/yapdd/utils"
)

var (
	account = &models.Account{}
	dns     = &models.DNSRecords{}
	domain  = &models.Domain{}
	domains = &models.Domains{}
)

// Domains render domains list table.
func Domains(verbose bool) {
	domains, err := domains.List()

	if err = utils.ErrorCheck(domains.Success, domains.Error, err); err != nil {
		fmt.Println(err)
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

			for i, d := range domains.Domains {
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

			for i, d := range domains.Domains {
				table.AddRow(i+1, d.Name, d.EmailsCount)
			}
		}
		fmt.Println(table.Render())
	}
}

// Accounts render accounts list in domain.
func Accounts(domainName string, verbose bool) {
	domain, err := domain.List(domainName)

	if err = utils.ErrorCheck(domain.Success, domain.Error, err); err != nil {
		fmt.Println(err)
	} else {
		table := termtables.CreateTable()
		table.AddTitle("A list of accounts in the domain.")

		// vebose output
		if verbose {
			table.AddHeaders("#", "Account", "Active/Ready", "Username/Birthday", "Question hint")

			for i, a := range domain.Accounts {
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

			for i, a := range domain.Accounts {
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
	account, err := account.UnreadMail(accountName)

	if err = utils.ErrorCheck(account.Success, account.Error, err); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Count of unread emails:", account.Counters.New)
		fmt.Println("Count of letters received since the last mailbox test:",
			 account.Counters.Unread,
			)
	}
}

// DomainStatus render connection status.
func DomainStatus(domainName string) {
	domain, err := domain.ConnectionStatus(domainName)

	if err = utils.ErrorCheck(domain.Success, domain.Error, err); err != nil {
		fmt.Println(err)
	} else {
		table := termtables.CreateTable()
		table.AddTitle("Domain connection status.")

		table.AddHeaders("Domain", "Status", "Check results", "Last check", "Next check")

		table.AddRow(
			domain.Domain,
			domain.Status,
			domain.CheckResults,
			domain.LastCheck,
			domain.NextCheck,
		)
		fmt.Println(table.Render())
	}
}

// DomainConfig render domain settings.
func DomainConfig(domainName string) {
	domain, err := domain.Config(domainName)

	if err = utils.ErrorCheck(domain.Success, domain.Error, err); err != nil {
		fmt.Println(err)
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
			domain.Domain,
			domain.Status,
			domain.Delegated,
			domain.Country,
			domain.ImapEnabled,
			domain.PopEnabled,
		)
		fmt.Println(table.Render())
	}
}

// DomainDNSRecords print DNS records in domain.
func DomainDNSRecords(domainName string) {
	dns, err := dns.DNSRecords(domainName)

	if err = utils.ErrorCheck(dns.Success, dns.Error, err); err != nil {
		fmt.Println(err)
	} else {
		for _, d := range dns.Records {
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
