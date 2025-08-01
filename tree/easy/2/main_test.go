package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindMode(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1 - Mode is 2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 2,
						},
					},
				},
			},
			want: []int{2},
		},
		{
			name: "Example 2 - Single node",
			args: args{
				root: &TreeNode{
					Val: 0,
				},
			},
			want: []int{0},
		},
		{
			name: "Two modes",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  1,
						Left: &TreeNode{Val: 1},
					},
					Right: &TreeNode{
						Val:   3,
						Right: &TreeNode{Val: 3},
					},
				},
			},
			want: []int{1, 3}, // both appear twice
		},
		{
			name: "All values unique",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
			},
			want: []int{1, 2, 3}, // all appear once
		},
		{
			name: "Left skewed tree with mode 5",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			want: []int{5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMode(tt.args.root)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
