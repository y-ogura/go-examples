package smtp

import (
	"net/smtp"
	"os"

	mailContent "github.com/go-examples/mail_driver_sample/mail/content"
)

// SMTP smtp struct
type SMTP struct{}

// Send send email
func (s *SMTP) Send(content mailContent.Content) error {
	// set config values
	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT")
	user := os.Getenv("MAIL_USERNAME")
	pass := os.Getenv("MAIL_PASSWORD")
	from := os.Getenv("MAIL_SEND_ADDRESS")

	// set massage
	var msg string
	msg = msg + "From:" + from + "\r\n"
	for _, to := range content.To {
		msg = msg + "To:" + to + "\r\n"
	}
	msg = msg + "Subject:" + content.Subject + "\r\n\r\n"
	msg = msg + content.Message

	// Set up authentication information.
	auth := smtp.PlainAuth("", user, pass, host)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(host+":"+port, auth, from, content.To, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}
