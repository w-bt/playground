package main

import "testing"

func TestIsValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1: Valid BST [2,1,3]",
			args: args{
				root: &TreeNode{2,
					&TreeNode{1, nil, nil},
					&TreeNode{3, nil, nil},
				},
			},
			want: true,
		},
		{
			name: "Example 2: Invalid BST [5,1,4,null,null,3,6]",
			args: args{
				root: &TreeNode{5,
					&TreeNode{1, nil, nil},
					&TreeNode{4,
						&TreeNode{3, nil, nil},
						&TreeNode{6, nil, nil},
					},
				},
			},
			want: false,
		},
		{
			name: "Single node",
			args: args{
				root: &TreeNode{10, nil, nil},
			},
			want: true,
		},
		{
			name: "Left child > root (invalid)",
			args: args{
				root: &TreeNode{10,
					&TreeNode{15, nil, nil},
					nil,
				},
			},
			want: false,
		},
		{
			name: "Right child < root (invalid)",
			args: args{
				root: &TreeNode{10,
					nil,
					&TreeNode{5, nil, nil},
				},
			},
			want: false,
		},
		{
			name: "Valid BST with 3 levels",
			args: args{
				root: &TreeNode{20,
					&TreeNode{10,
						&TreeNode{5, nil, nil},
						&TreeNode{15, nil, nil},
					},
					&TreeNode{30,
						&TreeNode{25, nil, nil},
						&TreeNode{35, nil, nil},
					},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBST(tt.args.root); got != tt.want {
				t.Errorf("IsValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
