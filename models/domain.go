package models

import (
	"encoding/json"
	"errors"

	"github.com/jezman/request"
	"github.com/jezman/yapdd/pdd"
	"github.com/jezman/yapdd/utils"
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
func (d *Domain) List(domain string, verbose bool) (*Domain, error) {
	response, err := request.Get(pdd.AccountsList, request.Options{
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"PddToken":     pdd.Token,
		},
		Body: map[string]string{
			"on_page": "500",
			"domain":  domain,
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

// Add domain into Yandex PDD.
func (d *Domain) Add(domain string) (*Domain, error) {
	body, err := request.Post(pdd.DomainAdd, request.Options{
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"PddToken":     pdd.Token,
		},
		Body: map[string]string{
			"domain": domain,
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

// ConnectionStatus gets domain connetion status.
func (d *Domain) ConnectionStatus(domain string) (*Domain, error) {
	body, err := request.Get(pdd.DomainStatus, request.Options{
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"PddToken":     pdd.Token,
		},
		Body: map[string]string{
			"domain": domain,
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

// Config gets domain settings.
func (d *Domain) Config(domain string) (*Domain, error) {
	body, err := request.Get(pdd.DomainConfig, request.Options{
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"PddToken":     pdd.Token,
		},
		Body: map[string]string{
			"domain": domain,
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

// Remove domain from YandexPDD.
func (d *Domain) Remove(domainName string) (*Domain, error) {
	// generates capcha for confirm remove
	capcha := utils.RandomInt(8)

	warning := "please confirm domain removed. input: " + capcha + "\n"
	// read user confirmation
	confirmation := utils.ReadStdIn(warning)

	// check confirmation
	if confirmation == capcha {
		// sends remove request
		body, err := request.Post(pdd.DomainDelete, request.Options{
			Headers: map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
				"PddToken":     pdd.Token,
			},
			Body: map[string]string{
				"domain": domainName,
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

	// wrong confirmation
	return nil, errors.New("confirmation error")
}
