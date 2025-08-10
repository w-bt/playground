package main

import "testing"

func TestMinDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty tree",
			args: args{
				root: nil,
			},
			want: 0,
		},
		{
			name: "Single node",
			args: args{
				root: &TreeNode{Val: 1},
			},
			want: 1,
		},
		{
			name: "Balanced tree depth 2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: 2,
		},
		{
			name: "Left child only depth 3",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
				},
			},
			want: 3,
		},
		{
			name: "Right child only depth 4",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val:   3,
							Right: &TreeNode{Val: 4},
						},
					},
				},
			},
			want: 4,
		},
		{
			name: "Minimum depth on left leaf",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:  4,
							Left: &TreeNode{Val: 5},
						},
					},
				},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDepth(tt.args.root); got != tt.want {
				t.Errorf("MinDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
