package models

import (
	"errors"
	"strings"

	"github.com/jezman/yapdd/pdd"
	"github.com/jezman/yapdd/utils"
	"github.com/levigross/grequests"
)

// Account struct
type Account struct {
	Aliases  []string  `json:"aliases"`    // email aliases
	Birthday string    `json:"birth_date"` // birthday YYYY-MM-DD
	Counters *Counters `json:"counters"`
	Domain   string    `json:"domain"`   // domain name
	Enabled  string    `json:"enabled"`  // email account status
	Error    string    `json:"error"`    // error code
	FName    string    `json:"fname"`    // last name
	IName    string    `json:"iname"`    // first name
	Login    string    `json:"login"`    // email address
	MailList string    `json:"maillist"` // email for newsletter
	Question string    `json:"hintq"`    // secret question
	Ready    string    `json:"ready"`    // ready to work
	Sex      int       `json:"sex"`      // 0 - not set; 1 - male; 2 - female
	Success  string    `json:"success"`  // request status
	UID      int       `json:"uid"`      // email id
	User     string    `json:"fio"`      // full name
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

	ro.Params["login"] = tmp[0]
	ro.Params["domain"] = tmp[1]

	response, err := grequests.Get(pdd.AccountUnreadEmails, ro)
	if err != nil {
		return nil, err
	}
	if err := response.JSON(a); err != nil {
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

		ro.Params["login"] = tmp[0]
		ro.Params["domain"] = tmp[1]
		ro.Params["password"] = password[1]

		// send request
		response, err := grequests.Post(pdd.AccountAdd, ro)
		if err != nil {
			return nil, err
		}
		if err := response.JSON(a); err != nil {
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

		ro.Params["login"] = tmp[0]
		ro.Params["domain"] = tmp[1]
		// sends remove request

		response, err := grequests.Post(pdd.AccountAdd, ro)
		if err != nil {
			return nil, err
		}
		if err := response.JSON(a); err != nil {
			return nil, err
		}
		return a, nil
	}
	// wrong confirmation
	return nil, errors.New("confirmation error")
}
