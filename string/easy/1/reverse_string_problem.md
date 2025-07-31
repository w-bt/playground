# 🔤 Reverse a String

**Difficulty:** Easy  
**Topic:** String

---

## 📝 Problem Statement

Write a function that takes a string `s` and returns the string reversed.

---

## ✨ Function Signature

```python
def reverse_string(s: str) -> str:
```

---

## 🔢 Input

- `s`: A string with length `1 <= len(s) <= 10^4`
- The string consists of printable ASCII characters.

---

## 🎯 Output

- Return a new string that is the reverse of the input `s`.

---

## 🔍 Examples

### Example 1

```python
Input: "hello"
Output: "olleh"
```

### Example 2

```python
Input: "OpenAI"
Output: "IAnepO"
```

### Example 3

```python
Input: "a"
Output: "a"
```

---

## 📈 Time Complexity

- Expected: **O(n)** where `n` is the length of the string

---

## ✅ Test Case Patterns to Validate

| Pattern                | Example Input         | Expected Output |
|------------------------|-----------------------|-----------------|
| Single Character       | `"x"`                 | `"x"`           |
| Palindrome             | `"madam"`             | `"madam"`       |
| Mixed Case             | `"HelloWorld"`        | `"dlroWolleH"`  |
| With Spaces            | `"hello world"`       | `"dlrow olleh"` |
| With Symbols           | `"abc123!@#"`         | `"#@!321cba"`   |
| Long String            | `"a"*10_000`          | `"a"*10_000`    |
