package main

func FirstUniqueChar(s string) int {
	mapHashCount := make(map[rune]int)
	mapHashIndex := make(map[rune]int)
	for i, item := range s {
		_, ok := mapHashCount[item]
		mapHashCount[item]++
		if !ok {
			mapHashIndex[item] = i
		}
	}

	for _, item := range s {
		val := mapHashCount[item]
		if val == 1 {
			return mapHashIndex[item]
		}
	}

	return -1
}
