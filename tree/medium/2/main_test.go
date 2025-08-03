package main

import "testing"

func TestIsSubtree(t *testing.T) {
	type args struct {
		root    *TreeNode
		subRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1: Valid subtree",
			args: args{
				root: &TreeNode{3,
					&TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}},
					&TreeNode{5, nil, nil},
				},
				subRoot: &TreeNode{4,
					&TreeNode{1, nil, nil},
					&TreeNode{2, nil, nil},
				},
			},
			want: true,
		},
		{
			name: "Example 2: Invalid subtree due to extra node",
			args: args{
				root: &TreeNode{3,
					&TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, &TreeNode{0, nil, nil}, nil}},
					&TreeNode{5, nil, nil},
				},
				subRoot: &TreeNode{4,
					&TreeNode{1, nil, nil},
					&TreeNode{2, nil, nil},
				},
			},
			want: false,
		},
		{
			name: "Subtree is entire tree",
			args: args{
				root:    &TreeNode{1, nil, nil},
				subRoot: &TreeNode{1, nil, nil},
			},
			want: true,
		},
		{
			name: "Subtree does not exist",
			args: args{
				root: &TreeNode{1,
					&TreeNode{2, nil, nil},
					&TreeNode{3, nil, nil},
				},
				subRoot: &TreeNode{4, nil, nil},
			},
			want: false,
		},
		{
			name: "Complex subtree match",
			args: args{
				root: &TreeNode{1,
					&TreeNode{2, &TreeNode{4, nil, nil}, nil},
					&TreeNode{3, nil, nil},
				},
				subRoot: &TreeNode{2,
					&TreeNode{4, nil, nil},
					nil,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubtree(tt.args.root, tt.args.subRoot); got != tt.want {
				t.Errorf("IsSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
