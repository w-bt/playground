# Reverse String

**Original link:** https://leetcode.com/problems/reverse-string/

## Ringkasan
Diberikan array karakter `s` (mis. `['h','e','l','l','o']`), balik urutannya **secara in-place** dan **kembalikan melalui modifikasi pada array yang sama** (tanpa membuat array baru berukuran `n`).

## Contoh
- `['h','e','l','l','o']` → `['o','l','l','e','h']`
- `['H','a','n','n','a','h']` → `['h','a','n','n','a','H']`

## Pendekatan: Two Pointers
Gunakan dua indeks: `l` di awal dan `r` di akhir. Selama `l < r`, tukar `s[l]` dan `s[r]`, lalu geser `l++`, `r--`.

- **Waktu:** `O(n)`
- **Ruang:** `O(1)` tambahan

> Catatan: Pada sebagian bahasa (mis. Go), string bersifat immutable; gunakan slice bertipe byte/char (atau rune untuk Unicode) yang dapat diubah di tempat.

## Implementasi

### Go (two pointers, in-place)
```go
package main

func reverseString(s []byte) {
    for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
        s[l], s[r] = s[r], s[l]
    }
}
```

### Python (two pointers, in-place)
```python
from typing import List

def reverseString(s: List[str]) -> None:
    l, r = 0, len(s) - 1
    while l < r:
        s[l], s[r] = s[r], s[l]
        l += 1
        r -= 1
```

### JavaScript (two pointers, in-place)
```js
function reverseString(s) {
  let l = 0, r = s.length - 1;
  while (l < r) {
    const tmp = s[l];
    s[l] = s[r];
    s[r] = tmp;
    l++; r--;
  }
}
```

## Edge Cases & Tips
- Panjang 0 atau 1 → tidak ada perubahan.
- Jika karakter diwakili byte ASCII, contoh di atas cukup. Untuk Unicode penuh (di luar ruang ASCII), pertimbangkan representasi yang sesuai (mis. Python `list[str]` atau Go `[]rune`).

## Disclaimer
Dokumen ini adalah ringkasan orisinal (bukan salinan teks LeetCode secara verbatim).
