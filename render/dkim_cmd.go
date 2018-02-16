package render

import (
	"fmt"

	"github.com/jezman/yapdd/utils"
)

// DKIMStatus render DKIM informations.
func DKIMStatus(domainName string) {
	domain, err := domain.DKIMStatus(domainName)

	if err = utils.ErrorCheck(domain.Success, domain.Error, err); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Domain:", domain.Domain)
		fmt.Println("Status:", domain.DKIM.Status)
		fmt.Println("TXT record:", domain.DKIM.TxtRecord)
		fmt.Println("Mail ready:", domain.DKIM.MailReady)
		fmt.Println("Secret key:", domain.DKIM.SecretKey)
	}
}

// DKIMEnable print result.
func DKIMEnable(domainName string) {
	domain, err := domain.DKIMEnable(domainName)

	if err = utils.ErrorCheck(domain.Success, domain.Error, err); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Domain:", domain.Domain)
		fmt.Println("Status:", domain.DKIM.Status)
		fmt.Println("TXT record:", domain.DKIM.TxtRecord)
	}
}

// DKIMDisable print result.
func DKIMDisable(domainName string) {
	domain, err := domain.DKIMDisable(domainName)

	if err = utils.ErrorCheck(domain.Success, domain.Error, err); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Domain:", domain.Domain)
		fmt.Println("Status:", domain.Success)
	}
}
