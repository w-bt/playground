package main

import "unicode"

func IsPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	pointerLeft := 0
	pointerRight := len(s) - 1
	for {
		if !unicode.IsLetter(rune(s[pointerLeft])) && !unicode.IsDigit(rune(s[pointerLeft])) {
			pointerLeft += 1
			if pointerLeft >= len(s) {
				break
			} else {
				continue
			}
		}
		if !unicode.IsLetter(rune(s[pointerRight])) && !unicode.IsDigit(rune(s[pointerRight])) {
			pointerRight -= 1
			if pointerRight < 0 {
				break
			} else {
				continue
			}
		}
		if pointerLeft > pointerRight {
			break
		}
		if unicode.ToLower(rune(s[pointerLeft])) != unicode.ToLower(rune(s[pointerRight])) {
			return false
		} else {
			pointerLeft += 1
			pointerRight -= 1
		}
	}

	return true
}
