# Summary Ranges

## Description
You are given a **sorted** unique integer array `nums`. Return the **smallest sorted list of ranges** that cover all the numbers in the array exactly.

- Each continuous range `[a, b]` (inclusive) is represented as:
  - `"a->b"` if `a != b`
  - `"a"` if `a == b`

### Example 1
**Input**: `nums = [0,1,2,4,5,7]`  
**Output**: `["0->2","4->5","7"]`  
Explanation: Ranges are `[0,2]`, `[4,5]`, and `[7,7]`.

### Example 2
**Input**: `nums = [0,2,3,4,6,8,9]`  
**Output**: `["0","2->4","6","8->9"]`  
Explanation: Ranges are `[0,0]`, `[2,4]`, `[6,6]`, and `[8,9]`.

---

## Approach

1. Use a variable `start` to mark the beginning of a range and iterate with index `i`.
2. If the next number is consecutive (`nums[i+1] == nums[i] + 1`), continue the loop.
3. Otherwise, the range ends here:
   - If `start == nums[i]`, append `"a"`.
   - Otherwise, append `"start->nums[i]"`.
   - Set `start` to the next element.
4. Repeat until the end of the array.

---

## Example Implementation (JavaScript)

```javascript
const summaryRanges = function(nums) {
  let start = null;
  const result = [];
  for (let i = 0; i < nums.length; i++) {
    if (start === null) {
      start = nums[i];
    }
    if (nums[i] === nums[i + 1] - 1) {
      continue;
    }
    if (nums[i] === start) {
      result.push(nums[i].toString());
    } else {
      result.push(`${start}->${nums[i]}`);
    }
    start = null;
  }
  return result;
};
```

---

## Complexity
- **Time Complexity:** O(n), since each element is visited once.
- **Space Complexity:** O(n), to store the resulting list of ranges.
