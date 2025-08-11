package main

func IsSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	if len(s) == 0 {
		return true
	}

	counter := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[counter] {
			if counter == len(s)-1 {
				return true
			}
			counter++
		}
	}

	return false
}
