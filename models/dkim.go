package models

import (
	"encoding/json"

	"github.com/jezman/request"
	"github.com/jezman/yapdd/pdd"
)

// DKIM structure
type DKIM struct {
	Status    string `json:"enabled"`   // dkim connection status
	TxtRecord string `json:"txtrecord"` // TXT record
	NSReady   string `json:"nsready"`   // presence of the TXT record
	MailReady string // YandexPDD readiness to sign letters according to DKIM
	SecretKey string // DKIM secret key
}

// DKIMStatus gets DKIM informations.
func (d *Domain) DKIMStatus(domainName string) (*Domain, error) {
	body, err := request.Get(pdd.DKIMStatus, request.Options{
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

func (d *Domain) DKIMEnable(domainName string) (*Domain, error) {
	body, err := request.Get(pdd.DKIMEnable, request.Options{
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

func (d *Domain) DKIMDisable(domainName string) (*Domain, error) {
	body, err := request.Get(pdd.DKIMDisable, request.Options{
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
