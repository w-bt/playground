package main

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1 - Balanced Tree",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "Example 2 - Single Node",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			want: [][]int{{1}},
		},
		{
			name: "Example 3 - Empty Tree",
			args: args{
				root: nil,
			},
			want: [][]int{},
		},
		{
			name: "Unbalanced Left Tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "Unbalanced Right Tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			want: [][]int{{1}, {2}, {3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
