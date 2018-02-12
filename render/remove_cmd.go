package render

import (
	"fmt"

	"github.com/jezman/yapdd/models"
)

func RemoveDomain(domainName string) {
	domain := &models.Domain{}
	json, err := domain.Remove(domainName)

	if err != nil {
		fmt.Println(err)
	} else if json.Success != "ok" {
		fmt.Printf("Error: %s\n", json.Error)
	} else {
		fmt.Printf("Domain '%s' has been successfully removed.\n", json.Domain)
	}

}

func RemoveAccount(accountName string) {
	account := &models.Account{}
	json, err := account.Remove(accountName)

	if err != nil {
		fmt.Println(err)
	} else if json.Success != "ok" {
		fmt.Printf("Error: %s\n", json.Error)
	} else {
		fmt.Printf("Domain '%s' has been successfully removed.\n", json.Login)
	}

}
