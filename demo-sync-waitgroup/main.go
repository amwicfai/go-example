package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var urls = []string{
		"http://www.golang.org",
		"http://www.google.com",
	}

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			http.Get(url)

			fmt.Println(url)
		}(url)
	}

	wg.Wait()

	fmt.Println("over")
}
