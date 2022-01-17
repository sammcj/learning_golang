package main

import (
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	// run checkLink in a go routine for concurrency
	for _, link := range links {
		go checkLink(link, c)
	}

	// expect one message returned on the channel for each link
	// for i := 0; i < len(links); i++ {

	// infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Alertnative infinite loop that makes it easier to see why
	// you have an infinite loop / what is blocking
	// note that you need to pass the link into the function literal so
	// that it's not incorrectly referencing the outer for loop l
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		println(link, "might be down!")
		c <- link
		return
	}

	println(link, "is up!")
	c <- link
}
