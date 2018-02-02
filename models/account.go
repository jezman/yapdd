package models

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
