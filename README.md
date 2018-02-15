# YaPDD
Command line application for administration Yandex PDD.
## Installing

Set environment variable **PDD_TOKEN**:
```bash
export PDD_TOKEN="yandex_pdd_api_token"
```
Install Gorion
If you have [Go](https://golang.org/) installed: 
```bash
go get github.com/jezman/yapdd && go install github.com/jezman/yapdd
```
Otherwise, please see [Go install](https://golang.org/doc/install).
## Features

- List of user domains.
- Domains management
  * Add domain
  * Remove domain
  * Get connection status
  * Get domain settings

- Accounts management
  * Add account
  * Remove account
  * Get count of unread emails

- DKIM managements
  * Get informations
  * Enable/Disable

## TODO
- set language for domain
- domain logo managements
- update account settings