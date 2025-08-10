# 108. Convert Sorted Array to Binary Search Tree

**Difficulty:** Easy

## Problem Description

Given an integer array `nums` where the elements are sorted in ascending order, convert it to a **height-balanced** binary search tree.

A height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of *every* node never differs by more than one.

---

## Example

**Example 1**  
**Input:** `nums = [-10, -3, 0, 5, 9]`  
**Output:** A height-balanced BST, for example:  
```
     0
    / \
  -10  5
    \   \
    -3   9
```
(Another valid output: `[0, -3, 9, -10, null, 5]`)

**Example 2**  
**Input:** `nums = [1, 3]`  
**Output:** A height-balanced BST, e.g.:  
```
  1
   \
    3
```
or
```
  3
 /
1
```

---

## Constraints

- `1 <= nums.length <= 10^4`  
- `-10^4 <= nums[i] <= 10^4`  
- `nums` is sorted in **strictly increasing** order
