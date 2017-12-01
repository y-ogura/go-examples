package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	sp "github.com/SparkPost/gosparkpost"
)

func main() {
	fmt.Println("vim-go")
	apiKey := os.Getenv("SPARKPOST_API_KEY")
	cfg := &sp.Config{
		BaseUrl:    "https://api.sparkpost.com",
		ApiKey:     apiKey,
		ApiVersion: 1,
	}
	var client sp.Client
	err := client.Init(cfg)
	if err != nil {
		log.Printf("SparkPost client init failed: %s", err)
		return
	}

	// excelファイルをbase64にエンコード
	file, _ := os.Open("./tmp/sample.xlsx")
	fileData, _ := ioutil.ReadAll(file)
	emcExcel := base64.StdEncoding.EncodeToString(fileData)
	from := os.Getenv("MAIL_SEND_ADDRESS")

	tx := &sp.Transmission{
		Recipients: []string{os.Getenv("MAIL_ADDRESS")},
		Content: sp.Content{
			HTML:    "<p>Hello</p>",
			From:    sp.From{Email: from, Name: "SCOREシステム"},
			Subject: "test mail",
			Attachments: []sp.Attachment{
				sp.Attachment{
					MIMEType: "application/ms-excel",
					Filename: "sample.xlsx",
					B64Data:  emcExcel,
				},
			},
		},
	}
	id, _, err := client.Send(tx)
	if err != nil {
		log.Printf("It failed in transmission of the mail: %s", err)
		return
	}
	log.Printf("Transmission sent with id [%s]\n", id)
}
