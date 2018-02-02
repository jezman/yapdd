package pdd

import "os"

var (
	// Token API get from os envinronment
	Token = os.Getenv("PDD_TOKEN")
)

const (
	// BaseURL API version 2
	BaseURL = "https://pddimp.yandex.ru/api2"

	DomainsList      = BaseURL + "/admin/domain/domains?&on_page=1000"
	DomainAdd        = BaseURL + "/admin/domain/register"
	DomainStatus     = BaseURL + "/admin/domain/registration_status"
	DomainConfig     = BaseURL + "/admin/domain/details"
	DomainDelete     = BaseURL + "/admin/domain/delete"
	DomainSetCountry = BaseURL + "/admin/domain/settings/set_country"

	AccountsList        = BaseURL + "/admin/email/list?&on_page=1000&domain="
	AccountAdd          = BaseURL + "/admin/email/add"
	AccountDelete       = BaseURL + "api2/admin/email/del"
	AccountUnreadEmails = BaseURL + "/api2/admin/email/counters"
	AccountUpdate       = BaseURL + "/api2/admin/email/edit"
)
