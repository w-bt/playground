# 🧠 Coding Challenge: Move Zeroes

## Level
Easy

## Category
Array / In-place Operations

## Problem Statement
Given an integer array `nums`, move all `0`'s to the **end** of it while maintaining the relative order of the **non-zero** elements.

You must do this **in-place** without making a copy of the array.

### Function Signature
```go
func MoveZeroes(nums []int)
```

## Example 1
### Input
```go
nums := []int{0, 1, 0, 3, 12}
```

### Output (in-place)
```go
[]int{1, 3, 12, 0, 0}
```

## Example 2
### Input
```go
nums := []int{0, 0, 1}
```

### Output
```go
[]int{1, 0, 0}
```

## Constraints
- You must modify the array in-place.
- The number of operations should be minimized.
- 1 <= len(nums) <= 10^4
- -2^31 <= nums[i] <= 2^31 - 1

## Input Patterns to Test
- All zeroes → `[]int{0, 0, 0}`
- No zeroes → `[]int{1, 2, 3}`
- Zeroes at the end → `[]int{1, 2, 3, 0, 0}`
- Zeroes at the start → `[]int{0, 0, 1, 2}`
- Interleaved zeroes → `[]int{0, 1, 0, 2, 0, 3}`
- Single element → `[]int{0}`, `[]int{1}`

## Expected Time Complexity
- O(n) time, O(1) space
