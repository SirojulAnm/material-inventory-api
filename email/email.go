package email

import (
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strings"
)

func SendingEmail(ReceiverEmail string, MaterialName string, SenderEmail string) error {
	// Set up MailHog configuration
	// smtpHost := "localhost"
	smtpHost := os.Getenv("MAIL")
	smtpPort := "1025"
	auth := smtp.PlainAuth("", "", "", smtpHost)

	data := struct {
		Receiver string
		Material string
		Sender   string
	}{
		Receiver: ReceiverEmail,
		Material: MaterialName,
		Sender:   SenderEmail,
	}

	bodyTemplate := `<!DOCTYPE html>
	<html>
	<head>
		<title>Notification From MIS!</title>
	</head>
	<body>
		<h1>Hello, {{.Receiver}}!</h1>
		<p>Please approve the submission of items transaction <b>{{.Material}}</b> by the user <u>{{.Sender}}</u></p>
		<p></p>
		<p>Thank You</p>
	</body>
	</html>
	`

	bodyMessage := new(strings.Builder)
	tmpl := template.Must(template.New("bodyTemplate").Parse(bodyTemplate))
	err := tmpl.Execute(bodyMessage, data)
	if err != nil {
		log.Printf("faile saat Execute body email %s", err.Error())
		return err
	}

	from := SenderEmail
	to := []string{ReceiverEmail}
	subject := "Notification From MIS!"

	auth = smtp.PlainAuth("", "", "", smtpHost)
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\nContent-Type: text/html\r\n\r\n%s", strings.Join(to, ","), subject, bodyMessage.String()))
	err = smtp.SendMail(fmt.Sprintf("%s:%s", smtpHost, smtpPort), auth, from, to, msg)
	if err != nil {
		log.Printf("faile saat send email %s", err.Error())
		return err
	}

	log.Println("Email sent successfully!")

	return nil
}
