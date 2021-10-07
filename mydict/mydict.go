package mydict

import "errors"

var (
	ErrSearch     = errors.New("Not Find")
	errCantUpdate = errors.New("Camt update non-exitsting word")
	ErrWordExists = errors.New("that word is already exists")
)

//Dictionary
type Dictionary map[string]string

//Search
func (d Dictionary) Search(word string) (string, error) {
	//value = d[word]
	//value에 값이 존재하면 value와 nil을 반환하고
	//value에 값이 없을 경우
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", ErrSearch
}
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case ErrSearch:
		d[word] = def
	case nil:
		return ErrSearch
	}
	return nil
}
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case ErrSearch:
		return errCantUpdate
	}
	return nil
}
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

// if err == ErrSearch {
// 	d[word] = def
// } else if err == nil {
// 	return ErrWordExists
// }
// return nil
