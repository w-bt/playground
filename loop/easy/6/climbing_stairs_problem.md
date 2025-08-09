# LeetCode Problem: Climbing Stairs

You are climbing a staircase. It takes `n` steps to reach the top.  
Each time you can either climb **1** or **2** steps. In how many distinct ways can you climb to the top?

---

## Examples

- **Example 1:**  
  Input: `n = 2`  
  Output: `2`  
  Explanation: There are two ways: `1 + 1` or `2`.

- **Example 2:**  
  Input: `n = 3`  
  Output: `3`  
  Explanation: There are three ways: `1 + 1 + 1`, `1 + 2`, `2 + 1`.

---

## Constraints
- `1 <= n <= 45`

---

## Approach (Fibonacci-like DP)

Define:
```
ways(n) = ways(n-1) + ways(n-2)
```
Base cases:
- `ways(1) = 1` (just one step)
- `ways(2) = 2` (either 1+1 or 2)

This is identical to the Fibonacci sequence, so we can solve it with **O(n)** time and **O(1)** space.

---

## Go Implementation (Bottom-Up, O(1) Space)

```go
func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    prev, curr := 1, 2
    for i := 3; i <= n; i++ {
        prev, curr = curr, prev+curr
    }
    return curr
}
```

---

## Go Unit Tests (Table-Driven)

```go
package main

import "testing"

type args struct { n int }

func Test_climbStairs(t *testing.T) {
    tests := []struct {
        name string
        args args
        want int
    }{
        {"n=1", args{n: 1}, 1},
        {"n=2", args{n: 2}, 2},
        {"n=3", args{n: 3}, 3},
        {"n=4", args{n: 4}, 5},
        {"n=5", args{n: 5}, 8},
        {"n=45", args{n: 45}, 1836311903},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := climbStairs(tt.args.n); got != tt.want {
                t.Errorf("climbStairs(%d) = %d, want %d", tt.args.n, got, tt.want)
            }
        })
    }
}
```

---
