# Intersection of Two Arrays

**Original link:** https://leetcode.com/problems/intersection-of-two-arrays/

## Ringkasan
Diberikan dua array bilangan bulat `nums1` dan `nums2`, kembalikan **himpunan irisan** (elemen-elemen **unik** yang muncul di **keduanya**). Urutan hasil tidak dipersyaratkan.

## Contoh
- `nums1 = [1,2,2,1]`, `nums2 = [2,2]` → hasil: `[2]`
- `nums1 = [4,9,5]`, `nums2 = [9,4,9,8,4]` → hasil: `[4,9]` (atau `[9,4]`)

## Batasan (umum)
- Ukuran array bisa berbeda; elemen dapat berulang; bisa negatif/positif.
- Hasil harus **unik** (tanpa duplikasi), elemen boleh dalam urutan apa pun.

## Pendekatan

### 1) Hash Set (paling sederhana)
- Simpan elemen **unik** dari `nums1` ke set.
- Iterasi `nums2`; jika elemen ada di set, masukkan ke **result set**.
- Konversi set hasil ke array.
- **Waktu:** `O(m + n)` • **Ruang:** `O(m + n)` dalam kasus terburuk.

### 2) Sort + Two Pointers
- Urutkan `nums1` dan `nums2`.
- Pakai dua indeks; geser yang lebih kecil, jika sama tambahkan ke hasil (hindari duplikasi dengan cek elemen terakhir yang ditambah).
- **Waktu:** `O(m log m + n log n)` • **Ruang:** `O(1)` tambahan (jika boleh sort in-place).

### 3) Domain Terbatas (opsional)
- Jika nilai berada di domain kecil (mis. 0..1000), bisa gunakan bitset/boolean array untuk menandai kehadiran dengan memori tetap.

## Implementasi

### Go — Hash Set
```go
package main

func intersection(nums1 []int, nums2 []int) []int {
    seen := make(map[int]struct{})
    for _, x := range nums1 {
        seen[x] = struct{}{}
    }
    resSet := make(map[int]struct{})
    for _, y := range nums2 {
        if _, ok := seen[y]; ok {
            resSet[y] = struct{}{}
        }
    }
    res := make([]int, 0, len(resSet))
    for v := range resSet {
        res = append(res, v)
    }
    return res
}
```

### Python — Set Intersection
```python
from typing import List

def intersection(nums1: List[int], nums2: List[int]) -> List[int]:
    return list(set(nums1) & set(nums2))
```

### JavaScript — Set + Filter
```js
function intersection(nums1, nums2) {
  const s1 = new Set(nums1);
  const res = new Set();
  for (const x of nums2) {
    if (s1.has(x)) res.add(x);
  }
  return Array.from(res);
}
```

### Go — Sort + Two Pointers (alternatif)
```go
import "sort"

func intersectionSorted(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    i, j := 0, 0
    res := []int{}
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] == nums2[j] {
            // tambahkan unik
            if len(res) == 0 || res[len(res)-1] != nums1[i] {
                res = append(res, nums1[i])
            }
            i++; j++
        } else if nums1[i] < nums2[j] {
            i++
        } else {
            j++
        }
    }
    return res
}
```

## Kompleksitas
- **Hash Set:** `O(m + n)` waktu, `O(m + n)` ruang (dalam praktik biasanya `O(min(m, n))` tambahan).
- **Sort + Two Pointers:** `O(m log m + n log n)` waktu, `O(1)` ruang ekstra jika sort in-place.

## Edge Cases & Tips
- Salah satu array kosong → hasil kosong.
- Banyak duplikasi → pastikan hasil tetap unik.
- Jika data sangat besar dan memori ketat, pertimbangkan pendekatan **sort + two pointers**.

## Catatan
Dokumen ini adalah ringkasan orisinal (bukan salinan verbatim dari LeetCode).
