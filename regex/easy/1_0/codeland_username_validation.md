## 🧩 Problem: Codeland Username Validation

**Difficulty**: Easy  
**Category**: String / Regex

### Description

Write a function `CodelandUsernameValidation` that takes a string `str` and determines if it is a valid username based on the following rules:

1. The username must be between **4 and 25 characters**.
2. It must **start with a letter** (A-Z or a-z).
3. It can only contain **letters, numbers, and the underscore** `_` character.
4. It must **not end with an underscore**.

Return `true` if the username is valid, otherwise return `false`.

---

### Function Signature

```go
func CodelandUsernameValidation(str string) bool
```

---

### Example

```go
CodelandUsernameValidation("user_1")      // true
CodelandUsernameValidation("1user")       // false
CodelandUsernameValidation("user_")       // false
CodelandUsernameValidation("us")          // false
CodelandUsernameValidation("user@name")   // false
CodelandUsernameValidation("username123") // true
```

---

### Constraints

- 1 ≤ len(str) ≤ 100
- All characters are ASCII
