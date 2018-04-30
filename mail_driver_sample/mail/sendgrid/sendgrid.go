package sendgrid

import (
	"fmt"
	"os"

	mailContent "github.com/go-examples/mail_driver_sample/mail/content"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Sendgrid sendgrid struct
type Sendgrid struct{}

// Send send email
func (s *Sendgrid) Send(content mailContent.Content) error {
	from := mail.NewEmail("NUTS.SHOP", os.Getenv("SEND_MAIL_ADDRESS"))
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	message := mailInit(from, content)
	res, err := client.Send(message)
	if err != nil {
		return err
	}
	fmt.Println(res.StatusCode)
	fmt.Println(res.Body)
	return nil
}

func mailInit(from *mail.Email, content mailContent.Content) *mail.SGMailV3 {
	m := new(mail.SGMailV3)
	m.SetFrom(from)
	m.Subject = content.Subject
	p := mail.NewPersonalization()
	for _, to := range content.To {
		p.AddTos(mail.NewEmail("", to))
	}
	m.AddPersonalizations(p)
	c := mail.NewContent("text/html", content.Message)
	m.AddContent(c)
	return m
}
