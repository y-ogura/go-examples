package content

import "os"

// Content mail content struct
type Content struct {
	To          []string
	Subject     string
	Message     string
	Attachments []Attachment
}

// Attachment attachment struct
type Attachment struct {
	Type     string
	FileName string
	Data     *os.File
}
