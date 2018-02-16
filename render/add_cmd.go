package render

import (
	"fmt"

	"github.com/jezman/yapdd/utils"
)

// AddDomain print domain add result.
func AddDomain(domainName string) {
	domain, err := domain.Add(domainName)

	if err = utils.ErrorCheck(domain.Success, domain.Error, err); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Domain %s has been successfully connected.", domain.Domain)
	}
}

// AddAccount print account add result.
func AddAccount(accountName string) {
	account, err := account.Add(accountName)

	if err = utils.ErrorCheck(account.Success, account.Error, err); err != nil {
		fmt.Println(err)
		fmt.Printf("Account '%s' was successfully added.\n", account.Login)
	}
}
