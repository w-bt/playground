# 136. Single Number

**Difficulty:** Easy

## Problem Description

Given a **non-empty** array of integers `nums`, every element appears **twice** except for one. Find that single one.

You must implement a solution with **O(n)** runtime complexity and **O(1)** extra space.

---

## Examples

**Example 1**
```
Input: nums = [2,2,1]
Output: 1
```

**Example 2**
```
Input: nums = [4,1,2,1,2]
Output: 4
```

**Example 3**
```
Input: nums = [1]
Output: 1
```

---

## Constraints

- `1 <= nums.length <= 3 * 10^4`
- `-3 * 10^4 <= nums[i] <= 3 * 10^4`
- Each element in the array appears twice except for one element which appears only once.

---

## Solution Approach

### Bit Manipulation (XOR)

Using the XOR operation:
- `a ^ a = 0` (number XOR itself is 0)
- `a ^ 0 = a` (number XOR 0 is the number)
- XOR is commutative and associative.

Thus, if we XOR all numbers in the array, pairs will cancel out and the single number will remain.

**Pseudocode (Go-style):**
```go
func singleNumber(nums []int) int {
    result := 0
    for _, num := range nums {
        result ^= num
    }
    return result
}
```

---

## Complexity Analysis

- **Time Complexity:** O(n) — Iterate once through the array.
- **Space Complexity:** O(1) — Constant extra space.

---
