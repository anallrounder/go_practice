package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("not Found")
var errWordExists = errors.New("That word alreay exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nill {
		return errWordExists
	}
	return nil
}
