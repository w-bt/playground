# 27. Remove Element

## 🌟 Problem Statement

Given an integer array `nums` and an integer `val`, remove all occurrences of `val` in `nums` **in‑place**. The order of the elements may be changed. Then return the number of elements in `nums` which are not equal to `val`.

Consider the number of elements in `nums` which are not equal to `val` be `k`, to get accepted, you need to do the following things:

- Change the array `nums` such that the first `k` elements of `nums` contain the elements which are not equal to `val`.
- The remaining elements of `nums` are not important as well as the size of `nums`.
- Return `k`.

Custom Judge:

```java
int[] nums = [...]; // Input array  
int val = ...;       // Value to remove  
int[] expectedNums = [...]; // The expected answer with correct length.
// It is sorted with no values equaling val.

int k = removeElement(nums, val); // Calls your implementation

assert k == expectedNums.length;
sort(nums, 0, k); // Sort the first k elements of nums
for (int i = 0; i < actualLength; i++) {
    assert nums[i] == expectedNums[i];
}
```

Jika semua assertion lolos → accepted.

---

## 📋 Examples

**Example 1:**

```
Input: nums = [3,2,2,3], val = 3  
Output: 2, nums = [2,2,_,_]  
Explanation:  
Return k = 2, dengan dua elemen pertama adalah nilai 2. Sisanya tidak penting.
```

**Example 2:**

```
Input: nums = [0,1,2,2,3,0,4,2], val = 2  
Output: 5, nums = [0,1,4,0,3,_,_,_]  
Explanation:  
Return k = 5, elemen pertama berisi nilai selain 2 (urutan bebas).
```

---

## 🔐 Constraints

- `0 <= nums.length <= 100`  
- `0 <= nums[i] <= 50`  
- `0 <= val <= 100`

---

## 🧠 Analysis & Approach

- Gunakan teknik **two pointers** (atau satu pointer `k` sebagai penampung index valid).
- Iterasi array `nums`, jika elemen `nums[i] != val`, maka taruh di `nums[k]`, lalu increment `k`.
- Setelah selesai, `k` adalah jumlah elemen valid, dan `nums[0..k-1]` berisi nilai yang dipertahankan. Sisanya bebas.

---

## 💻 Code: Implementasi Two-Pointers

### Java
```java
class Solution {
    public int removeElement(int[] nums, int val) {
        int k = 0;
        for (int j = 0; j < nums.length; j++) {
            if (nums[j] != val) {
                nums[k] = nums[j];
                k++;
            }
        }
        return k;
    }
}
```

### Python
```python
class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        k = 0
        for num in nums:
            if num != val:
                nums[k] = num
                k += 1
        return k
```

### JavaScript
```javascript
var removeElement = function(nums, val) {
    let k = 0;
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] !== val) {
            nums[k] = nums[i];
            k++;
        }
    }
    return k;
};
```

---

## ⏱️ Complexity

- **Time:** O(n) — satu iterasi penuh.
- **Space:** O(1) — hanya menggunakan variabel beberapa elemen (`k`, `i`, dll).

---

## ✅ Kesimpulan

- Problem ini menguji pemahaman tentang **modifikasi array in-place** tanpa membuat array baru.
- Implementasi optimal memakai pointer `k` untuk menomori ulang posisi elemen valid dari depan.
- Implementasi ini efisien dan memenuhi semua syarat dari custom judge LeetCode.
