# Isomorphic Strings

## Description
Given two strings `s` and `t`, determine whether they are **isomorphic**.  
Two strings are isomorphic if characters in `s` can be replaced to get `t`, such that:
- Each character in `s` maps to exactly one character in `t` (a function).
- The mapping is **consistent** across all positions.
- No two different characters in `s` map to the **same** character in `t` (bijection on used characters).

> Note: You may not reorder characters; only a consistent one-to-one character mapping is allowed.

---

## Examples

### Example 1
**Input:** `s = "egg"`, `t = "add"`  
**Output:** `true`  
Explanation: `e -> a`, `g -> d` works consistently.

### Example 2
**Input:** `s = "foo"`, `t = "bar"`  
**Output:** `false`  
Explanation: `o` would need to map to both `a` and `r`.

### Example 3
**Input:** `s = "paper"`, `t = "title"`  
**Output:** `true`  
Explanation: `p -> t`, `a -> i`, `e -> l`, `r -> e` works.

---

## Constraints
- `1 <= s.length, t.length <= 5 * 10^4`
- `s.length == t.length`
- `s` and `t` consist of valid ASCII characters (typically lower/upper letters and symbols in the original problem).

---

## Approaches

### 1) Dual-map (bijection check) — O(n)
Maintain two maps:
- `s2t[c]` = which char in `t` is `c` mapped to
- `t2s[d]` = which char in `s` maps to `d`

While iterating positions `i`:
- If `s[i]` is already mapped, it must equal `t[i]`.
- If `t[i]` is already mapped from a different char, fail.
This ensures one-to-one mapping both ways.

### 2) Pattern signature — O(n)
Transform each string into a pattern of first-occurrence indices (e.g., `"egg" -> [0,1,1]`, `"add" -> [0,1,1]`) and compare arrays. Equal patterns imply isomorphism.

---

## Reference Implementations

### Go (Dual-map)
```go
package main

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    s2t := make(map[byte]byte, 128)
    t2s := make(map[byte]byte, 128)
    for i := 0; i < len(s); i++ {
        cs, ct := s[i], t[i]
        if v, ok := s2t[cs]; ok && v != ct {
            return false
        }
        if v, ok := t2s[ct]; ok && v != cs {
            return false
        }
        s2t[cs] = ct
        t2s[ct] = cs
    }
    return true
}
```

### JavaScript (Dual-map)
```javascript
var isIsomorphic = function(s, t) {
  if (s.length !== t.length) return false;
  const s2t = new Map();
  const t2s = new Map();
  for (let i = 0; i < s.length; i++) {
    const a = s[i], b = t[i];
    if (s2t.has(a) && s2t.get(a) !== b) return false;
    if (t2s.has(b) && t2s.get(b) !== a) return false;
    s2t.set(a, b);
    t2s.set(b, a);
  }
  return true;
};
```

### Python (Dual-map)
```python
def isIsomorphic(s: str, t: str) -> bool:
    if len(s) != len(t):
        return False
    s2t, t2s = {}, {}
    for a, b in zip(s, t):
        if a in s2t and s2t[a] != b:
            return False
        if b in t2s and t2s[b] != a:
            return False
        s2t[a] = b
        t2s[b] = a
    return True
```

---

## Complexity
- **Time Complexity:** O(n), n = length of the strings.
- **Space Complexity:** O(Σ), Σ = alphabet size used (bounded by characters seen).

---

## Test Ideas
- Simple positives: `"egg" vs "add"`, `"paper" vs "title"`
- Negatives due to conflicting mapping: `"foo" vs "bar"`, `"ab" vs "aa"`
- Mixed cases and symbols: `"ab!a" vs "cd#c"`
- Edge lengths: minimal length `1`
