package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left != nil && right != nil {
		if left.Val == right.Val {
			leftOne := left.Left
			leftTwo := left.Right
			rightOne := right.Left
			rightTwo := right.Right
			return IsSameTree(leftOne, rightOne) && IsSameTree(leftTwo, rightTwo)
		}
	} else {
		return false
	}

	return false
}
