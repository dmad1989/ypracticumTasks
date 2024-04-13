package main

import "errors"

type Storage map[string]string

func (s Storage) Get(key string) (string, error) {
	val, found := s[key]
	if !found {
		return "", errors.New("not found")
	}
	return val, nil
}
