package main

func CharFrequency(s string) map[rune]int {
	frequency := make(map[rune]int)
	for _, char := range s {
		frequency[char]++
	}
	return frequency
}
