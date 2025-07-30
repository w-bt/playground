package main

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashMapFirst := make(map[rune]int)
	for _, v := range s {
		hashMapFirst[v]++
	}

	hashMapSecond := make(map[rune]int)
	for _, v := range t {
		hashMapSecond[v]++
	}

	if len(hashMapFirst) != len(hashMapSecond) {
		return false
	}

	isValid := true
	for k, v := range hashMapFirst {
		if v != hashMapSecond[k] {
			isValid = false
		}
	}

	return isValid
}
