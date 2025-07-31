# 🧠 Coding Challenge: First Unique Character

## Level
Easy

## Category
HashMap / String — Character Frequency

## Problem Statement
Given a string `s`, return the index of the **first non-repeating character**. If there is no such character, return -1.

### Function Signature
```go
func FirstUniqueChar(s string) int
```

## Example

### Input
```go
s := "leetcode"
```

### Output
```go
0
```

### Input
```go
s := "loveleetcode"
```

### Output
```go
2
```

### Input
```go
s := "aabb"
```

### Output
```go
-1
```

## Constraints
- 1 <= len(s) <= 10^5
- `s` consists of only lowercase English letters.

## Input Patterns to Test
- Unique at beginning: `"leetcode"` → `0`
- Unique in middle: `"loveleetcode"` → `2`
- No unique: `"aabb"` → `-1`
- One character: `"z"` → `0`
- All unique: `"abcde"` → `0`
- Unique at end: `"aabbc"` → `4`

## Expected Time Complexity
- O(n), where n is the length of the string
