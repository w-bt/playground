# LeetCode Problem: Sqrt(x)

Given a non-negative integer `x`, return the square root of `x` rounded down to the nearest integer. 
The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

---

## Examples

**Example 1:**  
Input: `x = 4`  
Output: `2`  
Explanation: The square root of 4 is 2, so we return 2.

**Example 2:**  
Input: `x = 8`  
Output: `2`  
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.

---

## Constraints
- `0 <= x <= 2^31 - 1`

---

## Approach

We can solve this problem using **Binary Search**:

1. If `x` is less than 2, return `x` directly.
2. Set `left = 1` and `right = x / 2`.
3. While `left <= right`:
   - Set `mid = left + (right - left) / 2`.
   - If `mid * mid` <= `x`, store `mid` in `result` and move `left` to `mid + 1`.
   - Otherwise, move `right` to `mid - 1`.
4. Return `result`.

**Time Complexity:** O(log x)  
**Space Complexity:** O(1)

---

## Go Implementation

```go
func mySqrt(x int) int {
    if x < 2 {
        return x
    }
    left, right := 1, x/2
    result := 0
    for left <= right {
        mid := left + (right - left) / 2
        if mid <= x/mid {
            result = mid
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return result
}
```

---

## Go Unit Tests

```go
package main

import "testing"

type args struct {
    x int
}

func Test_mySqrt(t *testing.T) {
    tests := []struct {
        name string
        args args
        want int
    }{
        {name: "Zero", args: args{x: 0}, want: 0},
        {name: "One", args: args{x: 1}, want: 1},
        {name: "Perfect square", args: args{x: 4}, want: 2},
        {name: "Non-perfect square", args: args{x: 8}, want: 2},
        {name: "Large square", args: args{x: 10000}, want: 100},
        {name: "Between squares", args: args{x: 15}, want: 3},
        {name: "Max int", args: args{x: 2147395600}, want: 46340},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := mySqrt(tt.args.x); got != tt.want {
                t.Errorf("mySqrt() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

---
