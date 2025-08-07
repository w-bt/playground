# 58. Length of Last Word

**Deskripsi Masalah:**  
Diberikan sebuah string `s` terdiri dari huruf dan spasi. Kembalikan panjang kata terakhir dalam string tersebut. Sebuah kata didefinisikan sebagai substring maksimal yang terdiri dari karakter non-space (spasi) saja ([leetcode.com](https://leetcode.com/problems/length-of-last-word/)).

---

##  Contoh Kasus

| Input                          | Output | Penjelasan                            |
|--------------------------------|--------|----------------------------------------|
| `"Hello World"`               | `5`    | Kata terakhir adalah `"World"`         |
| `"   fly me   to   the moon  "` | `4`    | Kata terakhir adalah `"moon"`          |
| `"luffy is still joyboy"`     | `6`    | Kata terakhir adalah `"joyboy"`        |

---

##  Pendekatan Umum (Two Pointers — O(n), O(1))

1. Mulai dari akhir string dan lewati dulu semua spasi yang mungkin ada.
2. Setelah menemukan karakter non-space pertama dari belakang, itu adalah akhir kata terakhir.
3. Terus mundur hingga menemukan spasi atau mencapai awal string.
4. Panjang kata terakhir adalah `endIndex - startIndex`.

---

##  Contoh Implementasi

### Python

```python
class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        i = len(s) - 1
        while i >= 0 and s[i] == ' ':
            i -= 1
        last = i
        while i >= 0 and s[i] != ' ':
            i -= 1
        return last - i
```

### Java

```java
public int lengthOfLastWord(String s) {
    int i = s.length() - 1;
    while (i >= 0 && s.charAt(i) == ' ') {
        i--;
    }
    int last = i;
    while (i >= 0 && s.charAt(i) != ' ') {
        i--;
    }
    return last - i;
}
```

### C++

```cpp
int lengthOfLastWord(string s) {
    int i = s.size() - 1;
    while (i >= 0 && s[i] == ' ') --i;
    int last = i;
    while (i >= 0 && s[i] != ' ') --i;
    return last - i;
}
```

---

##  Kompleksitas

- **Time Complexity:** O(n)
- **Space Complexity:** O(1)

---

##  Kesimpulan

- Gunakan pemindaian mundur (backward scan) dengan dua pointer untuk menemukan panjang kata terakhir dengan efisien.