package models

import (
	"errors"

	"github.com/jezman/yapdd/pdd"
	"github.com/jezman/yapdd/utils"
	"github.com/levigross/grequests"
)

// Domain struct
type Domain struct {
	Accounts       []*Account `json:"accounts"`      // user account
	Aliases        []string   `json:"aliases"`       // list of domain aliases
	CheckResults   string     `json:"check_results"` // last check result
	Country        string     `json:"country"`       // interface language for mailboxes by default.
	Delegated      string     `json:"delegated"`     // status of the domain name delegation on the Yandex servers
	DKIM           DKIM
	Domain         string   `json:"domain"`           // domain name
	EmailsCount    int      `json:"emails-count"`     // number of available mailboxes
	EmailsMaxCount int      `json:"emails-max-count"` // maximum number of mailboxes that can be created for the domain
	Error          string   `json:"error"`            // error message
	ImapEnabled    int      `json:"imap_enabled"`     // imap status
	LastCheck      string   `json:"last_check"`       // date and time of last check
	Logo           bool     `json:"logo_enabled"`     // presence logo
	LogoURL        string   `json:"logo_url"`         // logo URL
	Name           string   `json:"name"`             // domain name
	NextCheck      string   `json:"next_check"`       // date and time of next check
	NoDKIM         bool     `json:"nodkim"`           // a sign that DKIM is not connected
	NSDelegated    bool     `json:"nsdelegated"`      // status of the domain name delegation on the Yandex servers
	PopEnabled     int      `json:"pop_enabled"`      // pop status
	Secrets        *Secrets `json:"secrets"`          // Secret data test file (or CNAME records)
	Status         string   `json:"status"`           // domain status
	Success        string   `json:"success"`          // request status
	Total          int      `json:"total"`            // total counts of user account
}

// Secrets data struct
type Secrets struct {
	Name    string `json:"name"`    // secret part of the actual file name (or CNAME records)
	Content string `json:"content"` // secret contents of the test file
}

// List accounts in domain.
func (d *Domain) List(domainName string) (*Domain, error) {
	ro.Params["on_page"] = "1000"
	ro.Params["domain"] = domainName

	response, err := grequests.Get(pdd.AccountsList, ro)
	if err != nil {
		return nil, err
	}
	if err := response.JSON(d); err != nil {
		return nil, err
	}
	return d, nil
}

// Add domain into Yandex PDD.
func (d *Domain) Add(domainName string) (*Domain, error) {
	ro.Params["domain"] = domainName

	response, err := grequests.Post(pdd.DomainAdd, ro)
	if err != nil {
		return nil, err
	}
	if err := response.JSON(d); err != nil {
		return nil, err
	}
	return d, nil
}

// ConnectionStatus gets domain connetion status.
func (d *Domain) ConnectionStatus(domainName string) (*Domain, error) {
	ro.Params["domain"] = domainName

	response, err := grequests.Get(pdd.DomainStatus, ro)
	if err != nil {
		return nil, err
	}
	if err := response.JSON(d); err != nil {
		return nil, err
	}
	return d, nil
}

// Config gets domain settings.
func (d *Domain) Config(domainName string) (*Domain, error) {
	ro.Params["domain"] = domainName

	response, err := grequests.Get(pdd.DomainConfig, ro)
	if err != nil {
		return nil, err
	}
	if err := response.JSON(d); err != nil {
		return nil, err
	}
	return d, nil
}

// Remove domain from YandexPDD.
func (d *Domain) Remove(domainName string) (*Domain, error) {
	// generates capcha for confirm remove
	capcha := utils.RandomInt(8)

	warning := "please confirm domain removed. input: " + capcha + "\n"
	// read user confirmation
	confirmation := utils.ReadStdIn(warning)

	// check confirmation
	if confirmation == capcha {
		ro.Params["domain"] = domainName

		// sends remove request
		response, err := grequests.Get(pdd.DomainDelete, ro)
		if err != nil {
			return nil, err
		}
		if err := response.JSON(d); err != nil {
			return nil, err
		}
		return d, nil
	}
	// wrong confirmation
	return nil, errors.New("confirmation error")
}
