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

// List get list of user domains
func (d *Domains) List(verbose bool) (*Domains, error) {
	body, err := request.Get(pdd.DomainsList, request.Options{
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"PddToken":     pdd.Token,
		},
		Body: map[string]string{"on_page": "20"},
	})
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, d); err != nil {
		return nil, err
	}
	return d, nil
}
