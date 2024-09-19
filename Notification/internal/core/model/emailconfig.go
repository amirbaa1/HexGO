package model

type EmailConfig struct {
	UserName        string `json:"username"`
	Password        string `json:"password"`
	Server          string `json:"server"`
	To              string `json:"to"`
	Subject         string `json:"subject"`
	From            string `json:"from"`
	BodyHtml        string `json:"bodyHtml"`
	BodyText        string `json:"bodyText"`
	IsTransactional bool   `json:"isTransactional"`
}
