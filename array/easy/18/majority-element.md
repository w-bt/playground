# 169. Majority Element

**Difficulty:** Easy

## Problem Description

Given an array `nums` of size `n`, return *the majority element*.

The majority element is the element that appears **more than** `⌊n / 2⌋` times.  
You may assume that the majority element always exists in the array.

---

## Examples

**Example 1**
```
Input: nums = [3,2,3]
Output: 3
```

**Example 2**
```
Input: nums = [2,2,1,1,1,2,2]
Output: 2
```

---

## Constraints

- `n == nums.length`
- `1 <= n <= 5 * 10^4`
- `-10^9 <= nums[i] <= 10^9`

---

## Solution Approach

### Boyer–Moore Voting Algorithm

This algorithm works by maintaining a `count` and a `candidate` element:

1. Initialize `count = 0` and `candidate = 0` (or any value).
2. Iterate through the array:
   - If `count == 0`, set `candidate = num`.
   - If `num == candidate`, increment `count`, else decrement `count`.
3. Return `candidate`.

**Pseudocode (Go-style):**
```go
func majorityElement(nums []int) int {
    count := 0
    candidate := 0
    for _, num := range nums {
        if count == 0 {
            candidate = num
        }
        if num == candidate {
            count++
        } else {
            count--
        }
    }
    return candidate
}
```

---

## Complexity Analysis

- **Time Complexity:** O(n) — Single pass through the array.
- **Space Complexity:** O(1) — Constant extra space.

---
