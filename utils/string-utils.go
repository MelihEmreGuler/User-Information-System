package utils

import (
	. "../models"
)

func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	}
	return false
}

func LoadUsers() []User {
	return []User{}
}
