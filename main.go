package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sun")
	count("ogu")
}

func count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "looks great", i)
		time.Sleep(time.Second) //time은 go package이다
	}
}
