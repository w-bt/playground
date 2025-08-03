
# Roman to Integer - Easy

## Description

Roman numerals are represented by seven different symbols:

| Symbol | Value |
|--------|-------|
| I      | 1     |
| V      | 5     |
| X      | 10    |
| L      | 50    |
| C      | 100   |
| D      | 500   |
| M      | 1000  |

For example:
- 2 is written as `II` in Roman numeral (1 + 1).
- 12 is `XII` (10 + 2).
- 27 is `XXVII` (10 + 10 + 5 + 2).

Roman numerals are usually written from largest to smallest from left to right. However, for some numbers, subtraction is used:

- `I` before `V` (5) or `X` (10) makes 4 and 9.
- `X` before `L` (50) or `C` (100) makes 40 and 90.
- `C` before `D` (500) or `M` (1000) makes 400 and 900.

Given a Roman numeral, convert it to an integer.

---

## Examples

### Example 1:
```
Input: s = "III"
Output: 3
Explanation: III = 3
```

### Example 2:
```
Input: s = "LVIII"
Output: 58
Explanation: L = 50, V = 5, III = 3
```

### Example 3:
```
Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90, IV = 4
```

---

## Constraints

- 1 <= s.length <= 15  
- s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').  
- It is guaranteed that s is a valid Roman numeral in the range [1, 3999].
