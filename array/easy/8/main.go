package main

func FindLucky(arr []int) int {
	mapArr := make(map[int]int)
	for _, num := range arr {
		mapArr[num] += 1
	}

	var found bool
	largest := 0
	for key, value := range mapArr {
		if key == value {
			found = true
			if key > largest {
				largest = key
			}
		}
	}

	if !found {
		return -1
	}

	return largest
}
