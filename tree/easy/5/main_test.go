package main

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				root: treeFromSlice([]interface{}{1, nil, 2, 3}),
			},
			want: []int{1, 3, 2},
		},
		{
			name: "Empty tree",
			args: args{
				root: treeFromSlice(nil),
			},
			want: []int{},
		},
		{
			name: "Single node",
			args: args{
				root: treeFromSlice([]interface{}{1}),
			},
			want: []int{1},
		},
		{
			name: "Left skewed tree",
			args: args{
				root: treeFromSlice([]interface{}{3, 2, nil, 1}),
			},
			want: []int{1, 2, 3},
		},
		{
			name: "Right skewed tree",
			args: args{
				root: treeFromSlice([]interface{}{1, nil, 2, nil, 3}),
			},
			want: []int{1, 2, 3},
		},
		{
			name: "Complete tree",
			args: args{
				root: treeFromSlice([]interface{}{1, 2, 3, 4, 5, 6, 7}),
			},
			want: []int{4, 2, 5, 1, 6, 3, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func treeFromSlice(vals []interface{}) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
