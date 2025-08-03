package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	return isValidBSTWithLimit(root, nil, nil)
}

func isValidBSTWithLimit(root *TreeNode, leftLimit *int, rightLimit *int) bool {
	if root == nil {
		return true
	}

	left := root.Left
	if left != nil {
		if root.Val <= left.Val {
			return false
		}
		if leftLimit != nil && *leftLimit >= left.Val {
			return false
		}
	}

	right := root.Right
	if right != nil {
		if right.Val <= root.Val {
			return false
		}
		if rightLimit != nil && *rightLimit <= right.Val {
			return false
		}
	}

	return isValidBSTWithLimit(left, leftLimit, &root.Val) && isValidBSTWithLimit(right, &root.Val, rightLimit)
}
