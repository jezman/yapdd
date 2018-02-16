package models

import (
	"github.com/jezman/yapdd/pdd"
	"github.com/levigross/grequests"
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
	ro.Params["domain"] = domainName

	response, err := grequests.Get(pdd.DKIMStatus, ro)
	if err != nil {
		return nil, err
	}
	if err := response.JSON(d); err != nil {
		return nil, err
	}
	return d, nil
}

// DKIMEnable for domain.
func (d *Domain) DKIMEnable(domainName string) (*Domain, error) {
	ro.Params["domain"] = domainName

	response, err := grequests.Post(pdd.DKIMEnable, ro)
	if err != nil {
		return nil, err
	}
	if err := response.JSON(d); err != nil {
		return nil, err
	}
	return d, nil
}

// DKIMDisable for domain.
func (d *Domain) DKIMDisable(domainName string) (*Domain, error) {
	ro.Params["domain"] = domainName

	response, err := grequests.Post(pdd.DKIMDisable, ro)
	if err != nil {
		return nil, err
	}
	if err := response.JSON(d); err != nil {
		return nil, err
	}
	return d, nil
}
