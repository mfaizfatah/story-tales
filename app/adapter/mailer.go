package adapter

import (
	"log"
	"os"
	"strconv"

	"github.com/mfaizfatah/story-tales/app/helpers/mailer"
	"gopkg.in/gomail.v2"
)

var (
	configEmail    = os.Getenv("SMTP_EMAIL")
	configPort     = os.Getenv("SMTP_PORT")
	configHost     = os.Getenv("SMTP_HOST")
	configPassword = os.Getenv("SMTP_PASSWORD")
)

type email struct {
	dial     *gomail.Dialer
	subject  string
	template mailer.Template
}

//Messenger ...
type Messenger interface {
	SendEmail(to ...string) error
}

//MailClient ...
type MailClient interface {
	ForgotPassword(nama, link string) Messenger
	EmailVerification(link string) Messenger
}

func (e *email) SendEmail(to ...string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", configEmail)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", e.subject)

	body := e.template.Compose()
	m.SetBody("text/html", body)

	log.Printf("from: %v\nto: %v\nsub:  %v\nasset: %v\nclient: %v\n",
		configEmail, to, e.subject, e.template, e.dial)

	return e.dial.DialAndSend(m)
}

func (e *email) ForgotPassword(nama, link string) Messenger {
	e.subject = mailer.ForgotPasswordTitle
	e.template = mailer.NewforgotPassword(nama, link)
	return e
}

func (e *email) EmailVerification(link string) Messenger {
	e.subject = mailer.VerificationTitle
	e.template = mailer.NewVerification(link)
	return e
}

//NewSMTPClient ...
func NewSMTPClient() MailClient {
	port, _ := strconv.Atoi(configPort)
	client := gomail.NewDialer(configHost, port, configEmail, configPassword)
	return &email{dial: client}
}
