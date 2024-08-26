package main

import "errors"

type Dictionary map[string]string

var (
	ErrWordExists       = errors.New("word already exists")
	ErrNotFound         = errors.New("could not find the word you were looking for")
	ErrWordDoesNotExist = errors.New("cannot update word because it does not exist")
)

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, body string) error {
  _, err := d.Search(key)

  switch err {
  case ErrNotFound:
    d[key] = body
  case nil:
    return ErrWordExists
  default:
    return err
  }

	return nil
}

func (d Dictionary) Update(key, body string) error {
  _, err := d.Search(key)

  switch err {
  case ErrNotFound:
    return ErrWordDoesNotExist
  case nil:
    d[key] = body
  default:
    return err
  }
  
  return nil
}
