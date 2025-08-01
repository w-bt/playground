package main

import (
	"strings"
)

func NumberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	digits := splitDigit(num, 1000, 3)
	result := []string{}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 0 {
			continue
		}
		subDigits := splitDigit(digits[i], 100, 1)
		for j := len(subDigits) - 1; j >= 0; j-- {
			if subDigits[j] == 0 {
				continue
			}
			result = append(result, threeDigits(subDigits[j], j))
		}
		if i == 3 {
			result = append(result, "Billion")
		} else if i == 2 {
			result = append(result, "Million")
		} else if i == 1 {
			result = append(result, "Thousand")
		}
	}

	return strings.Join(result, " ")
}

func threeDigits(num, index int) string {
	result := ""
	if index == 1 {
		result = digitToWord(num) + " " + "Hundred"
	} else if index == 0 {
		result = digitToWord(num)
	}

	return result
}

func splitDigit(num, divNum, pow int) (nums []int) {
	div := num / divNum
	mod := num - (div * divNum)
	nums = append(nums, mod)
	newNum := num - mod
	if newNum > 0 {
		result := splitDigit(newNum/divNum, divNum, pow)
		nums = append(nums, result...)
	}
	return nums
}

func digitToWord(num int) string {
	switch num {
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	case 11:
		return "Eleven"
	case 12:
		return "Twelve"
	case 13:
		return "Thirteen"
	case 14:
		return "Fourteen"
	case 15:
		return "Fifteen"
	case 16:
		return "Sixteen"
	case 17:
		return "Seventeen"
	case 18:
		return "Eighteen"
	case 19:
		return "Nineteen"
	default:
		div := num / 10
		mod := num - (div * 10)
		str := ""
		switch div {
		case 2:
			str = "Twenty"
		case 3:
			str = "Thirty"
		case 4:
			str = "Forty"
		case 5:
			str = "Fifty"
		case 6:
			str = "Sixty"
		case 7:
			str = "Seventy"
		case 8:
			str = "Eighty"
		case 9:
			str = "Ninety"
		}

		if mod > 0 {
			str = str + " " + digitToWord(mod)
		}
		return str
	}
}
