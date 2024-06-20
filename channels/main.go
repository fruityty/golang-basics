package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facbook.com",
		"http://stackoverflow.com",
		"http://youtube.com",
		"http://golang.org",
		// "http://fruitytyty.com",
	}

	// create string channel
	c := make(chan string)

	for _, link := range links {
		// create child go routine
		go checkLink(link, c)
	}

	// go main routine wait to recieve message from channel
	// fmt.Println(<-c)
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	for l := range c {
		go func(link string) {
			// sleep go routine
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	// sleep go routine
	// time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// send message to channel
	c <- link
}
