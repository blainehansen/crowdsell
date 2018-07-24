package main

import (
	"gopkg.in/mailgun/mailgun-go.v1"
)

var domain string = environment["SERVER_DOMAIN"]
var privateAPIKey string = environment["MAIL_PRIVATE_API_KEY"]
var publicValidationKey string = environment["MAIL_PUBLIC_KEY"]

var mailgunClient mailgun.Mailgun = mailgun.NewMailgun(domain, privateAPIKey, publicValidationKey)

func sendMessage(sender string, subject string, body string, recipient string) error {
	message := mailgunClient.NewMessage(sender, subject, body, recipient)
	// resp, id, err := mailgunClient.Send(message)
	_, _, err := mailgunClient.Send(message)

	if err != nil {
		return err
	}

	return nil
}
