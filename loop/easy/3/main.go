package main

func FirstFactorial(num int) int {

	if num == 1 {
		return num
	}

	return num * FirstFactorial(num-1)

}
