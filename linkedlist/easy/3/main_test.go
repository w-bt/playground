package main

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				head: listFromSlice([]int{1, 1, 2}),
			},
			want: listFromSlice([]int{1, 2}),
		},
		{
			name: "Example 2",
			args: args{
				head: listFromSlice([]int{1, 1, 2, 3, 3}),
			},
			want: listFromSlice([]int{1, 2, 3}),
		},
		{
			name: "Empty list",
			args: args{
				head: listFromSlice([]int{}),
			},
			want: listFromSlice([]int{}),
		},
		{
			name: "No duplicates",
			args: args{
				head: listFromSlice([]int{1, 2, 3}),
			},
			want: listFromSlice([]int{1, 2, 3}),
		},
		{
			name: "All duplicates same number",
			args: args{
				head: listFromSlice([]int{2, 2, 2, 2}),
			},
			want: listFromSlice([]int{2}),
		},
		{
			name: "Duplicates at end",
			args: args{
				head: listFromSlice([]int{1, 2, 3, 3, 3}),
			},
			want: listFromSlice([]int{1, 2, 3}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHead := DeleteDuplicates(tt.args.head)

			got := sliceFromList(gotHead)
			want := sliceFromList(tt.want)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("DeleteDuplicates() = %v, want %v", got, want)
			}
		})
	}
}

func listFromSlice(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func sliceFromList(head *ListNode) []int {
	var out []int
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}
