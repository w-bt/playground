package main

func IntToRoman(num int) string {
	numbers := splitDigit(num, 10)
	romanStr := ""
	for i := len(numbers) - 1; i >= 0; i-- {
		romanStr = romanStr + convertToRoman(numbers[i])
	}
	return romanStr
}

func convertToRoman(num int) (result string) {
	if num < 4 {
		start := 1
		for start <= num {
			start = start + 1
			result = result + "I"
		}
	} else if num == 4 {
		result = result + "IV"
	} else if num == 5 {
		result = result + "V"
	} else if num < 9 {
		start := 5
		result = result + "V"
		for start < num {
			result = result + "I"
			start = start + 1
		}
	} else if num == 9 {
		result = result + "IX"
	} else if num == 10 {
		result = result + "X"
	} else if num < 40 {
		start := 10
		for start <= num {
			start = start + 10
			result = result + "X"
		}
	} else if num == 40 {
		result = result + "XL"
	} else if num == 50 {
		result = result + "L"
	} else if num < 90 {
		start := 50
		result = result + "L"
		for start < num {
			result = result + "X"
			start = start + 10
		}
	} else if num == 90 {
		result = result + "XC"
	} else if num == 100 {
		result = result + "C"
	} else if num < 400 {
		start := 100
		for start <= num {
			start = start + 100
			result = result + "C"
		}
	} else if num == 400 {
		result = result + "CD"
	} else if num == 500 {
		result = result + "D"
	} else if num < 900 {
		start := 500
		result = result + "D"
		for start < num {
			result = result + "C"
			start = start + 100
		}
	} else if num == 900 {
		result = result + "CM"
	} else if num == 1000 {
		result = result + "M"
	} else {
		start := 1000
		for start <= num {
			start = start + 1000
			result = result + "M"
		}
	}
	return result
}

func splitDigit(num, divNum int) (nums []int) {
	div := num / divNum
	mod := num - (div * divNum)
	nums = append(nums, mod)
	newNum := num - mod
	if divNum*10 <= newNum {
		result := splitDigit(newNum, divNum*10)
		nums = append(nums, result...)
	} else {
		nums = append(nums, newNum)
	}
	return nums
}
