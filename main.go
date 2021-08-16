package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

func main() {
	//empty map을 initialize하고 싶을때 (안하면 nil이라 어떤 값도 넣을 수 없다.)
	//var results map[string]string 정의되지 않음
	//var results = map[string]string{} //이렇게 정의해도 되고
	var results = make(map[string]string) //make해도 된다.

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
	results["hello"] = "Hello"
	for _, url := range urls {
		//fmt.Println(url)
		result := "OK"
		err := hitURL(url)
		if err != nil { //error가 있으면
			result = "FAILED"
		}
		results[url] = result
	}

	fmt.Println(results)
}

func hitURL(url string) error {
	fmt.Println("Checking  ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
