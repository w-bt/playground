# 🧠 Coding Challenge: Valid Anagram

## Level
Easy

## Category
HashMap / String — Character Count Matching

## Problem Statement
Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all original letters exactly once.

### Function Signature
```go
func IsAnagram(s string, t string) bool
```

## Example

### Input
```go
s := "anagram"
t := "nagaram"
```

### Output
```go
true
```

### Input
```go
s := "rat"
t := "car"
```

### Output
```go
false
```

## Constraints
- 1 <= len(s), len(t) <= 5 * 10^4
- `s` and `t` consist of lowercase English letters.

## Input Patterns to Test
- Valid anagram: `"listen", "silent"` → `true`
- Different lengths: `"abc", "ab"` → `false`
- Same letters, different counts: `"aabb", "ab"` → `false`
- Identical strings: `"test", "test"` → `true`
- Completely different strings: `"hello", "world"` → `false`

## Expected Time Complexity
- O(n), where n is the length of the strings
