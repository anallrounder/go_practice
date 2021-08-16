package main

import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

func main() {
	c := make(chan result)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func hitURL(url string, c chan<- result) {
	//fmt.Println("Checking  ", url)
	resp, err := http.Get(url)
	status := "OK"                            //default value
	if err != nil || resp.StatusCode >= 400 { //에러 있거나 400이상(에러)
		status = "FAILED"
		//c <- result{url: url, status: "FAILED"}
	}
	c <- result{url: url, status: status}
	//c <- result{url: url, status: "OK"}
}
