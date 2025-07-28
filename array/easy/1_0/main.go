package main

import "fmt"

func main() {
	//input := []int{3, 5, 1, 7, 2}
	//input := []int{-10, -3, -7, -1}
	//input := []int{5}
	//input := []int{42}
	//input := []int{1, 2, 3, 4, 5}
	//input := []int{-10, -5, -20}
	//input := []int{0, -1, 10, 3, -5}
	input := []int{7, 7, 7, 7}

	if len(input) == 0 {
		fmt.Println("Empty array")
		return
	}

	if len(input) == 1 {
		fmt.Println(input[0])
		return
	}

	max := input[0]
	for _, i := range input {
		if i > max {
			max = i
		}
	}

	fmt.Println(max)
}
