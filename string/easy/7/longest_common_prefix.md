# 14. Longest Common Prefix

**Deskripsi Masalah:**  
Tulislah sebuah fungsi yang menerima sebuah array string dan mengembalikan **prefix terpanjang** yang merupakan awalan (prefix) bersama dari semua string dalam array tersebut. Jika tidak ada prefix yang sama, fungsi harus mengembalikan string kosong `""`. ([leetcode.com](https://leetcode.com/problems/longest-common-prefix/?utm_source=chatgpt.com))

---

##  Contoh

| Input                              | Output |
|------------------------------------|--------|
| `["flower","flow","flight"]`       | `"fl"` |
| `["dog","racecar","car"]`          | `""`   |

---

##  Pendekatan Penyelesaian (Vertical Scanning)

1. Jika array kosong atau null → kembalikan `""`.
2. Iterasi setiap karakter dari string pertama (`strs[0]`):
   - Bandingkan karakter pada posisi `i` dengan karakter di posisi yang sama untuk semua string lainnya.
   - Jika indeks `i` melewati panjang salah satu string, atau karakter berbeda → hentikan dan kembalikan substring dari `0` hingga `i` dari string pertama.
3. Jika selesai tanpa mismatch, kembalikan keseluruhan `strs[0]` sebagai prefix terpanjang.

---

##  Implementasi dalam Beberapa Bahasa

### Python

```python
class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs:
            return ""
        for i in range(len(strs[0])):
            for s in strs[1:]:
                if i >= len(s) or s[i] != strs[0][i]:
                    return strs[0][:i]
        return strs[0]
```

---

### Java

```java
public String longestCommonPrefix(String[] strs) {
    if (strs == null || strs.length == 0) return "";
    for (int i = 0; i < strs[0].length(); i++) {
        for (int j = 1; j < strs.length; j++) {
            if (i >= strs[j].length() || strs[j].charAt(i) != strs[0].charAt(i)) {
                return strs[0].substring(0, i);
            }
        }
    }
    return strs[0];
}
```

---

### C++

```cpp
string longestCommonPrefix(vector<string>& strs) {
    if (strs.empty()) return "";
    for (int i = 0; i < strs[0].size(); i++) {
        for (int j = 1; j < strs.size(); j++) {
            if (i >= strs[j].size() || strs[j][i] != strs[0][i]) {
                return strs[0].substr(0, i);
            }
        }
    }
    return strs[0];
}
```

---

##  Kompleksitas

- **Time complexity:** `O(n * m)`  
  di mana `n` adalah jumlah string, dan `m` adalah panjang karakter string pertama (atau batas perbandingan) ([algo.monster](https://algo.monster/liteproblems/14?utm_source=chatgpt.com), [medium.com](https://medium.com/%40AlexanderObregon/solving-the-longest-common-prefix-problem-on-leetcode-a-java-walkthrough-dd7efa5c0b9f?utm_source=chatgpt.com)).
- **Space complexity:** `O(1)`  
  karena hanya menggunakan variabel tambahan tetap tanpa struktur data tambahan.

---

##  Kesimpulan

- Gunakan **vertical scanning**: membandingkan karakter per karakter secara vertikal antar string.
- Hentikan saat mismatch pertama terjadi untuk efisiensi.
- Cocok digunakan di berbagai bahasa pemrograman dengan logika yang sama.