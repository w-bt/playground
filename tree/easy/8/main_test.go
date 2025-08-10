package main

import "testing"

func TestIsBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty tree",
			args: args{
				root: nil,
			},
			want: true,
		},
		{
			name: "Single node",
			args: args{
				root: &TreeNode{Val: 1},
			},
			want: true,
		},
		{
			name: "Balanced tree depth 2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 5},
					},
					Right: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: true,
		},
		{
			name: "Unbalanced tree - left heavy",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}},
					},
					Right: &TreeNode{Val: 5},
				},
			},
			want: false,
		},
		{
			name: "Balanced but not perfect",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 5},
					},
				},
			},
			want: true,
		},
		{
			name: "Unbalanced - deeper right subtree",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2},
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val:   4,
							Right: &TreeNode{Val: 5},
						},
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBalanced(tt.args.root); got != tt.want {
				t.Errorf("IsBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
