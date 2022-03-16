package main

import (
	gomail "gopkg.in/gomail.v2"
	"crypto/tls"
)

func main(){
	var (
		mail = "cocodepub@gmail.com"
		pwd = "gm879170"
	)
	m := gomail.NewMessage()
	m.SetHeader("From", mail)
	m.SetHeader("To", mail)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")


	d := gomail.NewDialer("smtp.gmail.com", 587, mail, pwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
