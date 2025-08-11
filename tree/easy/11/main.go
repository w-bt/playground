package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CountNodes(root *TreeNode) int {
	return RecursiveCount(root)
}

func RecursiveCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + RecursiveCount(root.Left) + RecursiveCount(root.Right)
}
