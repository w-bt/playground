package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return heightBalanced(nums)
}

func heightBalanced(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	mid := len(nums) / 2
	var left *TreeNode
	if mid > 0 {
		left = heightBalanced(nums[:mid])
	}
	var right *TreeNode
	if mid+1 < len(nums) {
		right = heightBalanced(nums[mid+1:])
	}
	return &TreeNode{
		Val:   nums[mid],
		Left:  left,
		Right: right,
	}
}
