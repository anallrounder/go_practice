package main

import (
	"fmt"
	"go_test/modules/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First world"}
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(definition)
}
