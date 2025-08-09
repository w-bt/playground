# LeetCode Problem: Merge Sorted Array

You are given two integer arrays `nums1` and `nums2`, sorted in **non-decreasing order**, and two integers `m` and `n`, representing the number of elements in `nums1` and `nums2` respectively.

Merge `nums2` into `nums1` as one sorted array.

The final sorted array should not be returned by the function, but instead be stored inside the array `nums1`.  
To accommodate this, `nums1` has a length of `m + n`, where the first `m` elements denote the elements that should be merged, and the last `n` elements are set to 0 and should be ignored.

---

## Examples

**Example 1:**  
Input: `nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3`  
Output: `[1,2,2,3,5,6]`

**Example 2:**  
Input: `nums1 = [1], m = 1, nums2 = [], n = 0`  
Output: `[1]`

**Example 3:**  
Input: `nums1 = [0], m = 0, nums2 = [1], n = 1`  
Output: `[1]`

---

## Constraints

- `nums1.length == m + n`
- `nums2.length == n`
- `0 <= m, n <= 200`
- `1 <= m + n <= 200`
- `-10^9 <= nums1[i], nums2[j] <= 10^9`
- Both `nums1` and `nums2` are sorted in **non-decreasing order**.

---

## Approach

We can solve this problem **in-place** starting from the end of both arrays to avoid overwriting elements in `nums1`:

1. Initialize three pointers:
   - `p1 = m - 1` (last element in the initialized part of `nums1`)
   - `p2 = n - 1` (last element in `nums2`)
   - `p = m + n - 1` (last position in `nums1`)
2. Compare `nums1[p1]` and `nums2[p2]`, and place the larger one at `nums1[p]`.
3. Move pointers accordingly.
4. If there are remaining elements in `nums2`, copy them into `nums1`.

**Time Complexity:** O(m + n)  
**Space Complexity:** O(1)

---

## Go Implementation

```go
func merge(nums1 []int, m int, nums2 []int, n int) {
    p1 := m - 1
    p2 := n - 1
    p := m + n - 1

    for p1 >= 0 && p2 >= 0 {
        if nums1[p1] > nums2[p2] {
            nums1[p] = nums1[p1]
            p1--
        } else {
            nums1[p] = nums2[p2]
            p2--
        }
        p--
    }

    // Copy any remaining elements from nums2
    for p2 >= 0 {
        nums1[p] = nums2[p2]
        p2--
        p--
    }
}
```

---

## Go Unit Tests

```go
package main

import (
    "reflect"
    "testing"
)

type args struct {
    nums1 []int
    m     int
    nums2 []int
    n     int
}

func Test_merge(t *testing.T) {
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            name: "Example 1",
            args: args{
                nums1: []int{1, 2, 3, 0, 0, 0},
                m:     3,
                nums2: []int{2, 5, 6},
                n:     3,
            },
            want: []int{1, 2, 2, 3, 5, 6},
        },
        {
            name: "Example 2",
            args: args{
                nums1: []int{1},
                m:     1,
                nums2: []int{},
                n:     0,
            },
            want: []int{1},
        },
        {
            name: "Example 3",
            args: args{
                nums1: []int{0},
                m:     0,
                nums2: []int{1},
                n:     1,
            },
            want: []int{1},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
            if !reflect.DeepEqual(tt.args.nums1, tt.want) {
                t.Errorf("merge() = %v, want %v", tt.args.nums1, tt.want)
            }
        })
    }
}
```

---
