package maps

import "errors"

var (
	ErrNotFound         = errors.New("Word not found")
	ErrWordExists       = errors.New("Word already exists")
	ErrWordDoesNotExist = errors.New("Word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return err
}

func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return definition, nil
}

// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }
