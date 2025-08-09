# LeetCode Problem: Binary Tree Inorder Traversal

Given the `root` of a binary tree, return the inorder traversal of its nodes' values.

---

## Examples

**Example 1:**  
Input: `root = [1,null,2,3]`  
Output: `[1,3,2]`

**Example 2:**  
Input: `root = []`  
Output: `[]`

**Example 3:**  
Input: `root = [1]`  
Output: `[1]`

---

## Constraints

- The number of nodes in the tree is in the range `[0, 100]`.  
- `-100 <= Node.val <= 100`

---

## Approach 1: Recursive (DFS)

Inorder traversal visits nodes in the order: **left → root → right**.

```go
func inorderTraversal(root *TreeNode) []int {
    var res []int
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        res = append(res, node.Val)
        dfs(node.Right)
    }
    return res
}
```

---

## Approach 2: Iterative (Stack)

```go
func inorderTraversal(root *TreeNode) []int {
    var res []int
    stack := []*TreeNode{}
    curr := root
    
    for curr != nil || len(stack) > 0 {
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, curr.Val)
        curr = curr.Right
    }
    return res
}
```

---

## Approach 3: Morris Traversal (O(1) space)

```go
func inorderTraversal(root *TreeNode) []int {
    var res []int
    curr := root
    
    for curr != nil {
        if curr.Left == nil {
            res = append(res, curr.Val)
            curr = curr.Right
        } else {
            predecessor := curr.Left
            for predecessor.Right != nil && predecessor.Right != curr {
                predecessor = predecessor.Right
            }
            if predecessor.Right == nil {
                predecessor.Right = curr
                curr = curr.Left
            } else {
                predecessor.Right = nil
                res = append(res, curr.Val)
                curr = curr.Right
            }
        }
    }
    return res
}
```

---

## Complexity

| Method       | Time Complexity | Space Complexity |
|--------------|-----------------|------------------|
| Recursive    | O(n)            | O(n)             |
| Iterative    | O(n)            | O(n)             |
| Morris       | O(n)            | O(1)             |

---

## Go Unit Tests

```go
package main

import (
    "reflect"
    "testing"
)

type args struct {
    root *TreeNode
}

func Test_inorderTraversal(t *testing.T) {
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            name: "Example 1",
            args: args{root: treeFromSlice([]interface{}{1, nil, 2, 3})},
            want: []int{1, 3, 2},
        },
        {
            name: "Empty tree",
            args: args{root: treeFromSlice(nil)},
            want: []int{},
        },
        {
            name: "Single node",
            args: args{root: treeFromSlice([]interface{}{1})},
            want: []int{1},
        },
        {
            name: "Left skewed",
            args: args{root: treeFromSlice([]interface{}{3, 2, nil, 1})},
            want: []int{1, 2, 3},
        },
        {
            name: "Right skewed",
            args: args{root: treeFromSlice([]interface{}{1, nil, 2, nil, 3})},
            want: []int{1, 2, 3},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
            }
        })
    }
}
```
