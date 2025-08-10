package main

import "testing"

func TestHasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty tree, target 0",
			args: args{
				root:      nil,
				targetSum: 0,
			},
			want: false,
		},
		{
			name: "Single node equals target",
			args: args{
				root:      &TreeNode{Val: 5},
				targetSum: 5,
			},
			want: true,
		},
		{
			name: "Single node not equal target",
			args: args{
				root:      &TreeNode{Val: 1},
				targetSum: 2,
			},
			want: false,
		},
		{
			name: "Example true path exists (5->4->11->2 = 22)",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   11,
							Left:  &TreeNode{Val: 7},
							Right: &TreeNode{Val: 2},
						},
					},
					Right: &TreeNode{
						Val:  8,
						Left: &TreeNode{Val: 13},
						Right: &TreeNode{
							Val:   4,
							Right: &TreeNode{Val: 1},
						},
					},
				},
				targetSum: 22,
			},
			want: true,
		},
		{
			name: "No path equals target",
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
				targetSum: 5, // 1->2=3, 1->3=4
			},
			want: false,
		},
		{
			name: "Sum matches non-leaf only (should be false)",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:  4,
							Left: &TreeNode{Val: 5},
						},
					},
				},
				targetSum: 4, // 1->3 = 4
			},
			want: false,
		},
		{
			name: "Right skewed true",
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
				targetSum: 1 + 2 + 3 + 4, // 10
			},
			want: true,
		},
		{
			name: "Contains zero values",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val:  0,
						Left: &TreeNode{Val: 0},
					},
					Right: &TreeNode{Val: 1},
				},
				targetSum: 0, // path 0->0->0
			},
			want: true,
		},
		{
			name: "Negative values path",
			args: args{
				root: &TreeNode{
					Val: -2,
					Right: &TreeNode{
						Val: -3,
					},
				},
				targetSum: -5, // -2->-3
			},
			want: true,
		},
		{
			name: "Multiple leaves, only one matches",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 2},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   6,
						Right: &TreeNode{Val: 1},
					},
				},
				targetSum: 12, // 5->3->4 = 12
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("HasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
