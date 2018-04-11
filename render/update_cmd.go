package render

import (
	"fmt"

	"github.com/jezman/yapdd/utils"
)

// UpdateAccount print results
func UpdateAccount(accountName string, params map[string]string) {
	account, err := account.Update(accountName, params)

	if err = utils.ErrorCheck(account.Success, account.Error, err); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Account '%s' successfully updated.\n", account.Login)
	}
}
