# 🔁 Coding Challenge: Valid Palindrome

## Level
Easy

## Category
String / Two Pointers

## Problem Statement
A phrase is a **palindrome** if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward.

Given a string `s`, return `true` if it is a palindrome, or `false` otherwise.

### Function Signature
```go
func IsPalindrome(s string) bool
```

## Example 1
### Input
```go
s := "A man, a plan, a canal: Panama"
```

### Output
```go
true
```

## Example 2
### Input
```go
s := "race a car"
```

### Output
```go
false
```

## Constraints
- 1 <= len(s) <= 2 * 10^5
- `s` consists only of printable ASCII characters.

## Input Patterns to Test
- Mixed case and punctuation → "No 'x' in Nixon"
- Single character → "a"
- Empty string → ""
- All symbols → "!!?.,"
- Palindrome with numbers → "12321"

## Expected Time Complexity
- O(n) time
- O(1) space
