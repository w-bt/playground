package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return isSubtreeChain(root, subRoot, false)
}

func isSubtreeChain(root *TreeNode, subRoot *TreeNode, prevResult bool) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root != nil && subRoot != nil {
		if root.Val != subRoot.Val {
			if prevResult {
				return false
			}
			return isSubtreeChain(root.Left, subRoot, false) || isSubtreeChain(root.Right, subRoot, false)
		} else {
			if subRoot.Left == nil && subRoot.Right == nil && (root.Left != nil || root.Right != nil) {
				return isSubtreeChain(root.Left, subRoot, false) || isSubtreeChain(root.Right, subRoot, false)
			}
			leftResult := isSubtreeChain(root.Left, subRoot.Left, true)
			rightResult := isSubtreeChain(root.Right, subRoot.Right, true)
			if !(leftResult && rightResult) {
				if prevResult {
					return false
				}
				return isSubtreeChain(root.Left, subRoot, false) || isSubtreeChain(root.Right, subRoot, false)
			}
			return true
		}
	}

	return false
}
