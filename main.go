package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"sun", "ogu"}
	for _, person := range people {
		go isSexy(person, c)
	}
	//result := <-c
	//fmt.Println(result)
	fmt.Println(<-c)
	fmt.Println(<-c) //두개 받으려면 하나더 추가
	//fmt.Println(<-c) //fatal error: all goroutines are asleep - deadlock!
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println((person))
	c <- true
}
