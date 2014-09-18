package main

import "fmt"
import "time"

func getUserName(userId int, cname chan string) {
	fmt.Println("call getUserName")
	time.Sleep(100 * time.Millisecond)
	cname <- "zeeman"

	//关闭可选
	close(cname)
}

func getUserGroup(userNmae string, cgroups chan []string) {
	fmt.Println("call getUserGroup, User: ", userNmae)
	time.Sleep(40 * time.Millisecond)
	cgroups <- []string{"acct", "mis", "bts"}

	//关闭可选
	close(cgroups)
}

func getUserAuth(groups []string) string {
	time.Sleep(80 * time.Millisecond)
	fmt.Println("call getUserAuth, Groups: ", groups)
	return "auth list"
}

func main() {
	cname := make(chan string)
	go getUserName(1, cname)

	name, ok := <-cname
	fmt.Println("第一次读Name,", name, ok)

	cgroups := make(chan []string)
	go getUserGroup(name, cgroups)

	//Channel的发送和接收次数必须相同, 可用第二个返回值测试
	//go getUserGroup(<-cname, cgroups)

	name2, ok := <-cname
	fmt.Println("第二次读Name,", name2, ok)

	auth := getUserAuth(<-cgroups)
	fmt.Println(auth)
}
