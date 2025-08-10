# 112. Path Sum

**Difficulty:** Easy

## Problem Description

Given the `root` of a binary tree and an integer `targetSum`, return `true` if the tree has a **root-to-leaf** path such that the sum of all node values along the path equals `targetSum`. A *leaf* is a node with no children.

---

## Examples

**Example 1**  
```
Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true
```
Explanation: Path `5 → 4 → 11 → 2` sums to 22.

**Example 2**  
```
Input: root = [1,2,3], targetSum = 5
Output: false
```
Explanation: Paths are `1→2` = 3 and `1→3` = 4; none equals 5.

**Example 3**  
```
Input: root = [], targetSum = 0
Output: false
```
Explanation: Empty tree has no root-to-leaf path.

---

## Solution Approach

A common and efficient approach is to use **DFS recursion** with accumulating sum:

1. If `root` is `nil`, return `false`.
2. Subtract `root.Val` from `targetSum`.
3. If current node is a **leaf** (`root.Left == nil && root.Right == nil`), check if `targetSum == 0`.
4. Otherwise, recursively check either subtree:
```go
return hasPathSum(root.Left, targetSum-root.Val) ||
       hasPathSum(root.Right, targetSum-root.Val)
```

---

## Complexity Analysis

| Approach            | Time Complexity | Space Complexity                     |
|---------------------|------------------|----------------------------------------|
| Recursive DFS       | O(n)             | O(h) — recursion stack (h = tree height) |

Where `n` is number of nodes, and `h` is the height of the tree.

---
