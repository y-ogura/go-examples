package main

import (
	"fmt"

	"github.com/go-examples/mail_driver_sample/mail"
	"github.com/go-examples/mail_driver_sample/mail/content"
)

func main() {
	mail := mail.New()
	c := content.Content{
		To:      []string{"y-ogura@example.com"},
		Subject: "Mail driver sample",
		Message: "test message",
	}
	err := mail.Send(c)
	if err != nil {
		fmt.Println(err.Error())
	}
}
