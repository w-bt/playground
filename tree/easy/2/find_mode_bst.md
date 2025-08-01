# Find Mode in Binary Search Tree - Easy

Given the root of a **binary search tree (BST)** with **duplicates**, return all the mode(s) (i.e., the most frequently occurred element) in it.

If the tree has more than one mode, return them in any order.

---

### 📘 BST Properties:
- The left subtree of a node contains only nodes with keys **less than or equal** to the node's key.
- The right subtree of a node contains only nodes with keys **greater than or equal** to the node's key.
- Both the left and right subtrees must also be binary search trees.
- Duplicate values **are allowed**.

---

### 🧠 What is a "Mode"?
A **mode** is a value that appears **most frequently** in a dataset.

---

### 🧪 Example 1:

```
   1
    \
     2
    /
   2
```

**Input:**  
`root = [1,null,2,2]`  
**Output:** `[2]`  
**Explanation:**  
2 appears twice, 1 appears once → the mode is 2.

---

### 🧪 Example 2:

```
   0
```

**Input:**  
`root = [0]`  
**Output:** `[0]`  
**Explanation:**  
Only one node, so 0 is the mode.

---

### 📌 Constraints:
- The number of nodes in the tree is in the range `[1, 10^4]`.
- `-10^5 <= Node.val <= 10^5`

---

### 🚀 Follow-up:
Can you do it **without using any extra space**?  
(Assume that the implicit stack space incurred due to recursion **does not count**.)

**Hint:**  
You can do two in-order traversals:
1. First to find the max frequency.
2. Second to collect the values with that frequency.
