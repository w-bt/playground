package main

import (
	"testing"
)

// ===== Helpers =====

// Validate if the tree is a Binary Search Tree (BST).
func isBST(n *TreeNode, min *int, max *int) bool {
	if n == nil {
		return true
	}
	if min != nil && n.Val <= *min {
		return false
	}
	if max != nil && n.Val >= *max {
		return false
	}
	return isBST(n.Left, min, &n.Val) && isBST(n.Right, &n.Val, max)
}

// Calculate if the tree is height-balanced.
func isBalanced(n *TreeNode) bool {
	_, ok := heightOK(n)
	return ok
}

func heightOK(n *TreeNode) (int, bool) {
	if n == nil {
		return 0, true
	}
	lh, lok := heightOK(n.Left)
	rh, rok := heightOK(n.Right)
	if !lok || !rok {
		return 0, false
	}
	if lh-rh > 1 || rh-lh > 1 {
		return 0, false
	}
	if lh > rh {
		return lh + 1, true
	}
	return rh + 1, true
}

// Calculate the size of the tree.
func size(n *TreeNode) int {
	if n == nil {
		return 0
	}
	return 1 + size(n.Left) + size(n.Right)
}

// ===== Tests =====

func TestSortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Single element array",
			args: args{nums: []int{1}},
		},
		{
			name: "Two elements array",
			args: args{nums: []int{1, 3}},
		},
		{
			name: "Odd number of elements",
			args: args{nums: []int{-10, -3, 0, 5, 9}},
		},
		{
			name: "Even number of elements",
			args: args{nums: []int{-5, -3, 0, 2}},
		},
		{
			name: "Larger array",
			args: args{nums: []int{-8, -5, -3, -1, 0, 2, 4, 7, 9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortedArrayToBST(tt.args.nums)

			// 1) Valid BST
			if !isBST(got, nil, nil) {
				t.Fatalf("tree is not a valid BST for input %v", tt.args.nums)
			}

			// 2) Height-balanced
			if !isBalanced(got) {
				t.Fatalf("tree is not height-balanced for input %v", tt.args.nums)
			}

			// 3) Node count matches input size
			if sz := size(got); sz != len(tt.args.nums) {
				t.Fatalf("size = %d, want %d for input %v", sz, len(tt.args.nums), tt.args.nums)
			}
		})
	}
}
