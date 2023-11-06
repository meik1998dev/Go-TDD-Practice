package main

import "errors"

func main() {}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	def, ok := d[key]

	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return def, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	keyIsAlreadyExisted := err == nil

	if keyIsAlreadyExisted {
		return errors.New("word already exists")
	}

	d[key] = value

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	keyIsExisted := err == nil

	if !keyIsExisted {
		return errors.New("word does not exist")
	}

	d[key] = value

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
