package entity

type EmailTemplate struct {
	From     string `json:"from"`
	Subject  string `json:"subject"`
	MimeType string `json:"mimeType"`
	Body     string `json:"body"`
}

type OutputEmail struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	MimeType string `json:"mimeType"`
	Body     string `json:"body"`
}
