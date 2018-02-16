package models

import (
	"github.com/jezman/yapdd/pdd"
	"github.com/jezman/yapdd/utils"
	"github.com/levigross/grequests"
)

var ro = utils.RequestOptions

// Domains struct
type Domains struct {
	Total   int       `json:"total"`   // total counts of user domains
	Domains []*Domain `json:"domains"` // list of domains
	Success string    `json:"success"` // request status
	Error   string    `json:"error"`   // error message
}

// List gets list of user domains.
func (d *Domains) List() (*Domains, error) {
	ro.Params["on_page"] = "20"

	response, err := grequests.Get(pdd.DomainsList, ro)
	if err != nil {
		return nil, err
	}

	if err := response.JSON(d); err != nil {
		return nil, err
	}
	return d, nil
}
