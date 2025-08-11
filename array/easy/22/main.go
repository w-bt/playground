package main

func FindTheDifference(s string, t string) byte {
	mapHashS := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mapHashS[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		if _, ok := mapHashS[t[i]]; !ok || mapHashS[t[i]] == 0 {
			return t[i]
		} else {
			mapHashS[t[i]]--
		}
	}

	return byte(0)
}
