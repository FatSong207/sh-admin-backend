package utils

import (
	"SH-admin/global"
	"gopkg.in/gomail.v2"
)

func SendMail(subject string, body string, to ...string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "tom@shadmin.com")
	m.SetHeader("To", to...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, global.Config.Mail.UserName, global.Config.Mail.Password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
