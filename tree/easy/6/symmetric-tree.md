# 101. Symmetric Tree

Given the `root` of a binary tree, return `true` *if and only if* the tree is a mirror of itself (i.e., **symmetric** around its center); otherwise return `false`.

A binary tree is symmetric if the left subtree is a mirror reflection of the right subtree: the root values are equal, the left child of the left subtree equals the right child of the right subtree, and the right child of the left subtree equals the left child of the right subtree, recursively.

---

## Examples

### Example 1
![](sandbox:/mnt/data/example1.jpg)

**Output:** `true`  
**Explanation:** The left subtree is a mirror reflection of the right subtree.

### Example 2
![](sandbox:/mnt/data/example2.jpg)

**Output:** `false`  
**Explanation:** The structure is not symmetric.

---

## Constraints
- The number of nodes in the tree is in the range **[1, 1000]**.
- Each node value is in the range **[-100, 100]**.

---

## Follow-up
- Solve it **recursively**.  
- Then solve it **iteratively** (e.g., using a queue to compare nodes in pairs).

---

## Function Signature (Go)

```go
// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// Return true if the tree is symmetric.
func isSymmetric(root *TreeNode) bool
```

### Hints
- Two trees are mirrors if:
  1. `root1.Val == root2.Val`
  2. `root1.Left` mirrors `root2.Right`
  3. `root1.Right` mirrors `root2.Left`
- For the iterative version, push pairs `(root.Left, root.Right)` into a queue and always compare in mirrored order.
