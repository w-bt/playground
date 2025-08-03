
# Subtree of Another Tree - Easy

## Description

Given the roots of two binary trees `root` and `subRoot`, return `true` if there is a subtree of `root` with the same structure and node values as `subRoot`, and `false` otherwise.

A **subtree** of a binary tree `tree` is a tree that consists of a node in `tree` and all of this node's descendants. The tree `tree` could also be considered as a subtree of itself.

---

## Examples

### Example 1

**Input:**  
root = [3,4,5,1,2]  
subRoot = [4,1,2]

**Image:**  
![example1](example1.jpeg)

**Output:**  
`true`

---

### Example 2

**Input:**  
root = [3,4,5,1,2,null,null,null,null,0]  
subRoot = [4,1,2]

**Image:**  
![example2](example2.jpeg)

**Output:**  
`false`

---

## Constraints

- The number of nodes in the root tree is in the range [1, 2000].  
- The number of nodes in the subRoot tree is in the range [1, 1000].  
- -10⁴ <= root.val <= 10⁴  
- -10⁴ <= subRoot.val <= 10⁴
