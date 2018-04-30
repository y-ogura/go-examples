package sparkpost

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	sp "github.com/SparkPost/gosparkpost"
	mailContent "github.com/go-examples/mail_driver_sample/mail/content"
)

// Sparkpost sparkpost
type Sparkpost struct{}

// Send send email
func (s *Sparkpost) Send(content mailContent.Content) error {
	// Get our API key from the environment; configure.
	apiKey := os.Getenv("SPARKPOST_API_KEY")
	cfg := &sp.Config{
		BaseUrl:    "https://api.sparkpost.com",
		ApiKey:     apiKey,
		ApiVersion: 1,
	}

	var client sp.Client
	err := client.Init(cfg)
	if err != nil {
		return fmt.Errorf("SparkPost client init failed: %s", err)
	}

	attachments, err := b64EncodeAttachementFiles(content.Attachments)
	if err != nil {
		return fmt.Errorf("Can't parse attachment files: %s", err)
	}

	from := sp.From{
		Email: os.Getenv("MAIL_SEND_ADDRESS"),
		Name:  "NUTS.SHOP",
	}

	tx := &sp.Transmission{
		Recipients: content.To,
		Content: sp.Content{
			HTML:        content.Message,
			From:        from,
			Subject:     content.Subject,
			Attachments: attachments,
		},
	}
	id, _, err := client.Send(tx)
	if err != nil {
		return fmt.Errorf("It failed in transmission of the mail: %s", err)
	}

	log.Printf("Transmission sent with id [%s]\n", id)

	return nil
}

// b64EncodeAttachementFiles 添付ファイルデータをBase64にエンコード
func b64EncodeAttachementFiles(attachments []mailContent.Attachment) ([]sp.Attachment, error) {
	var res = []sp.Attachment{}
	for _, attachment := range attachments {
		var data = sp.Attachment{
			MIMEType: attachment.Type,
			Filename: attachment.FileName,
		}
		fileData, err := ioutil.ReadAll(attachment.Data)
		if err != nil {
			return res, err
		}
		fileEnc := base64.StdEncoding.EncodeToString(fileData)
		data.B64Data = fileEnc
		res = append(res, data)
	}
	return res, nil
}
