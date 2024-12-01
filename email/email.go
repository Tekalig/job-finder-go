// email/email.go
package email

import (
	"gopkg.in/gomail.v2"
)

type EmailService struct {
	dialer *gomail.Dialer
	from   string
}

func NewEmailService(host string, port int, username, password string) *EmailService {
	return &EmailService{
		dialer: gomail.NewDialer(host, port, username, password),
		from:   username,
	}
}

func (s *EmailService) SendVerificationEmail(to, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Verify your email")
	m.SetBody("text/html", generateVerificationTemplate(token))

	return s.dialer.DialAndSend(m)
}

// Add this function to fix the error
func generateVerificationTemplate(token string) string {
	return "<p>Please verify your email using the following token: " + token + "</p>"
}
