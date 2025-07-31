# 🧠 Coding Challenge: Prefix and Suffix Sum

## Level
Easy

## Category
Array & String — Prefix/Suffix Sum

## Problem Statement
Given an integer array `nums`, return two new arrays:
1. `prefixSums`, where each element at index `i` is the sum of all elements from `0` to `i`.
2. `suffixSums`, where each element at index `i` is the sum of all elements from `i` to the end.

### Function Signature
```go
func PrefixAndSuffixSums(nums []int) ([]int, []int)
```

## Example

### Input
```go
nums := []int{1, 2, 3, 4}
```

### Output
```go
prefixSums := []int{1, 3, 6, 10}
suffixSums := []int{10, 9, 7, 4}
```

## Constraints
- 1 <= len(nums) <= 1000
- -10^4 <= nums[i] <= 10^4

## Input Patterns to Test
- Normal array: `[1, 2, 3, 4]`
- Single element: `[5]`
- All zeros: `[0, 0, 0]`
- Negative numbers: `[-1, -2, -3]`
- Mixed values: `[3, -2, 5, -1]`

## Expected Time Complexity
- O(n), where n is the length of the array
