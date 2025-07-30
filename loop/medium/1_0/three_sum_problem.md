# 🔺 Coding Challenge: 3Sum

## Level
Medium

## Category
Array / Two Pointers / Sorting

## Problem Statement
Given an integer array `nums`, return all the **triplets** `[nums[i], nums[j], nums[k]]` such that:

- `i != j`, `i != k`, and `j != k`
- `nums[i] + nums[j] + nums[k] == 0`

**Notice** that the solution set must **not contain duplicate triplets**.

### Function Signature
```go
func ThreeSum(nums []int) [][]int
```

## Example 1
### Input
```go
nums := []int{-1, 0, 1, 2, -1, -4}
```

### Output
```go
[[-1, -1, 2], [-1, 0, 1]]
```

### Explanation
Possible valid triplets:
- (-1) + 0 + 1 = 0
- (-1) + (-1) + 2 = 0  
Avoid duplicate triplets in output.

## Example 2
### Input
```go
nums := []int{0, 1, 1}
```

### Output
```go
[]
```

## Example 3
### Input
```go
nums := []int{0, 0, 0}
```

### Output
```go
[[0, 0, 0]]
```

## Constraints
- 3 <= nums.length <= 3000
- -10^5 <= nums[i] <= 10^5

## Input Patterns to Test
- Duplicates with same sum → [-1,0,1,2,-1,-4]
- No valid triplet → [0, 1, 1]
- All zero → [0, 0, 0]
- Large size + edge values → random in range [-100000, 100000]

## Expected Time Complexity
- O(n^2)
