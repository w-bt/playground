# 🧠 Coding Challenge: Valid Parentheses

## Level
Easy

## Category
Stack / Queue — Basic Implementation

## Problem Statement
Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is **valid**.

An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every closing bracket has a corresponding open bracket of the same type.

### Function Signature
```go
func IsValid(s string) bool
```

## Examples

### Input
```go
s := "()"
```

### Output
```go
true
```

### Input
```go
s := "()[]{}"
```

### Output
```go
true
```

### Input
```go
s := "(]"
```

### Output
```go
false
```

### Input
```go
s := "([)]"
```

### Output
```go
false
```

### Input
```go
s := "{[]}"
```

### Output
```go
true
```

## Constraints
- 1 <= len(s) <= 10^4
- `s` consists only of `'()[]{}'`

## Input Patterns to Test
- Nested and interleaved brackets: `"([]{})"` → `true`
- Incorrect closure order: `"(]"` → `false`
- Single bracket: `"["` → `false`
- Empty string: `""` → `true`

## Expected Time Complexity
- O(n), where n is the length of the string
