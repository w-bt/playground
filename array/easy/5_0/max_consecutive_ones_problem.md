# 🧠 Coding Challenge: Max Consecutive Ones

## Level
Easy

## Category
Array / Sliding Window

## Problem Statement
Given a binary array `nums`, return the **maximum number of consecutive 1's** in the array.

### Function Signature
```go
func FindMaxConsecutiveOnes(nums []int) int
```

## Example 1
### Input
```go
nums := []int{1, 1, 0, 1, 1, 1}
```

### Output
```go
3
```

## Example 2
### Input
```go
nums := []int{1, 0, 1, 1, 0, 1}
```

### Output
```go
2
```

## Constraints
- 1 <= len(nums) <= 10^5
- nums[i] is either 0 or 1

## Input Patterns to Test
- All 1's → `[]int{1, 1, 1, 1}`
- All 0's → `[]int{0, 0, 0}`
- Alternating → `[]int{1, 0, 1, 0, 1}`
- Single element → `[]int{1}`, `[]int{0}`
- Ends with longest → `[]int{0, 1, 1, 1}`
- Starts with longest → `[]int{1, 1, 1, 0, 0}`
- Very long input → length ~100000 (for performance)

## Expected Time Complexity
- O(n) time
- O(1) space
