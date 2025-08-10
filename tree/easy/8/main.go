package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	_, invalid := DeepCount(root, 0)
	return !invalid
}

func DeepCount(root *TreeNode, count int) (int, bool) {
	if root == nil {
		return count, false
	}
	count++
	leftCount := count
	leftBreak := false
	if root.Left != nil {
		leftCount, leftBreak = DeepCount(root.Left, count)
	}
	if leftBreak {
		return leftCount, true
	}
	rightCount := count
	rightBreak := false
	if root.Right != nil {
		rightCount, rightBreak = DeepCount(root.Right, count)
	}
	if rightBreak {
		return rightCount, true
	}
	if leftCount > rightCount {
		if leftCount-rightCount > 1 {
			return leftCount, true
		}
		return leftCount, false
	} else if rightCount > leftCount {
		if rightCount-leftCount > 1 {
			return rightCount, true
		}
		return rightCount, false
	}

	return leftCount, false
}
