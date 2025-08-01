# Distribute Candies to People - Easy

We distribute some number of candies to a row of `n = num_people` people in the following way:

We then give 1 candy to the first person, 2 candies to the second person, and so on until we give `n` candies to the last person.

Then, we go back to the start of the row, giving `n + 1` candies to the first person, `n + 2` candies to the second person, and so on until we give `2 * n` candies to the last person.

This process repeats (with us giving one more candy each time, and moving to the start of the row after we reach the end) until we run out of candies. The last person will receive all of our remaining candies (not necessarily one more than the previous gift).

Return an array (of length `num_people` and sum `candies`) that represents the final distribution of candies.

---

### Example 1:
**Input:**  
`candies = 7`, `num_people = 4`  
**Output:** `[1, 2, 3, 1]`  
**Explanation:**  
- Turn 1: ans[0] += 1 → `[1, 0, 0, 0]`  
- Turn 2: ans[1] += 2 → `[1, 2, 0, 0]`  
- Turn 3: ans[2] += 3 → `[1, 2, 3, 0]`  
- Turn 4: ans[3] += 1 → `[1, 2, 3, 1]`

---

### Example 2:
**Input:**  
`candies = 10`, `num_people = 3`  
**Output:** `[5, 2, 3]`  
**Explanation:**  
- Turn 1: ans[0] += 1 → `[1, 0, 0]`  
- Turn 2: ans[1] += 2 → `[1, 2, 0]`  
- Turn 3: ans[2] += 3 → `[1, 2, 3]`  
- Turn 4: ans[0] += 4 → `[5, 2, 3]`

---

### Constraints:
- `1 <= candies <= 10^9`  
- `1 <= num_people <= 1000`
