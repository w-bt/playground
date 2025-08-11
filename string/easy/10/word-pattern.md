# Word Pattern

## Description
Given a pattern and a string `s`, find if `s` **follows** the same pattern.

Here **follow** means a full **bijection** between a character in `pattern` and a **non-empty** word in `s`:
- Each pattern character maps to exactly one word.
- Each word maps to exactly one pattern character.
- The mapping must be consistent across positions.

The string `s` is a sequence of words separated by single spaces.

---

## Examples

### Example 1
**Input:** `pattern = "abba"`, `s = "dog cat cat dog"`  
**Output:** `true`

### Example 2
**Input:** `pattern = "abba"`, `s = "dog cat cat fish"`  
**Output:** `false`

### Example 3
**Input:** `pattern = "aaaa"`, `s = "dog cat cat dog"`  
**Output:** `false`

### Example 4
**Input:** `pattern = "abba"`, `s = "dog dog dog dog"`  
**Output:** `false`

---

## Constraints
- `1 <= pattern.length <= 300`
- `1 <= s.length <= 3000`
- `pattern` contains only lowercase English letters.
- `s` contains lowercase English letters and spaces, with single spaces separating words.

---

## Approaches

### Dual-map (bijection check) — O(n)
1. Split `s` by spaces to get words; if `len(words) != len(pattern)`, return `false`.
2. Maintain two maps:
   - `p2w[ch]` → word mapped from pattern char `ch`
   - `w2p[word]` → pattern char mapped from word
3. Iterate over indices `i`:
   - If `p2w[ch]` exists and `p2w[ch] != words[i]`, return `false`.
   - If `w2p[word]` exists and `w2p[word] != ch`, return `false`.
   - Otherwise record the mapping both ways.
4. If loop finishes, return `true`.

**Time Complexity:** O(n) where n = number of words/characters.  
**Space Complexity:** O(k) where k = number of unique chars/words seen.

---
