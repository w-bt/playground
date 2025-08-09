# LeetCode Problem: Plus One

You are given a large integer represented as an integer array `digits`, where each `digits[i]` is the i-th digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading zeros.

Your task is to **increment the large integer by one** and return the resulting array of digits.

---

##  Examples

- **Example 1**  
  **Input:** `digits = [1, 2, 3]`  
  **Output:** `[1, 2, 4]`  
  **Explanation:** The array represents the integer 123. Incrementing by one gives 124.

- **Example 2**  
  **Input:** `digits = [9]`  
  **Output:** `[1, 0]`  
  **Explanation:** The array represents the integer 9. Incrementing by one gives 10.

---

##  Constraints

- `1 <= digits.length <= 100`
- `0 <= digits[i] <= 9`
- `digits` does not contain leading zeros.

---

##  Approach / Solution Outline

1. Traverse the array from the end (least significant digit) toward the beginning.
2. Increment the current digit by 1.
3. If the result is less than 10, **you're done**—return the updated array.
4. If the result is 10, set the current digit to 0 and **carry over** to the next digit (to the left).
5. If you finish the loop and all digits were 9 (now all zeros), insert a `1` at the front of the array.

###  Complexity
- **Time Complexity:** O(n), where n = number of digits (you may need to traverse the entire array).
- **Space Complexity:** O(1) extra, or O(n) if you create a new array when all digits are 9, requiring an extra leading digit.

---

##  Example Code (Go)

```go
func plusOne(digits []int) []int {
    for i := len(digits) - 1; i >= 0; i-- {
        digits[i]++
        digits[i] %= 10
        if digits[i] != 0 {
            return digits
        }
    }
    // All digits were 9
    return append([]int{1}, digits...)
}
