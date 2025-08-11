# Is Subsequence

**Original link:** https://leetcode.com/problems/is-subsequence/

## Ringkasan
Diberikan dua string `s` dan `t`. Tentukan apakah `s` adalah **subsequence** dari `t`. Sebuah subsequence berarti kita dapat memperoleh `s` dengan **menghapus** beberapa (atau tidak ada) karakter dari `t` **tanpa mengubah urutan** karakter yang tersisa.

## Contoh
- `s = "abc"`, `t = "ahbgdc"` → **true** (ambil `a`, `b`, `c` dari urutan `t`)
- `s = "axc"`, `t = "ahbgdc"` → **false** (`x` tidak bisa dipertahankan urutannya dari `t`)
- `s = ""`, `t = "anything"` → **true** (string kosong selalu subsequence dari string apa pun)

## Batasan (umum)
- `0 ≤ |s| ≤ 100`
- `0 ≤ |t| ≤ 10^4`
- Karakter biasanya huruf kecil `'a'..'z'` (solusi di bawah juga berlaku untuk set karakter umum).

## Pendekatan Utama

### 1) Two Pointers (Linear Scan)
Gunakan dua indeks: `i` untuk `s`, `j` untuk `t`. Iterasi `t` dari kiri ke kanan; setiap kali `t[j] == s[i]`, geser `i`. Jika di akhir `i == len(s)`, maka `s` adalah subsequence dari `t`.

- **Waktu:** `O(|t|)` (karena `s` hanya maju saat cocok)
- **Ruang:** `O(1)`

### 2) Banyak Query `s` terhadap `t` yang sama (Optimisasi)
Jika kita perlu mengecek **banyak** string `s` terhadap **satu** `t` yang sama, buat **tabel loncat (next occurrence)** atau **index list per karakter**:
- **Tabel loncat (`next[i][c]`)**: posisi berikutnya kemunculan karakter `c` mulai dari indeks `i` di `t`. Cek setiap karakter `s` dengan melompat ke posisi berikutnya via tabel.  
  - Preprocessing: `O(|t| * Σ)`; Cek per `s`: `O(|s|)`.
- **Daftar indeks per karakter + binary search**: simpan semua posisi tiap karakter di `t`. Untuk setiap karakter `s`, binary search posisi berikutnya yang > posisi saat ini.  
  - Preprocessing: `O(|t|)`; Cek per `s`: `O(|s| log K)`.

## Implementasi

### Go — Two Pointers
```go
package main

func isSubsequence(s string, t string) bool {
    i := 0
    for j := 0; j < len(t) && i < len(s); j++ {
        if s[i] == t[j] {
            i++
        }
    }
    return i == len(s)
}
```

### Python — Two Pointers
```python
def isSubsequence(s: str, t: str) -> bool:
    i = 0
    for ch in t:
        if i < len(s) and s[i] == ch:
            i += 1
    return i == len(s)
```

### JavaScript — Two Pointers
```js
function isSubsequence(s, t) {
  let i = 0;
  for (const ch of t) {
    if (i < s.length && s[i] === ch) i++;
  }
  return i === s.length;
}
```

## Kompleksitas
- **Two Pointers:** Waktu `O(|t|)`, Ruang `O(1)`.
- **Banyak Query (t tetap):** Preprocessing `O(|t| * Σ)` atau `O(|t|)`, lalu per-query `O(|s|)` atau `O(|s| log K)`.

## Edge Cases & Tips
- `s` kosong → selalu `true`.
- `s` lebih panjang dari `t` → langsung `false`.
- Karakter yang tidak muncul di `t` → `false`.
- Untuk use case produksi dengan banyak query, pertimbangkan tabel loncat atau index list + binary search.

## Catatan
Dokumen ini adalah ringkasan orisinal (bukan salinan teks LeetCode).
