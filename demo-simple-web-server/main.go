package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type String string

func (h String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h)
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.Handle("/string", String("I'm a frayed knot."))

	err := http.ListenAndServe(":8087", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("server listening on 8087")
}
