package main

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
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
			name: "Both lists are empty",
			args: args{
				l1: nil,
				l2: nil,
			},
			want: nil,
		},
		{
			name: "First list is empty",
			args: args{
				l1: nil,
				l2: &ListNode{Val: 0},
			},
			want: &ListNode{Val: 0},
		},
		{
			name: "Second list is empty",
			args: args{
				l1: &ListNode{Val: 1},
				l2: nil,
			},
			want: &ListNode{Val: 1},
		},
		{
			name: "Normal case with merge",
			args: args{
				l1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
				l2: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			},
			want: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 1, Next: &ListNode{
						Val: 2, Next: &ListNode{
							Val: 3, Next: &ListNode{
								Val: 4, Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "No overlapping, one list entirely smaller",
			args: args{
				l1: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
				l2: &ListNode{Val: 3, Next: &ListNode{Val: 4}},
			},
			want: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 2, Next: &ListNode{
						Val: 3, Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
