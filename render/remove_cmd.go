package render

import (
	"fmt"

	"github.com/jezman/yapdd/utils"
)

// RemoveDomain print result of domain removed.
func RemoveDomain(domainName string) {
	domain, err := domain.Remove(domainName)

	if err = utils.ErrorCheck(domain.Success, domain.Error, err); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Domain '%s' has been successfully removed.\n", domain.Domain)
	}

}

// RemoveAccount print result of account removed.
func RemoveAccount(accountName string) {
	account, err := account.Remove(accountName)

	if err = utils.ErrorCheck(account.Success, account.Error, err); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Account '%s' has been successfully removed.\n", account.Login)
	}

}
