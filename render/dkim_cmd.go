package render

import (
	"fmt"

	"github.com/jezman/yapdd/models"
)

// DKIMStatus render DKIM informations.
func DKIMStatus(domainName string) {
	d := &models.Domain{}
	json, err := d.DKIMStatus(domainName)

	if err != nil {
		fmt.Println(err)
	} else if json.Success != "ok" {
		fmt.Printf("Error: %s\n", json.Error)
	} else {
		fmt.Println("Domain:", json.Domain)
		fmt.Println("Status:", json.DKIM.Status)
		fmt.Println("TXT record:", json.DKIM.TxtRecord)
		fmt.Println("Mail ready:", json.DKIM.MailReady)
		fmt.Println("Secret key:", json.DKIM.SecretKey)
	}
}

// DKIMEnable print result.
func DKIMEnable(domainName string) {
	d := &models.Domain{}
	json, err := d.DKIMEnable(domainName)

	if err != nil {
		fmt.Println(err)
	} else if json.Success != "ok" {
		fmt.Printf("Error: %s\n", json.Error)
	} else {
		fmt.Println("Domain:", json.Domain)
		fmt.Println("Status:", json.DKIM.Status)
		fmt.Println("TXT record:", json.DKIM.TxtRecord)
	}
}

// DKIMDisable print result.
func DKIMDisable(domainName string) {
	d := &models.Domain{}
	json, err := d.DKIMDisable(domainName)

	if err != nil {
		fmt.Println(err)
	} else if json.Success != "ok" {
		fmt.Printf("Error: %s\n", json.Error)
	} else {
		fmt.Println("Domain:", json.Domain)
		fmt.Println("Status:", json.Success)
	}
}
