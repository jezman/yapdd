package models

import (
	"github.com/jezman/yapdd/pdd"
	"github.com/levigross/grequests"
)

// DNSRecords structure.
type DNSRecords struct {
	Records []*DNSRecord `json:"records"` // dns records
	Success string       `json:"success"` // request status
	Error   string       `json:"error"`   // error message
}

// DNSRecord structure.
type DNSRecord struct {
	Content   string `json:"content"`
	Domain    string `json:"domain"`
	FQDN      string `json:"fqdn"`
	RecordID  int64  `json:"record_id"`
	Subdomain string `json:"subdomain"`
	TTL       int    `json:"ttl"`
	Type      string `json:"type"`
}

// DNSRecords gets list of dns records in domain.
func (d *DNSRecords) DNSRecords(domainName string) (*DNSRecords, error) {
	ro.Params["domain"] = domainName

	response, err := grequests.Get(pdd.DNSList, ro)
	if err != nil {
		return nil, err
	}
	if err := response.JSON(d); err != nil {
		return nil, err
	}
	return d, nil
}
