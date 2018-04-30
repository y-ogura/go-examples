package mail

import (
	"os"

	"github.com/go-examples/mail_driver_sample/mail/content"
	"github.com/go-examples/mail_driver_sample/mail/sendgrid"
	"github.com/go-examples/mail_driver_sample/mail/smtp"
	sp "github.com/go-examples/mail_driver_sample/mail/sparkpost"
)

// Mail mail interface
type Mail interface {
	Send(content content.Content) error
}

// New mount mail driver
func New() Mail {
	driver := os.Getenv("MAIL_DRIVER")
	switch driver {
	case "sparlpost":
		return &sp.Sparkpost{}
	case "sendgrid":
		return &sendgrid.Sendgrid{}
	default:
		return &smtp.SMTP{}
	}
}
