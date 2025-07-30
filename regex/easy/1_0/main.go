package main

import (
	"regexp"
)

func CodelandUsernameValidation(str string) string {
	match, err := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_]{2,23}[A-Za-z0-9]$", str)

	if err != nil || !match || str[len(str)-1] == '_' {
		return "false"
	}

	return "true"
}
