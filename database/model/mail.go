package model

type MailModel struct {
	To          []string          `json:"to"`
	Subject     string            `json:"subject"`
	BCC         []string          `json:"bcc"`
	CC          []string          `json:"cc"`
	Body        string            `json:"body"`
	Attachments []AttachmentModel `json:"attachments"`
}

type AttachmentModel struct {
	Attachment string `json:"attachment"`
	FileName   string `json:"file_name"`
}
