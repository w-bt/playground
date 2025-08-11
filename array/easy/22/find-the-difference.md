
# Find the Difference

**Original link:** https://leetcode.com/problems/find-the-difference/

## Ringkasan
Diberikan dua string `s` dan `t`. String `t` berasal dari `s` yang diacak lalu **ditambah tepat satu** karakter ekstra. Temukan karakter ekstra tersebut.

## Contoh
- `s = "abcd"`, `t = "cbdae"` → hasil: `"e"`
- `s = ""`, `t = "y"` → hasil: `"y"`
- `s = "aa"`, `t = "aax"` → hasil: `"x"`

## Batasan
- `0 ≤ |s| ≤ 10^5`
- `|t| = |s| + 1`
- Biasanya huruf kecil `'a'..'z'` (solusi di bawah tetap berlaku untuk byte umum).

## Pendekatan
1. **XOR semua karakter** (O(n), O(1)) — pasangan karakter yang sama saling meniadakan; yang tersisa adalah karakter tambahan.
2. **Selisih jumlah kode karakter** (O(n), O(1)) — jumlahkan `ord` semua karakter `t` lalu kurangi jumlah `ord` semua karakter `s`.
3. **Hitung frekuensi** (O(n), O(1) untuk alfabet tetap) — gunakan array 26 elemen (atau map) untuk menghitung selisih frekuensi.

## Implementasi

### Go (XOR)
```go
package main

func findTheDifference(s, t string) byte {
    var x byte = 0
    for i := 0; i < len(s); i++ { x ^= s[i] }
    for i := 0; i < len(t); i++ { x ^= t[i] }
    return x
}
```

### Python (XOR)
```python
def findTheDifference(s: str, t: str) -> str:
    x = 0
    for ch in s: x ^= ord(ch)
    for ch in t: x ^= ord(ch)
    return chr(x)
```

### JavaScript (XOR)
```js
function findTheDifference(s, t) {
  let x = 0;
  for (const ch of s) x ^= ch.charCodeAt(0);
  for (const ch of t) x ^= ch.charCodeAt(0);
  return String.fromCharCode(x);
}
```

## Kompleksitas
- **Waktu:** O(n)
- **Ruang:** O(1)

## Catatan
Dokumen ini merupakan ringkasan orisinal (bukan salinan teks LeetCode).
