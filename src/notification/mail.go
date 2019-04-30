package notification

import (
	"gopkg.in/gomail.v2"
)

// SendMail is ..
func SendMail() {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "sandytiger@example.com", "Sandy Sender")
	m.SetAddressHeader("To", "kitty@example.com", "Katy Perry")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "This is the body of the message.")

	d := gomail.NewPlainDialer("smtp.example.com", 587, "user", "123456")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
