package main

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				l1: buildList([]int{2, 4, 3}),
				l2: buildList([]int{5, 6, 4}),
			},
			want: buildList([]int{7, 0, 8}),
		},
		{
			name: "Example 2 - zeroes",
			args: args{
				l1: buildList([]int{0}),
				l2: buildList([]int{0}),
			},
			want: buildList([]int{0}),
		},
		{
			name: "Example 3 - carry over multiple",
			args: args{
				l1: buildList([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: buildList([]int{9, 9, 9, 9}),
			},
			want: buildList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}
