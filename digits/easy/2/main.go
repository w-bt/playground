package main

func AddDigits(num int) int {
	return RecursiveAddDigits(num)
}

func RecursiveAddDigits(num int) int {
	if num < 10 {
		return num
	}
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return RecursiveAddDigits(sum)
}
