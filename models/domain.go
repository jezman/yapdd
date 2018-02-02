package models

import (
	"encoding/json"

	"github.com/jezman/request"
	"github.com/jezman/yapdd/pdd"
)

// Domains struct
type Domains struct {
	Total   int       `json:"total"`   // total counts of user domains
	Domains []*Domain `json:"domains"` // list of domains
	Success string    `json:"success"` // request status
	Error   string    `json:"error"`   // error message
}

// Domain struct
type Domain struct {
	Domain         string     `json:"domain"`           // domain name
	Total          int        `json:"total"`            // total counts of user account
	Accounts       []*Account `json:"accounts"`         // user account
	Name           string     `json:"name"`             // domain name
	Status         string     `json:"status"`           // domain status
	Delegated      string     `json:"delegated"`        // status of the domain name delegation on the Yandex servers
	Country        string     `json:"country"`          // interface language for mailboxes by default.
	PopEnabled     int        `json:"pop_enabled"`      // pop status
	ImapEnabled    int        `json:"imap_enabled"`     // imap status
	Aliases        []string   `json:"aliases"`          // list of domain aliases
	Logo           bool       `json:"logo_enabled"`     // presence logo
	LogoURL        string     `json:"logo_url"`         // logo URL
	NSDelegated    bool       `json:"nsdelegated"`      // status of the domain name delegation on the Yandex servers
	EmailsMaxCount int        `json:"emails-max-count"` // maximum number of mailboxes that can be created for the domain
	EmailsCount    int        `json:"emails-count"`     // number of available mailboxes
	NoDKIM         bool       `json:"nodkim"`           // a sign that DKIM is not connected
	Success        string     `json:"success"`          // request status
	Secrets        *Secrets   `json:"secrets"`          // Secret data test file (or CNAME records)
	Error          string     `json:"error"`            // error message
	CheckResults   string     `json:"check_results"`    // last check result
	NextCheck      string     `json:"next_check"`       // date and time of next check
	LastCheck      string     `json:"last_check"`       // date and time of last check
}

// Secrets data struct
type Secrets struct {
	Name    string `json:"name"`    // secret part of the actual file name (or CNAME records)
	Content string `json:"content"` // secret contents of the test file
}

// List get list of user domains
func (d *Domains) List(verbose bool) (*Domains, error) {
	body, err := request.Get(pdd.DomainsList, request.Options{
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"PddToken":     pdd.Token,
		},
	})
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, d); err != nil {
		return nil, err
	}

	return d, nil
}

// List accounts in domain
func (d *Domain) List(domain string, verbose bool) (*Domain, error) {
	url := pdd.AccountsList + domain

	response, err := request.Get(url, request.Options{
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"PddToken":     pdd.Token,
		},
	})
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(response, d); err != nil {
		return nil, err
	}

	return d, nil
}
