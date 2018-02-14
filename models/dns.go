package models

import (
	"encoding/json"

	"github.com/jezman/request"
	"github.com/jezman/yapdd/pdd"
)

type DNSRecords struct {
	Records []*DNSRecord `json:"records"` // dns records
	Success string       `json:"success"` // request status
	Error   string       `json:"error"`   // error message
}

type DNSRecord struct {
	RecordID  int64  `json:"record_id"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	Domain    string `json:"domain"`
	FQDN      string `json:"fqdn"`
	TTL       int    `json:"ttl"`
	Subdomain string `json:"subdomain"`
}

func (d *DNSRecords) DNSRecords(domainName string) (*DNSRecords, error) {
	body, err := request.Get(pdd.DNSList, request.Options{
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
