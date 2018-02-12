package render

import (
	"fmt"

	"github.com/jezman/yapdd/models"
)

// AddDomain to Yandex PDD
func AddDomain(domainName string) {
	domain := &models.Domain{}
	json, err := domain.Add(domainName)

	if err != nil {
		fmt.Println(err)
	} else if json.Success != "ok" {
		fmt.Printf("Error: %s\n", json.Error)
	} else {
		fmt.Printf("Domain %s has been successfully connected.", json.Domain)
	}
}

// AddAccount to Yandex PDD
func AddAccount(accountName string) {
	account := &models.Account{}
	json, err := account.Add(accountName)

	if err != nil {
		fmt.Println(err)
	} else if json.Success != "ok" {
		fmt.Printf("Error: %s\n", json.Error)
	} else {
		fmt.Printf("Account '%s' was successfully added.\n", json.Login)
	}
}
