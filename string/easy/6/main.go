package main

func StrStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	indexNeedle := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[indexNeedle] {
			result := checkCandidate(haystack, needle, i, indexNeedle)
			if result != -1 {
				return result
			}
		}
	}

	return -1
}

func checkCandidate(haystack string, needle string, indexHaystack, indexNeedle int) int {
	firstIndex := -1
	for i := indexHaystack; i < len(haystack); i++ {
		if indexNeedle < len(needle) {
			if haystack[i] == needle[indexNeedle] {
				if firstIndex == -1 {
					firstIndex = i
				}
				if indexNeedle == len(needle)-1 {
					return firstIndex
				}
				indexNeedle++
			} else {
				return -1
			}
		} else {
			return -1
		}
	}

	return -1
}
