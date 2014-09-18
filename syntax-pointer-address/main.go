package main

import (
	"fmt"
)

type IMessage interface {
	Send() bool
}

type Message struct {
	Id   int
	Body string
}

func (m *Message) Send() bool {
	fmt.Println(m.Body)
	m.Body = "sent"
	return true
}

func main() {
	var msg_face IMessage
	m := Message{Id: 1, Body: "test"}

	//Send方法定义在：*Message
	msg_face = &m

	//如果Send方法定义在: Message
	//msg_face = m

	msg_face.Send()
}
