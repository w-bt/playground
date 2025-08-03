package main

import "testing"

func TestIsSameTree(t *testing.T) {
	type args struct {
		left  *TreeNode
		right *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Identical trees [1,2,3]",
			args: args{
				left:  &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
				right: &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
			},
			want: true,
		},
		{
			name: "Different structure [1,2] vs [1,null,2]",
			args: args{
				left:  &TreeNode{1, &TreeNode{2, nil, nil}, nil},
				right: &TreeNode{1, nil, &TreeNode{2, nil, nil}},
			},
			want: false,
		},
		{
			name: "Same structure, different values [1,2,1] vs [1,1,2]",
			args: args{
				left:  &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{1, nil, nil}},
				right: &TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}},
			},
			want: false,
		},
		{
			name: "Both empty trees",
			args: args{
				left:  nil,
				right: nil,
			},
			want: true,
		},
		{
			name: "One empty, one not",
			args: args{
				left:  &TreeNode{1, nil, nil},
				right: nil,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSameTree(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("IsSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
