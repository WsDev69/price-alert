package models

type EmailNotification struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}
