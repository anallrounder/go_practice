package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [2]string{"sun", "ogu"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println("Wating for messages")
	fmt.Println("Received this message: ", <-c)
	fmt.Println("Received this message: ", <-c)
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + "is sexy"
}
