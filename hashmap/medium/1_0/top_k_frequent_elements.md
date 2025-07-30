## 📊 Top K Frequent Elements - Medium

**Problem**  
Given an integer array `nums` and an integer `k`, return the `k` most frequent elements. You may return the answer in **any order**.

### 🧪 Examples

**Example 1:**
```
Input: nums = [1,1,1,2,2,3], k = 2  
Output: [1,2]
```

**Example 2:**
```
Input: nums = [1], k = 1  
Output: [1]
```

---

### 🧾 Constraints
- `1 <= nums.length <= 10⁵`
- `-10⁴ <= nums[i] <= 10⁴`
- `k` is in the range `[1, the number of unique elements in the array]`.
- It is **guaranteed** that the answer is **unique**.

---

### 🔍 Follow-Up
Your algorithm's time complexity must be **better than O(n log n)**, where `n` is the array's size.

---

### 🧠 Category

- Hash Map
- Heap / Priority Queue
- Sorting (with optimization)
- Bucket Sort (advanced)
