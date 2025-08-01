package main

func DistributeCandies(candies int, num_people int) []int {
	slice := make([]int, num_people)
	remainder := candies
	counter := 1
	pointer := 0
	for true {
		pointer = pointer % num_people
		if remainder < counter {
			slice[pointer] += remainder
			break
		}

		slice[pointer] += counter
		remainder -= counter
		counter += 1
		pointer += 1
	}

	return slice
}
