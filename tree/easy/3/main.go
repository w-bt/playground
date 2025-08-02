package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	counter := 0
	if root == nil {
		return counter
	}
	counter += 1
	left := root.Left
	counterLeft := maxNodeDepth(left, counter+1)
	right := root.Right
	counterRight := maxNodeDepth(right, counter+1)
	max := counterLeft
	if counterRight > max {
		max = counterRight
	}
	return max
}

func maxNodeDepth(t *TreeNode, counter int) int {
	if t == nil {
		return counter - 1
	}
	left := t.Left
	counterLeft := counter
	counterRight := counter
	if left != nil {
		counterLeft = maxNodeDepth(left, counter+1)
	}
	right := t.Right
	if right != nil {
		counterRight = maxNodeDepth(right, counter+1)
	}
	max := counterLeft
	if counterRight > max {
		max = counterRight
	}
	return max
}
