package main

import (
	"strings"
)

type User struct {
	name     string
	password string
}

var users = []User{
	{
		name: "marcelo", password: "marcelo2022",
	},
	{
		name: "josiane", password: "josi1977",
	},
}

func checkCredentials(name string, password string) bool {
	for _, value := range users {
		if strings.EqualFold(value.name, name) && strings.EqualFold(value.password, password) {
			return true
		}
	}
	return false
}
