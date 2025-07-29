# 🧠 Coding Challenge: Character Frequency Counter

## Level
Easy

## Category
Array & String — Frequency Count

## Problem Statement
Given a string `s`, return a dictionary (or map) where each key is a character from `s`, and the corresponding value is the number of times that character appears in `s`.

The function should be **case-sensitive**, and spaces should also be counted.

### Function Signature
```go
func CharFrequency(s string) map[rune]int
```

## Example

### Input
```go
s := "hello world"
```

### Output
```go
map[rune]int{
    'h': 1,
    'e': 1,
    'l': 3,
    'o': 2,
    ' ': 1,
    'w': 1,
    'r': 1,
    'd': 1,
}
```

## Constraints
- 1 <= len(s) <= 1000
- `s` contains printable ASCII characters

## Input Patterns to Test
- Normal string: `"hello world"`
- Empty string: `""`
- All same character: `"aaaaa"`
- Special characters: `"a! a!"`
- Mixed case: `"AaAa"`

## Expected Time Complexity
- O(n), where n is the length of the string
