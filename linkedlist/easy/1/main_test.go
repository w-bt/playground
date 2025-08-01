package main

import "testing"

func TestHasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1 - cycle at pos 1",
			args: args{
				head: makeCyclicList([]int{3, 2, 0, -4}, 1),
			},
			want: true,
		},
		{
			name: "Example 2 - cycle at pos 0",
			args: args{
				head: makeCyclicList([]int{1, 2}, 0),
			},
			want: true,
		},
		{
			name: "Example 3 - no cycle",
			args: args{
				head: makeCyclicList([]int{1}, -1),
			},
			want: false,
		},
		{
			name: "Empty list",
			args: args{
				head: nil,
			},
			want: false,
		},
		{
			name: "Single node with cycle to self",
			args: args{
				head: makeCyclicList([]int{1}, 0),
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasCycle(tt.args.head); got != tt.want {
				t.Errorf("HasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeCyclicList(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	nodes := make([]*ListNode, len(values))
	for i, val := range values {
		nodes[i] = &ListNode{Val: val}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	if pos >= 0 {
		nodes[len(values)-1].Next = nodes[pos]
	}
	return nodes[0]
}
