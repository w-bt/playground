# 🧠 Coding Challenge: Merge Two Sorted Arrays

## Level
Easy

## Category
Two Pointers / Array Manipulation

## Problem Statement
You are given two integer arrays `nums1` and `nums2`, both sorted in **non-decreasing** order. Return a new array that merges the two arrays into one sorted array.

> Note: You **must not** use any built-in sorting functions.

### Function Signature
```go
func MergeSortedArrays(nums1, nums2 []int) []int
```

## Examples

### Input
```go
nums1 := []int{1, 3, 5}
nums2 := []int{2, 4, 6}
```

### Output
```go
[]int{1, 2, 3, 4, 5, 6}
```

### Input
```go
nums1 := []int{1, 2, 3}
nums2 := []int{}
```

### Output
```go
[]int{1, 2, 3}
```

### Input
```go
nums1 := []int{}
nums2 := []int{0}
```

### Output
```go
[]int{0}
```

## Constraints
- Both `nums1` and `nums2` are sorted in non-decreasing order.
- 0 <= len(nums1), len(nums2) <= 10^4

## Input Patterns to Test
- Both arrays empty → `[]int{}, []int{}`
- One empty, one non-empty → `[]int{1,2}, []int{}`
- Duplicates across arrays → `[]int{1,2}, []int{2,3}` → `[]int{1,2,2,3}`
- Already interleaved → `[]int{1,4}, []int{2,3}` → `[]int{1,2,3,4}`

## Expected Time Complexity
- O(n + m), where n = len(nums1), m = len(nums2)
