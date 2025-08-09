package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	var node *ListNode
	var first *ListNode
	for {
		if head == nil {
			break
		}
		if node == nil {
			node = &ListNode{Val: head.Val}
			first = node
			head = head.Next
		} else {
			if node.Val != head.Val {
				node.Next = &ListNode{Val: head.Val}
				node = node.Next
			}
			head = head.Next
		}
	}
	return first
}
