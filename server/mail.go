package main

import (
	"fmt"
	"strconv"
	"gopkg.in/mailgun/mailgun-go.v1"
)

var shouldMail bool = func() bool {
	innerBool, parseError := strconv.ParseBool(environment["SHOULD_MAIL"])
	if parseError != nil {
		panic(parseError)
	}

	return innerBool
}()

var domain string = environment["SERVER_DOMAIN"]
var privateAPIKey string = environment["MAIL_PRIVATE_API_KEY"]
var publicValidationKey string = environment["MAIL_PUBLIC_KEY"]

var mailgunClient mailgun.Mailgun = mailgun.NewMailgun(domain, privateAPIKey, publicValidationKey)

func sendMessage(sender string, subject string, body string, recipient string) error {
	message := mailgunClient.NewMessage(sender, subject, body, recipient)
	if shouldMail {
		_, _, err := mailgunClient.Send(message)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("fake emailed message: ")
		fmt.Println(message)
	}

	return nil
}
