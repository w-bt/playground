package main

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Remove 2nd node from end of [1,2,3,4,5]",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
				n:    2,
			},
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}},
		},
		{
			name: "Remove 1st node from [1]",
			args: args{
				head: &ListNode{1, nil},
				n:    1,
			},
			want: nil,
		},
		{
			name: "Remove 1st node from [1,2]",
			args: args{
				head: &ListNode{1, &ListNode{2, nil}},
				n:    1,
			},
			want: &ListNode{1, nil},
		},
		{
			name: "Remove head (n = 2) from [1,2,3]",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
				n:    3,
			},
			want: &ListNode{2, &ListNode{3, nil}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
