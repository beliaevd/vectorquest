package email

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

const (
	fromAddress       = "dbeliaev@fixsick.ru"
	fromEmailPassword = "sE4411Jx7dd@"
	smtpServer        = "smtp.yandex.ru"
	smtpPort          = "465"
)

// SendEmail will send email to given address
func SendEmail(message string, toAddress []string) (response bool, err error) {

	var auth = smtp.PlainAuth("", fromAddress, fromEmailPassword, smtpServer)
	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, fromAddress, toAddress, []byte(message))

	if err == nil {
		return true, nil
	}
	return false, err

}

func SendYa(toEmail string, subj string, body string) {
	from := mail.Address{Name: "Name", Address: "dbeliaev@fixsick.ru"}
	to := mail.Address{Name: "", Address: toEmail}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	servername := "smtp.yandex.ru:465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", "dbeliaev@fixsick.ru", "sE4411Jx7dd@", host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	a, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = a.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = a.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = a.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := a.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	a.Quit()
}

func SendGmail() {
	// Sender data.
	from := "belyaevpilot@gmail.com"
	password := "pkfujhj,t81"

	// Receiver email address.
	to := []string{
		"EricLic@yandex.ru",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
