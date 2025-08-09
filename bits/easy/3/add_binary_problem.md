# Problem: Add Binary

Given two binary strings `a` and `b`, return their sum as a **binary string**.

---

## Examples

- **Example 1**  
  **Input:** `a = "11"`, `b = "1"`  
  **Output:** `"100"`

- **Example 2**  
  **Input:** `a = "1010"`, `b = "1011"`  
  **Output:** `"10101"`

---

## Constraints
- `1 <= len(a), len(b) <= 10^4`  
- `a` and `b` consist only of characters `'0'` and `'1'`.  
- Strings may be long; avoid integer conversion that can overflow.

---

## Approach (Two-Pointer + Carry)
1. Use two indices `i` and `j` starting from the ends of `a` and `b`.
2. Maintain `carry` (0 or 1).
3. At each step, add the current bits (if any) plus `carry` → compute new bit and next `carry`.
4. Append the resulting bit to a buffer, and move left.
5. After the loop, if `carry` is 1, append it.
6. Reverse the buffer to get the final binary string.

**Time Complexity:** `O(n + m)`  
**Space Complexity:** `O(n + m)` for the result buffer.

---

## Go Implementation

```go
package main

func addBinary(a string, b string) string {
    i, j := len(a)-1, len(b)-1
    carry := 0
    // use a byte slice as a stack for performance
    res := make([]byte, 0, max(len(a), len(b))+1)

    for i >= 0 || j >= 0 || carry > 0 {
        sum := carry
        if i >= 0 {
            if a[i] == '1' {
                sum++
            }
            i--
        }
        if j >= 0 {
            if b[j] == '1' {
                sum++
            }
            j--
        }
        // current bit
        if sum%2 == 1 {
            res = append(res, '1')
        } else {
            res = append(res, '0')
        }
        // next carry
        carry = sum / 2
    }

    // reverse res in-place
    for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
        res[l], res[r] = res[r], res[l]
    }
    return string(res)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

---

## Go Unit Tests (Table-Driven)

```go
package main

import "testing"

type args struct {
    a string
    b string
}

func Test_addBinary(t *testing.T) {
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "basic 1",
            args: args{a: "11", b: "1"},
            want: "100",
        },
        {
            name: "basic 2",
            args: args{a: "1010", b: "1011"},
            want: "10101",
        },
        {
            name: "carry through all digits",
            args: args{a: "111", b: "1"},
            want: "1000",
        },
        {
            name: "different lengths",
            args: args{a: "1", b: "11101"},
            want: "11110",
        },
        {
            name: "zeros and ones",
            args: args{a: "0", b: "0"},
            want: "0",
        },
        {
            name: "longer carry chain",
            args: args{a: "101111", b: "111"},
            want: "111110",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
                t.Errorf("addBinary() = %v, want %v", got, tt.want)
            }
        })
    }
}
```
