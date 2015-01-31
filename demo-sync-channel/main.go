/*
通达channel同步协程可以获取返回值
*/

package main

import (
	"fmt"
	"net/http"
)

func main() {
	c := make(chan string, 2)

	var urls = []string{
		"http://www.golang.org",
		"http://www.google.com",
	}

	for _, url := range urls {
		go func(url string, c chan string) {
			http.Get(url)

			c <- "this is return value: " + url
		}(url, c)
	}

	for i := 0; i < 2; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("over")
}
