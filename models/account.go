package models

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/jezman/request"
	"github.com/jezman/yapdd/pdd"
	"github.com/jezman/yapdd/utils"
)

// Account struct
type Account struct {
	Domain   string    `json:"domain"`     // domain name
	Aliases  []string  `json:"aliases"`    // email aliases
	Login    string    `json:"login"`      // email address
	UID      int       `json:"uid"`        // email id
	Enabled  string    `json:"enabled"`    // email account status
	User     string    `json:"fio"`        // full name
	FName    string    `json:"fname"`      // last name
	IName    string    `json:"iname"`      // first name
	Birthday string    `json:"birth_date"` // birthday YYYY-MM-DD
	Sex      int       `json:"sex"`        // 0 - not set; 1 - male; 2 - female
	Question string    `json:"hintq"`      // secret question
	Ready    string    `json:"ready"`      // ready to work
	MailList string    `json:"maillist"`   // email for newsletter
	Success  string    `json:"success"`    // request status
	Error    string    `json:"error"`      // error code
	Counters *Counters `json:"counters"`
}

// Counters of unread mails struct
type Counters struct {
	Unread int `json:"unread"`
	New    int `json:"new"`
}

// UnreadMail gets count of unread emails in account.
func (a *Account) UnreadMail(account string) (*Account, error) {
	tmp, err := utils.SplitAccount(account)
	if err != nil {
		return nil, err
	}

	accountName := tmp[0]
	domainName := tmp[1]

	response, err := request.Get(pdd.AccountUnreadEmails, request.Options{
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"PddToken":     pdd.Token,
		},
		Body: map[string]string{
			"domain": domainName,
			"login":  accountName,
		},
	})
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(response, a); err != nil {
		return nil, err
	}
	return a, nil
}

// Add account in domain.
func (a *Account) Add(account string) (*Account, error) {
	var password [2]string

	ask := utils.ReadStdIn("Generate password? (Yes): ")
	// generate password is answer "yes"
	if strings.ToLower(ask) == "yes" || strings.ToLower(ask) == "" {
		password[0] = utils.GeneratePassword(11)
		password[1] = password[0]
	// hand password input
	} else {
		// first password input
		password[0] = utils.ReadStdIn("Password: ")
		// confirmation password input
		password[1] = utils.ReadStdIn("Confirm password: ")
	}

	// check passwords match
	if password[0] == password[1] {
		tmp, err := utils.SplitAccount(account)
		if err != nil {
			return nil, errors.New("invalid email format")
		}

		accountName := tmp[0]
		domainName := tmp[1]

		// send request
		response, err := request.Post(pdd.AccountAdd, request.Options{
			Headers: map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
				"PddToken":     pdd.Token,
			},
			Body: map[string]string{
				"domain":   domainName,
				"login":    accountName,
				"password": password[1],
			},
		})
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(response, a); err != nil {
			return nil, err
		}
		return a, nil
	}
	
	return nil, errors.New("passwords don't match")
}

// Remove domain from YandexPDD.
func (a *Account) Remove(accountName string) (*Account, error) {
	// generates capcha for confirm remove
	capcha := utils.RandomInt(8)

	warning := "please confirm account removed. input: " + capcha + "\n"
	// read user confirmation
	confirmation := utils.ReadStdIn(warning)

	// check confirmation
	if confirmation == capcha {
		tmp, err := utils.SplitAccount(accountName)
		if err != nil {
			return nil, errors.New("invalid email format")
		}

		accountName := tmp[0]
		domainName := tmp[1]
		// sends remove request
		body, err := request.Post(pdd.AccountDelete, request.Options{
			Headers: map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
				"PddToken":     pdd.Token,
			},
			Body: map[string]string{
				"login":  accountName,
				"domain": domainName,
			},
		})
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(body, a); err != nil {
			return nil, err
		}
		return a, nil
	}

	// wrong confirmation
	return nil, errors.New("confirmation error")
}
