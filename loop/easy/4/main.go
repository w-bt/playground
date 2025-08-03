package main

import "strings"

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func RomanToInt(s string) int {
	chars := strings.Split(s, "")
	total := 0
	index := 0
	for index < len(chars) {
		if index+1 < len(chars) {
			firstChar := chars[index]
			secondChar := chars[index+1]
			isValid := validateRule(firstChar, secondChar)
			if isValid {
				tempValue := romanMap[secondChar] - romanMap[firstChar]
				total = total + tempValue
				index = index + 2
			} else {
				total = total + romanMap[chars[index]]
				index = index + 1
			}
		} else {
			total = total + romanMap[chars[index]]
			index = index + 1
		}
	}

	return total
}

func validateRule(firstChar, secondChar string) bool {
	switch firstChar {
	case "I":
		if secondChar == "V" || secondChar == "X" {
			return true
		}
	case "X":
		if secondChar == "L" || secondChar == "C" {
			return true
		}
	case "C":
		if secondChar == "D" || secondChar == "M" {
			return true
		}
	}

	return false
}
