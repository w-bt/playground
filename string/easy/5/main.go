package main

import "regexp"

func LongestWord(sen string) string {
	pattern := regexp.MustCompile("[^a-zA-Z0-9]")
	words := pattern.Split(sen, -1)
	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}
