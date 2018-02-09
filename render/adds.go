package render

import (
	"fmt"
	"github.com/jezman/yapdd/models"
)
// AddDomain to Yandex PDD
func AddDomain(domainName string) {
	domain := &models.Domain{}
	list, err := domain.Add(domainName)

	if err != nil {
		fmt.Println(err)
	} else if list.Success != "ok" {
		fmt.Printf("Error: %s\n", list.Error)
	} else {
		fmt.Printf("Domain %s has been successfully connected.", list.Domain)
	}
}
