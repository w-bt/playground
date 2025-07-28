# 🧮 Find the Maximum Number in an Array

**Difficulty:** Easy  
**Topic:** Array

---

## 📝 Problem Statement

Given an array of integers `nums`, return the **maximum number** in the array.

> **Note:** You **must not** use the built-in `max()` function.

---

## ✨ Function Signature

```python
def find_max(nums: List[int]) -> int:
```

---

## 🔢 Input

- `nums`: A list of integers, with `1 <= len(nums) <= 10^4`
- Each integer: `-10^5 <= nums[i] <= 10^5`

---

## 🎯 Output

- Return a single integer — the maximum number in the array.

---

## 🔍 Examples

### Example 1

```python
Input: [3, 5, 1, 7, 2]
Output: 7
```

### Example 2

```python
Input: [-10, -3, -7, -1]
Output: -1
```

### Example 3

```python
Input: [5]
Output: 5
```

---

## 📈 Time Complexity

- Expected: **O(n)** where `n` is the length of the array

---

## ✅ Test Case Patterns to Validate

| Pattern                | Example Input               | Expected Output |
|------------------------|-----------------------------|-----------------|
| Single Element         | `[42]`                      | `42`            |
| All Positive Numbers   | `[1, 2, 3, 4, 5]`           | `5`             |
| All Negative Numbers   | `[-10, -5, -20]`            | `-5`            |
| Mixed Pos/Neg          | `[0, -1, 10, 3, -5]`        | `10`            |
| Duplicates             | `[7, 7, 7, 7]`              | `7`             |
| Large Input (10,000)   | Randomized large array      | (Max value)     |
