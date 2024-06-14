package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facbook.com",
		"http://stackoverflow.com",
		"http://youtube.com",
		"http://golang.org",
		"http://fruitytyty.com",
	}

	for _, link := range links {
		go checkLink(link)
	}

}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
