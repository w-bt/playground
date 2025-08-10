# 111. Minimum Depth of Binary Tree

**Difficulty:** Easy

## Problem Description

Given the `root` of a binary tree, return its **minimum depth**.

The minimum depth is defined as the number of nodes along the shortest path from the root node down to the **closest leaf node**. A *leaf* is a node with no children.

---

## Examples

**Example 1**  
```
Input: root = [3,9,20,null,null,15,7]
Output: 2
```  
The minimum path is: `3 → 9`.

**Example 2**  
```
Input: root = [2,null,3,null,4,null,5,null,6]
Output: 5
```  
This tree is like a right-linked list: depth equals the number of nodes.

---

## Solution Approaches

### 1. Recursive (DFS) Method  
- If the tree is empty (`root == nil`), return `0`.
- If one child is `nil`, ignore it and recurse on the other child (to avoid counting non-leaf paths).
- If both children exist, return `1 + min(leftDepth, rightDepth)`.

#### Go-style pseudocode:
```go
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil {
        return minDepth(root.Right) + 1
    }
    if root.Right == nil {
        return minDepth(root.Left) + 1
    }
    left := minDepth(root.Left)
    right := minDepth(root.Right)
    return 1 + min(left, right)
}
```

### 2. Iterative (BFS / Level-order) Method  
- Use a queue to traverse the tree level by level.
- As soon as you encounter the first leaf node (both children `nil`), return its depth.

---

## Complexity

| Approach       | Time Complexity | Space Complexity              |
|----------------|-----------------|--------------------------------|
| Recursive DFS  | O(n)            | O(h) for recursion stack      |
| BFS Iterative  | O(n)            | O(n) for queue in worst case  |

Balanced trees may have significantly less space usage for DFS, while BFS can find a leaf faster in practice.

---
