package main

import "strings"

func IsValid(s string) bool {
	if len(s)%3 != 0 {
		return false
	}

	result := s
	for result != "" {
		tempResult := strings.Replace(result, "abc", "", -1)
		if tempResult == result {
			return false
		}
		result = tempResult
	}
	return true
}
