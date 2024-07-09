package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"

	"asrlan-monolight/api/models"
	"asrlan-monolight/config"
)

type EmailData struct {
	Code string
}

// func SendEmail(to []string, subject string, cfg config.Config, htmlpath string, body models.EmailData) error {
// 	t, err := template.ParseFiles(htmlpath)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	var k bytes.Buffer
// 	err = t.Execute(&k, body)
// 	if err != nil {
// 		log.Println("failed to executing email body", err.Error())
// 		return err
// 	}

// 	if k.String() == "" {
// 		log.Println("Error buffer")
// 	}
// 	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
// 	msg := []byte(fmt.Sprintf("Subject: %s", subject) + mime + k.String())

// 	// Authentication.
// 	auth := smtp.PlainAuth("", cfg.SMTPEmail, cfg.SMTPEmailPass, cfg.SMTPHost)

// 	// Sending email.
// 	err = smtp.SendMail(cfg.SMTPHost+":"+cfg.SMTPPort, auth, cfg.SMTPEmail, to, msg)
// 	return err
// }

func SendEmailReset(to []string, subject string, cfg config.Config, htmlpath string, body models.ResetData) error {
	t, err := template.ParseFiles(htmlpath)
	if err != nil {
		log.Println(err)
		return err
	}

	var k bytes.Buffer
	err = t.Execute(&k, body)
	if err != nil {
		log.Println("failed to executing email body", err.Error())
		return err
	}

	if k.String() == "" {
		log.Println("Error buffer")
	}
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(fmt.Sprintf("Subject: %s", subject) + mime + k.String())

	// Authentication.
	auth := smtp.PlainAuth("", cfg.SMTPEmail, cfg.SMTPEmailPass, cfg.SMTPHost)

	// Sending email.
	err = smtp.SendMail(cfg.SMTPHost+":"+cfg.SMTPPort, auth, cfg.SMTPEmail, to, msg)
	return err
}

func SendEmailSignup(to []string, subject string, cfg config.Config, htmlpath string, body models.EmailData) error {
	t, err := template.ParseFiles(htmlpath)
	if err != nil {
		log.Println(err)
		return err
	}

	var k bytes.Buffer
	err = t.Execute(&k, body)
	if err != nil {
		log.Println("failed to executing email body", err.Error())
		return err
	}

	if k.String() == "" {
		log.Println("Error buffer")
	}
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(fmt.Sprintf("Subject: %s", subject) + mime + k.String())

	// Authentication.
	auth := smtp.PlainAuth("", cfg.SMTPEmail, cfg.SMTPEmailPass, cfg.SMTPHost)

	// Sending email.
	err = smtp.SendMail(cfg.SMTPHost+":"+cfg.SMTPPort, auth, cfg.SMTPEmail, to, msg)
	return err
}

type EmailCertificate struct {
	Name string
	Url  string
}

func SendEmailCertificate(to []string, subject string, htmlpath string, body EmailCertificate) error {
	t, err := template.ParseFiles(htmlpath)
	if err != nil {
		log.Println(err)
		return err
	}

	var k bytes.Buffer
	err = t.Execute(&k, body)
	if err != nil {
		log.Println("failed to executing email body", err.Error())
		return err
	}

	if k.String() == "" {
		log.Println("Error buffer")
	}
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(fmt.Sprintf("Subject: %s", subject) + mime + k.String())

	// Authentication.
	auth := smtp.PlainAuth("", "oybekatamatov999@gmail.com", "wgbtvlkeufaypcfr", "smtp.gmail.com")

	// Sending email.
	err = smtp.SendMail("smtp.gmail.com"+":"+"587", auth, "oybekatamatov999@gmail.com", to, msg)
	return err
}
