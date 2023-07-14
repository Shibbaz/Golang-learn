package main

import(
	"errors"
)

func Search(dictionary map[string]string, index string) (string, error) {
	definition, ok := dictionary[index]
	if !ok{
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}