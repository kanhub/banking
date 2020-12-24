package dict

import "errors"

var (
	errNotInDictionary = errors.New("Not in Dictionary")
	errWordExists      = errors.New("That word already exists")
	errCantUpdate      = errors.New("Can't Update")
)

// Dictionary type
type Dictionary map[string]string

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotInDictionary
}

// Add a word to dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case errNotInDictionary:
		d[word] = definition
	case nil:
		return errWordExists
	}
	return nil
}

// Update a existing key in the dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotInDictionary:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
