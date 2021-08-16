package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("not Found")
	errWordExists = errors.New("That word alreay exists")
	errCantUpdate = errors.New("can't update non-exmmisting word")
) // looks better

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	//단어가 존재하면 삭제하도록 구현해보자.
	delete(d, word)
}
