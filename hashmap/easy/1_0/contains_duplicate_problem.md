# 🧠 Coding Challenge: Contains Duplicate

## Level
Easy

## Category
HashMap / Set — Deteksi Duplikat

## Problem Statement
Given an integer array `nums`, return `true` if any value appears at least **twice** in the array, and return `false` if every element is distinct.

### Function Signature
```go
func ContainsDuplicate(nums []int) bool
```

## Example

### Input
```go
nums := []int{1, 2, 3, 1}
```

### Output
```go
true
```

### Input
```go
nums := []int{1, 2, 3, 4}
```

### Output
```go
false
```

## Constraints
- 1 <= len(nums) <= 10^5
- -10^9 <= nums[i] <= 10^9

## Input Patterns to Test
- Array with duplicates: `[1, 2, 3, 1]`
- Array with all distinct values: `[1, 2, 3, 4]`
- Single element array: `[1]`
- Large array with late duplicate
- All identical elements: `[5, 5, 5, 5]`

## Expected Time Complexity
- O(n), where n is the length of the array
