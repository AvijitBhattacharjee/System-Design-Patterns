package main

import (
	"fmt"
)

type Email struct {
	mailID string
	msg    string
}

type Text struct {
	mobNumber int
	msg       string
}

type Push struct {
	mobIP string
	msg   string
}

func (e *Email) Send() {
	fmt.Println("Sending email message to = ", e.mailID, e.msg)
}

func (p *Push) Send() {
	fmt.Println("Sending push notification to = ", p.mobIP, p.msg)
}

func (t *Text) Send() {
	fmt.Println("Sending text sms to = ", t.mobNumber, t.msg)
}

func main() {

	var text = &Text{mobNumber: 9123664918, msg: "this is a text SMS"}
	var mail = &Email{mailID: "avijit@gmail.com", msg: "this is a reminder mail"}
	var push = &Push{mobIP: "localhost", msg: "this is a push notification"}

	text.Send()
	mail.Send()
	push.Send()
}
