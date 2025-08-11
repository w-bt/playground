package main

import "strings"

func WordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(pattern) != len(words) {
		return false
	}

	eqCharPattern := make(map[byte]string)
	eqCharS := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		if _, ok := eqCharPattern[pattern[i]]; !ok {
			if _, ok := eqCharS[words[i]]; ok {
				return false
			}
			eqCharPattern[pattern[i]] = words[i]
			eqCharS[words[i]] = pattern[i]
		} else if eqCharPattern[pattern[i]] != words[i] {
			return false
		}
	}

	return true
}
