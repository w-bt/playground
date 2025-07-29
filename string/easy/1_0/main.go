package main

import "fmt"

func main() {
	// input := "hello world" // dlrow olleh
	// input := "OpenAI" // IAnepO
	// input := "a" // a
	// input := "HelloWorld" // dlroWolleH
	// input := "x" // x
	// input := "madam" // madam
	input := "abc123!@#" // #@!321cba

	if len(input) <= 1 {
		fmt.Println(input)
		return
	}

	result := ""
	for i := len(input) - 1; i >= 0; i-- {
		result += string(input[i])
	}

	fmt.Println(result)
}

// ReverseString returns the reversed string.
func ReverseString(input string) string {
	if len(input) <= 1 {
		return input
	}
	result := ""
	for i := len(input) - 1; i >= 0; i-- {
		result += string(input[i])
	}
	return result
}
