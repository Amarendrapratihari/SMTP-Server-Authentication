package email

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to, subject, text, message string) error {
	from := "amarendrapratihari03@gmail.com"
	password := "fjqusevejkvdmpne"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Compose the email message
	body := "To: " + to + "\r\nSubject: " + subject + "\r\n\r\n" + text + "\n\n" + message

	// Authenticate with the SMTP server
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send the email message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(body))
	if err != nil {
		return fmt.Errorf("Error sending email: %v", err)
	}
	return nil
}
