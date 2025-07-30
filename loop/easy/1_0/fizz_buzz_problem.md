# 🎉 Coding Challenge: FizzBuzz

## Level
Easy

## Category
Logic / Loop

## Problem Statement
Given an integer `n`, return a string array `answer` (1-indexed) where:

- `answer[i] == "FizzBuzz"` if `i` is divisible by 3 and 5.
- `answer[i] == "Fizz"` if `i` is divisible by 3.
- `answer[i] == "Buzz"` if `i` is divisible by 5.
- `answer[i] == i` (as a string) if none of the above conditions are true.

### Function Signature
```go
func FizzBuzz(n int) []string
```

## Example 1
### Input
```go
n := 5
```

### Output
```go
["1", "2", "Fizz", "4", "Buzz"]
```

## Example 2
### Input
```go
n := 15
```

### Output
```go
["1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"]
```

## Constraints
- 1 <= n <= 10^4

## Input Patterns to Test
- n = 1
- n = 3 (should contain "Fizz")
- n = 5 (should contain "Buzz")
- n = 15 (should contain "FizzBuzz")

## Expected Time Complexity
- O(n)
