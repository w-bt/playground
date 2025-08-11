# Valid Perfect Square

**Original link:** https://leetcode.com/problems/valid-perfect-square/

## Summary
Given a **positive integer** `num`, determine whether it is a **perfect square** (i.e., `x * x == num` for some integer `x ≥ 1`). Return `true` if yes, otherwise `false`.

## Examples
- `num = 16` → `true` (4 × 4)
- `num = 14` → `false`
- `num = 1` → `true`
- `num = 808201` → `true` (899 × 899)

## Approaches

### 1) Binary Search (O(log n), O(1))
Search `x` in range `[1, num]` (or `[1, min(num, 46340)]` if using 32-bit). Compare `mid*mid` with `num` and narrow the range until found or exhausted. Use a wider type (e.g., 64-bit) to avoid overflow.

### 2) Newton's Method / Integer Sqrt (O(log n), O(1))
Iteratively refine `x ≈ sqrt(num)` via `x = (x + num // x) // 2` until `x*x ≤ num`, then check equality.

### 3) Sum of Odd Numbers (O(√n), O(1))
Squares are sums of the first `k` odd numbers: `1 + 3 + 5 + ... + (2k-1) = k²`.
Repeatedly subtract odd numbers from `num` until it becomes `0` (perfect square) or negative (not a square).

## Implementations

### Go — Binary Search (safe with int64)
```go
package main

func isPerfectSquare(num int) bool {
    if num == 1 {
        return true
    }
    n := int64(num)
    var l, r int64 = 1, n
    for l <= r {
        m := l + (r-l)/2
        sq := m * m
        if sq == n {
            return true
        } else if sq < n {
            l = m + 1
        } else {
            r = m - 1
        }
    }
    return false
}
```

### Python — Newton's Method
```python
def isPerfectSquare(num: int) -> bool:
    if num == 1:
        return True
    x = num
    while x * x > num:
        x = (x + num // x) // 2
    return x * x == num
```

### JavaScript — Binary Search
```js
function isPerfectSquare(num) {
  if (num === 1) return true;
  let l = 1, r = num;
  while (l <= r) {
    const m = Math.floor(l + (r - l) / 2);
    const sq = m * m;
    if (sq === num) return true;
    if (sq < num) l = m + 1;
    else r = m - 1;
  }
  return false;
}
```

### (Optional) Python — Sum of Odd Numbers
```python
def isPerfectSquare_odd(num: int) -> bool:
    odd = 1
    while num > 0:
        num -= odd
        odd += 2
    return num == 0
```

## Complexity
- **Binary Search / Newton:** `O(log n)` time, `O(1)` space.
- **Sum of Odds:** `O(√n)` time, `O(1)` space.

## Edge Cases & Notes
- `num = 1` should return `true`.
- Use 64-bit math if your language’s default integer may overflow on `mid*mid`.
- For extremely large inputs or many queries, Newton’s method is fast and branch-light.

## Disclaimer
This write-up is an original summary and solution set (not a verbatim copy of LeetCode's text).
