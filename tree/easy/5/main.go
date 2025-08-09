package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	result := []int{}
	recursiveTraversal(root, &result)
	return result
}

func recursiveTraversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	recursiveTraversal(root.Left, result)
	*result = append(*result, root.Val)
	recursiveTraversal(root.Right, result)
}
