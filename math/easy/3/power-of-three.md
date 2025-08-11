# Power of Three

**Original link:** https://leetcode.com/problems/power-of-three/

## Ringkasan
Diberikan sebuah bilangan bulat `n`, tentukan apakah `n` merupakan **pangkat dari 3** — yaitu ada bilangan bulat `k ≥ 0` sehingga `n = 3^k`.

## Contoh
- `n = 27` → `true` (3^3)
- `n = 0` → `false`
- `n = 9` → `true` (3^2)
- `n = 45` → `false`
- `n = 1` → `true` (3^0)

## Pendekatan

### 1) Bagi Tiga Berulang (Iterative / Loop)
Selama `n > 0` dan `n % 3 == 0`, bagi `n` dengan 3. Di akhir, `n` harus menjadi 1 agar `true`.
- **Waktu:** `O(log_3 n)`
- **Ruang:** `O(1)`

### 2) Cek Pembagi dari Pangkat Tertinggi (Trik 32-bit)
Untuk integer 32-bit bertanda, pangkat 3 tertinggi yang muat adalah `3^19 = 1162261467`. Jika `n` adalah pembagi dari nilai tersebut dan `n > 0`, maka `n` pasti pangkat 3.
- **Waktu:** `O(1)`
- **Ruang:** `O(1)`
> Catatan: Trik ini **khusus** untuk batas 32-bit. Untuk tipe lebih besar, ganti dengan pangkat 3 tertinggi yang muat di tipe tersebut.

### 3) Basis-3 (Opsional)
Jika merepresentasikan `n` dalam basis-3, maka pangkat 3 memiliki bentuk `1` diikuti hanya `0` (contoh: `1`, `10`, `100`, ...). Namun implementasi biasanya lebih panjang dari metode (1)/(2).

## Implementasi

### Go — Loop bagi 3
```go
package main

func isPowerOfThree(n int) bool {
    if n <= 0 {
        return false
    }
    for n%3 == 0 {
        n /= 3
    }
    return n == 1
}
```

### Go — Trik 32-bit
```go
func isPowerOfThreeFast(n int) bool {
    // 3^19 = 1162261467, masih muat di int32
    const maxPow3 = 1162261467
    return n > 0 && maxPow3%n == 0
}
```

### Python — Loop bagi 3
```python
def isPowerOfThree(n: int) -> bool:
    if n <= 0:
        return False
    while n % 3 == 0:
        n //= 3
    return n == 1
```

### JavaScript — Loop bagi 3
```js
function isPowerOfThree(n) {
  if (n <= 0) return false;
  while (n % 3 === 0) n = Math.floor(n / 3);
  return n === 1;
}
```

## Edge Cases & Tips
- `n <= 0` → `false`
- `n = 1` → `true`
- Untuk input besar dan tipe berbeda (64-bit), gunakan metode loop atau hitung pangkat 3 tertinggi yang aman untuk tipe tersebut.

## Kompleksitas
- **Loop bagi 3:** `O(log_3 n)` waktu, `O(1)` ruang.
- **Trik 32-bit:** `O(1)` waktu, `O(1)` ruang.

## Catatan
Dokumen ini adalah ringkasan orisinal (bukan salinan verbatim dari LeetCode).
