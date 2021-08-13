package main

import (
	"fmt"
	"go_test/modules/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First world"}
	fmt.Println(dictionary["first"])
}
