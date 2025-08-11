# Count Complete Tree Nodes

## Description
Given the `root` of a **complete** binary tree, return the number of the nodes in the tree.

A complete binary tree is a binary tree in which every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible.

You must design an algorithm that runs in **less than O(n)** time.

---

## Examples

### Example 1
**Input:** root = [1,2,3,4,5,6]  
**Output:** 6

### Example 2
**Input:** root = []  
**Output:** 0

### Example 3
**Input:** root = [1]  
**Output:** 1

---

## Constraints
- The number of nodes in the tree is in the range `[0, 5 * 10^4]`.
- `0 <= Node.val <= 10^5`
- The tree is **complete**.

---

## Approaches

### 1) Height-based binary search (O(log^2 n))
Key property of a complete tree: compare the **leftmost height** and **rightmost height** from the current root.
- Let `lh` be the height going down the leftmost path and `rh` the height going down the rightmost path (counting edges or levels consistently).
- If `lh == rh`, the tree is a **perfect** binary tree with node count `2^h - 1` (where `h = lh + 1` if height counted by edges).
- Otherwise, recurse to `left` and `right` subtrees.

This yields `O(log^2 n)` time because each recursion computes heights in `O(log n)` and the recursion depth is `O(log n)`.

### 2) Index-range binary search on the last level (O(log^2 n))
- Compute tree height `h`.
- For the last level (indices in `[0, 2^{h-1}-1]`), binary-search how many nodes exist by checking whether a node at index `mid` exists via path bits.
- Each existence check takes `O(log n)`; the outer binary search also `O(log n)`.

---