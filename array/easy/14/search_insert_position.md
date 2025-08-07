# 35. Search Insert Position

**Deskripsi Masalah:**  
Diberikan sebuah array terurut berisi integer berbeda (`nums`) dan sebuah nilai `target`. Tugasnya adalah mengembalikan indeks jika `target` ditemukan di dalam array. Jika tidak, kembalikan indeks di mana `target` harus disisipkan agar urutan tetap terjaga. Solusi harus menggunakan kompleksitas waktu **O(log n)**.  
Sumber: LeetCode ([leetcode.com](https://leetcode.com/problems/search-insert-position/?utm_source=chatgpt.com))

---

##  Contoh Kasus

| Input                     | Output |
|---------------------------|--------|
| `nums = [1, 3, 5, 6], target = 5` | `2` (ditemukan di indeks 2) |
| `nums = [1, 3, 5, 6], target = 2` | `1` (harus disisip di indeks 1) |
| `nums = [1, 3, 5, 6], target = 7` | `4` (disisip di akhir) |

---

##  Pendekatan Optimal: Binary Search (O(log n))

1. Inisialisasi dua penanda: `left = 0`, `right = len(nums)`
2. Lakukan iterasi selama `left < right`:
   - Hitung `mid = (left + right) / 2`
   - Jika `nums[mid] >= target`, geser `right = mid`
   - Jika `nums[mid] < target`, geser `left = mid + 1`
3. Ketika loop selesai, `left` menunjukkan indeks yang benar — baik ditemukan atau posisi yang tepat untuk disisipkan.

---

##  Contoh Ilustrasi

- Untuk `nums = [1, 3, 5, 6]`, `target = 2`:
  1. `left = 0`, `right = 4`, `mid = 2` → `nums[mid] = 5` (≥ target), `right = 2`
  2. `left = 0`, `right = 2`, `mid = 1` → `nums[mid] = 3` (≥ target), `right = 1`
  3. `left = 0`, `right = 1`, `mid = 0` → `nums[mid] = 1` (< target), `left = 1`
  4. Loop selesai → kembalikan `left = 1` sebagai posisi sisipan

---

##  Implementasi Kode (Beberapa Bahasa)

### Python
```python
class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums)
        while left < right:
            mid = (left + right) // 2
            if nums[mid] >= target:
                right = mid
            else:
                left = mid + 1
        return left
```

### Java
```java
public int searchInsert(int[] nums, int target) {
    int left = 0, right = nums.length;
    while (left < right) {
        int mid = (left + right) >>> 1;
        if (nums[mid] >= target) {
            right = mid;
        } else {
            left = mid + 1;
        }
    }
    return left;
}
```

### C++
```cpp
int searchInsert(vector<int>& nums, int target) {
    int left = 0, right = nums.size();
    while (left < right) {
        int mid = (left + right) / 2;
        if (nums[mid] >= target) {
            right = mid;
        } else {
            left = mid + 1;
        }
    }
    return left;
}
```

---

##  Kompleksitas

- **Time Complexity:** O(log n)  
- **Space Complexity:** O(1)

---

##  Kesimpulan

- Gunakan binary search untuk efisiensi maksimal.
- `left` bukan hanya menemukan nilai, tapi juga bisa digunakan sebagai titik penyisipan jika `target` tidak ada dalam array.