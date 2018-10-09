# YaPDD
Command line application for administration Yandex PDD.
## Installing

- get [PDD-token](https://pddimp.yandex.ru/api2/admin/get_token)

- set environment variable **PDD_TOKEN**:
```bash
export PDD_TOKEN="yandex_pdd_api_token"
```
- install Gorion
If you have [Go](https://golang.org/) installed: 
```bash
go get github.com/jezman/yapdd && go install github.com/jezman/yapdd
```
- otherwise, please see [Go install](https://golang.org/doc/install).
## Features

- List of user domains.
- List of accounts in domain.
- Domains management
  * Add domain
  * Remove domain
  * Get connection status
  * Get domain settings

- Accounts management
  * Add account
  * Remove account
  * Get count of unread emails

- Update account settings
  * password
  * first name
  * last name
  * lock/unlock
  * birthday
  * sex
  * secret question
  * secret answer

- DKIM managements
  * Get informations
  * Enable/Disable

- Set language for domain

## License
MIT © 2017 jezman
