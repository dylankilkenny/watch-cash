package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
)

type Email struct {
	from    string
	to      []string
	subject string
	body    string
}

var config = Config{}

const (
	mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

// NewEmail returns a new Request obj
func NewEmail(to []string) *Email {
	return &Email{
		to:      to,
		subject: "[watch-cash] Address Activity Detected",
	}
}

func (r *Email) parseTemplate(fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}
	r.body = buffer.String()
	return nil
}

func (r *Email) sendMail() bool {
	body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + mime + "\r\n" + r.body
	SMTP := fmt.Sprintf("%s:%d", config.Server, config.Port)
	fmt.Println(config.Email)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", config.Email, config.Password, config.Server), config.Email, r.to, []byte(body)); err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(config.Email)
	return true
}

//Send processes the email
func (r *Email) Send(templateName string, items interface{}) {
	config.Load()
	err := r.parseTemplate(templateName, items)
	if err != nil {
		log.Fatal(err)
	}
	if ok := r.sendMail(); ok {
		log.Printf("Email has been sent to %s\n", r.to)
	} else {
		log.Printf("Failed to send the email to %s\n", r.to)
	}
}
