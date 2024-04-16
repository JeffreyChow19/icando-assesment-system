package service

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"html/template"
	"icando/internal/worker/task"
	"icando/lib"
	"icando/utils/logger"
	"os"
	"strings"
)

type EmailService struct {
	smtpEmail    string
	smtpPassword string
	smtpHost     string
	smtpPort     int
	smtpUser     string
}

func NewEmailService(config *lib.Config) EmailService {
	return EmailService{
		smtpEmail:    config.SmtpEmail,
		smtpPassword: config.SmtpPassword,
		smtpHost:     config.SmtpHost,
		smtpPort:     config.SmtpPort,
		smtpUser:     config.SmtpUser,
	}
}

func (s *EmailService) SendEmail(recipientEmail, subject, templatePath string, data interface{}) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", s.smtpEmail)
	mailer.SetHeader("To", recipientEmail)
	mailer.SetHeader("Subject", subject)

	htmlContent, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	tmpl, err := template.New(templatePath).Parse(string(htmlContent))
	if err != nil {
		return err
	}

	var bodyText string
	buffer := &strings.Builder{}
	err = tmpl.Execute(buffer, data)
	if err != nil {
		return err
	}
	bodyText = buffer.String()

	mailer.SetBody("text/html", bodyText)

	dialer := gomail.NewDialer(
		s.smtpHost,
		s.smtpPort,
		s.smtpUser,
		s.smtpPassword,
	)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err = dialer.DialAndSend(mailer)
	if err != nil {
		logger.Log.Error("Error sending email:", err)
		return err
	}

	return nil
}

func (s *EmailService) SendQuizEmail(payload task.SendQuizEmailPayload) error {
	subject := "Kuis Baru | sekolah.mu"
	templatePath := "internal/templates/quiz_email.html"
	return s.SendEmail(payload.StudentEmail, subject, templatePath, payload)
}
