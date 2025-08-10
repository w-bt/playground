package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return RecursiveHasPathSum(root, targetSum, 0)
}

func RecursiveHasPathSum(root *TreeNode, targetSum int, currentSum int) bool {
	if root == nil {
		return false
	}

	currentSum += root.Val

	if root.Left == nil && root.Right == nil {
		return currentSum == targetSum
	}

	return RecursiveHasPathSum(root.Left, targetSum, currentSum) || RecursiveHasPathSum(root.Right, targetSum, currentSum)
}
