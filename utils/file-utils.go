package utils

import (
	"errors"
	"io/ioutil"
)

func ReadFile(fileName string) (string, error) {
	if IsEmpty(fileName) {
		return "", errors.New("boş veri dosya adı olarak kabul edilemez")
	}
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
