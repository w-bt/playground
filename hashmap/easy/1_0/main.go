package main

func ContainsDuplicate(input []int) bool {
	intMap := map[int]int{}

	for _, v := range input {
		_, ok := intMap[v]
		if !ok {
			intMap[v] = 0
		} else {
			return true
		}
	}

	return false
}
