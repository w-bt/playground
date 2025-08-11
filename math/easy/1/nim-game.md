# Nim Game

## Description
You are playing the following Nim Game with your friend: there is a heap with `n` stones. You and your friend **alternate turns**, and on each turn you may remove **1, 2, or 3** stones. **You move first.** The player who removes the **last** stone **wins**.

Given the integer `n`, return `true` if you can win the game **assuming both players play optimally**, otherwise return `false`.

---

## Examples

- **Input:** `n = 4`  
  **Output:** `false`

- **Input:** `n = 1`  
  **Output:** `true`

- **Input:** `n = 2`  
  **Output:** `true`

- **Input:** `n = 3`  
  **Output:** `true`

---

## Constraints
- `1 <= n <= 2^31 - 1`

---

## Approach (O(1))
Key observation: positions where `n` is a **multiple of 4** are losing positions for the first player. From any non-multiple of 4, you can remove `k ∈ {1,2,3}` stones to leave a multiple of 4 for your opponent and mirror their moves thereafter.

Therefore, return `n % 4 != 0`.

---

## Proof Sketch
- Base cases: `n ∈ {1,2,3}` → first player takes all stones and wins. `n=4` → whatever the first player takes (1/2/3), the second player can take the remainder and win.
- Induction: if all `4k` are losing, then `4k+1`, `4k+2`, `4k+3` are winning by moving to `4k`. Conversely, from `4(k+1)` every move goes to a winning state (`4k+1..4k+3`).

---

## Reference Implementations

### Go
```go
package main

func canWinNim(n int) bool {
    return n%4 != 0
}
```

### JavaScript
```javascript
var canWinNim = function(n) {
  return n % 4 !== 0;
};
```

### Python
```python
def canWinNim(n: int) -> bool:
    return n % 4 != 0
```

---

## Test Ideas
- Small values: `n = 1, 2, 3, 4`
- Around boundaries: `n = 7, 8, 9`
- Large values near the limit: `n = 2**31 - 1`
