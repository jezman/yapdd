package pdd

import "os"

var (
	// Token API get from os envinronment
	Token = os.Getenv("PDD_TOKEN")
)

const (
	// BaseURL API version 2
	BaseURL = "https://pddimp.yandex.ru/api2"

	DomainsList      = BaseURL + "/admin/domain/domains"
	DomainAdd        = BaseURL + "/admin/domain/register"
	DomainStatus     = BaseURL + "/admin/domain/registration_status"
	DomainConfig     = BaseURL + "/admin/domain/details"
	DomainDelete     = BaseURL + "/admin/domain/delete"
	DomainSetCountry = BaseURL + "/admin/domain/settings/set_country"

	AccountsList        = BaseURL + "/admin/email/list"
	AccountAdd          = BaseURL + "/admin/email/add"
	AccountDelete       = BaseURL + "/admin/email/del"
	AccountUnreadEmails = BaseURL + "/admin/email/counters"
	AccountUpdate       = BaseURL + "/admin/email/edit"

	DNSList = BaseURL + "/admin/dns/list"
)
