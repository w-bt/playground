# 283. Move Zeroes

**Difficulty:** Easy

## Problem Description

Given an integer array `nums`, move all `0`'s to the end of it while maintaining the relative order of the non-zero elements.

You must do this **in-place** without making a copy of the array.

---

## Examples

**Example 1**
```
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
```

**Example 2**
```
Input: nums = [0]
Output: [0]
```

---

## Constraints

- `1 <= nums.length <= 10^4`
- `-2^31 <= nums[i] <= 2^31 - 1`

---

## Solution Approach

### Two Pointers

Use a pointer `lastNonZeroFoundAt` to track the index of the last non-zero element's position.

1. Iterate through the array:
   - If `nums[i] != 0`, assign `nums[lastNonZeroFoundAt] = nums[i]` and increment `lastNonZeroFoundAt`.
2. After all non-zero elements are moved forward, fill the remaining positions with zeros.

**Pseudocode (Go-style):**
```go
func moveZeroes(nums []int) {
    lastNonZeroFoundAt := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[lastNonZeroFoundAt], nums[i] = nums[i], nums[lastNonZeroFoundAt]
            lastNonZeroFoundAt++
        }
    }
}
```

---

## Complexity Analysis

- **Time Complexity:** O(n) — Iterate once through the array.
- **Space Complexity:** O(1) — In-place operation.

---
