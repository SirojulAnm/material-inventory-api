package email

import (
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"net/textproto"
	"os"
)

func SendingEmail(ReceiverEmail string, MaterialName string, SenderEmail string) error {
	smtpServer := os.Getenv("MAILHOG")

	// Create the email message
	from := mail.Address{Address: SenderEmail}
	to := mail.Address{Address: ReceiverEmail}
	subject := "Notification From MIS!"
	body := fmt.Sprintf(`<!DOCTYPE html>
		<html>
		<head>
			<title>Notification From MIS!</title>
		</head>
		<body>
			<h2>Hello, %s!</h2>
			<p>Please approve the submission of items transaction materials <b>%s</b> by the user <u>%s</u></p>
			<p><br></p>
			<p>Thank You</p>
		</body>
		</html>
	`, ReceiverEmail, MaterialName, SenderEmail)

	// Compose the MIME message
	header := make(textproto.MIMEHeader)
	header.Set("From", from.String())
	header.Set("To", to.String())
	header.Set("Subject", subject)
	header.Set("Content-Type", "text/html")

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v[0])
	}
	message += "\r\n" + body

	// Dial the SMTP server
	client, err := smtp.Dial(smtpServer)
	if err != nil {
		return err
	}
	defer client.Close()

	// Send the email
	if err = client.Mail(from.Address); err != nil {
		return err
	}
	if err = client.Rcpt(to.Address); err != nil {
		return err
	}
	writer, err := client.Data()
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte(message))
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	log.Println("Email sent successfully!")

	return nil
}
