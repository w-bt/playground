package main

import (
	"fmt"
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	result := [][]string{}
	hashMap := make(map[string][]string)
	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)
		hashChr := fmt.Sprintf("%+v", chars)
		hashMap[hashChr] = append(hashMap[hashChr], str)
	}

	for _, val := range hashMap {
		result = append(result, val)
	}
	return result
}
