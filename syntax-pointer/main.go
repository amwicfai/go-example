package main

import "fmt"

type Message struct {
	Id   int
	Body string
}

//去掉*则为传值，即函数内修改不影响参数原始值
func (m *Message) Send() bool {
	fmt.Println(m.Body)
	m.Body = "sent"
	return true
}

func main() {
	m := Message{Id: 1, Body: "test"}
	fmt.Println(m.Send())
	fmt.Println(m.Body)

	//var mp2 *Message
	//mp2 = &m
	//fmt.Println("pointer: ", mp2)

	mp := &m
	fmt.Println(mp.Id)

	mpv := *mp
	fmt.Println(mpv.Id)

	//fmt.Println(*m)
}
