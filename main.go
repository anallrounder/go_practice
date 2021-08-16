package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sun")
	go count("ogu")
	time.Sleep(time.Second * 5) //*5 goroutines가 5초동안 살아있을 수 있다.
}

func count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "looks great", i)
		time.Sleep(time.Second) //time은 go package이다
	}
}
