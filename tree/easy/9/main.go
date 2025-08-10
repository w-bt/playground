package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MinDepth(root *TreeNode) int {
	return RecursiveMinDepth(root, 0)
}

func RecursiveMinDepth(root *TreeNode, count int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return count + 1
	}
	leftCount := RecursiveMinDepth(root.Left, count+1)
	rightCount := RecursiveMinDepth(root.Right, count+1)
	if leftCount > rightCount {
		if rightCount == 0 {
			return leftCount
		}
		return rightCount
	} else if rightCount > leftCount {
		if leftCount == 0 {
			return rightCount
		}
		return leftCount
	}
	return leftCount
}
