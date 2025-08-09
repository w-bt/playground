package main

import "strings"

func AddBinary(a string, b string) string {
	newA, newB := standardizeBinary(a, b)
	remainder := false
	result := ""
	for i := len(newA) - 1; i >= 0; i-- {
		if newA[i] == '1' && newB[i] == '1' {
			if remainder {
				result = "1" + result
			} else {
				result = "0" + result
				remainder = true
			}
		} else if newA[i] == '0' && newB[i] == '0' {
			if remainder {
				result = "1" + result
				remainder = false
			} else {
				result = "0" + result
			}
		} else {
			if remainder {
				result = "0" + result
			} else {
				result = "1" + result
			}
		}
	}

	if remainder {
		result = "1" + result
	}

	return result
}

func standardizeBinary(a, b string) (string, string) {
	if len(a) > len(b) {
		padding := strings.Repeat("0", len(a)-len(b))
		b = padding + b
	} else if len(b) > len(a) {
		padding := strings.Repeat("0", len(b)-len(a))
		a = padding + a
	}

	return a, b
}
