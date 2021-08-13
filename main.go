package main

import (
	"fmt"
	"go_test/modules/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First world"}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update("dddd", "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}
