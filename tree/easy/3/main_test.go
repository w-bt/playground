package main

import "testing"

func TestMaxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1 - Full binary tree with depth 3",
			args: args{
				root: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: 3,
		},
		{
			name: "Example 2 - Right skewed tree",
			args: args{
				root: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 2},
				},
			},
			want: 2,
		},
		{
			name: "Single node tree",
			args: args{
				root: &TreeNode{Val: 42},
			},
			want: 1,
		},
		{
			name: "Empty tree",
			args: args{
				root: nil,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDepth(tt.args.root); got != tt.want {
				t.Errorf("MaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
