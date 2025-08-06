# 26. Remove Duplicates from Sorted Array

## 🌟 Problem Statement

Given an integer array `nums` sorted in non-decreasing order, remove the duplicates **in-place** such that each unique element appears only once. The relative order of the elements should be preserved. Then return the number of unique elements in `nums`.

After getting `k` (jumlah elemen unik), kamu wajib:
- Memodifikasi array `nums` sedemikian rupa sehingga **`nums[0:k]`** berisi elemen unik sesuai urutan awal.
- Sisanya setelah `k` **tidak penting nilainya**.
- Mengembalikan nilai `k`.

Contoh:
```
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
```
Dengan `k = 5`, elemen unik terletak di bagian awal array.

---

## 📋 Examples

**Example 1**  
Input: `nums = [1,1,2]`  
Output: `2`, `nums = [1,2,_]`

**Example 2**  
Input: `nums = [0,0,1,1,1,2,2,3,3,4]`  
Output: `5`, `nums = [0,1,2,3,4,_,_,_,_,_]`

---

## 🔐 Constraints

- `0 ≤ nums.length ≤ 3 × 10⁴`
- `-10⁴ ≤ nums[i] ≤ 10⁴`
- `nums` diurutkan non-decreasing
- Hanya **O(1)** ruang ekstra dan **in-place** modifikasi diizinkan.

---

## 🧠 Analysis & Approach

Karena `nums` sudah diurutkan, duplikat akan berdekatan. Gunakan **two-pointer** atau teknik `k` sebagai slow-pointer:

1. Inisialisasi `k = 0`.
2. Iterasi setiap elemen `x` dalam `nums`.
3. Jika `k == 0` atau `x != nums[k-1]`, lakukan:
   - `nums[k] = x`
   - `k++`
4. Setelah selesai, `k` adalah jumlah elemen unik, dan `nums[0:k]` memuat elemen tersebut.

---

## 💻 Code: Two-Pointer Approach

### Python
```python
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        k = 0
        for x in nums:
            if k == 0 or x != nums[k - 1]:
                nums[k] = x
                k += 1
        return k
```

### Java
```java
class Solution {
    public int removeDuplicates(int[] nums) {
        int k = 0;
        for (int x : nums) {
            if (k == 0 || x != nums[k - 1]) {
                nums[k] = x;
                k++;
            }
        }
        return k;
    }
}
```

### JavaScript
```javascript
var removeDuplicates = function(nums) {
    let k = 0;
    for (const x of nums) {
        if (k === 0 || x !== nums[k - 1]) {
            nums[k] = x;
            k++;
        }
    }
    return k;
};
```

---

## ⏱️ Complexity

- **Time:** O(n) — hanya satu kali iterasi.
- **Space:** O(1) — modifikasi inplace saja, tanpa struktur tambahan.

---

## ✅ Kesimpulan

- Gunakan metode **in-place** tanpa array baru.
- Teknik two-pointer (`k` slow-pointer) sangat efisien untuk memindahkan elemen unik ke depan.
- Output `k` merupakan jumlah elemen unik, dan `nums[0:k]` memuat mereka dalam urutan asli.
