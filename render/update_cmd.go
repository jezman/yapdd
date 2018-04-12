package render

import (
	"fmt"

	"github.com/jezman/yapdd/utils"
)

// EnableAccount in domain
func EnableAccount(accountName string) {
	account, err := account.Enable(accountName)

	if err = utils.ErrorCheck(account.Success, account.Error, err); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s - enabled.\n", account.Login)
	}
}

// DisableAccount in domain
func DisableAccount(accountName string) {
	account, err := account.Disable(accountName)

	if err = utils.ErrorCheck(account.Success, account.Error, err); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s - disabled.\n", account.Login)
	}
}

// UpdateAccount print results
func UpdateAccount(accountName string, params map[string]string) {
	account, err := account.Update(accountName, params)

	if err = utils.ErrorCheck(account.Success, account.Error, err); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Account '%s' successfully updated.\n", account.Login)
	}
}
