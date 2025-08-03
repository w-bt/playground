package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var firstNode *ListNode
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil {
			if result == nil {
				result = l2
				firstNode = result
			} else {
				(*result).Next = l2
				result = l2
			}
			l2 = l2.Next
			continue
		}
		if l2 == nil {
			if result == nil {
				result = l1
				firstNode = result
			} else {
				(*result).Next = l1
				result = l1
			}
			l1 = l1.Next
			continue
		}
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				if result == nil {
					result = l1
					firstNode = result
				} else {
					(*result).Next = l1
					result = l1
				}
				l1 = l1.Next
			} else {
				if result == nil {
					result = l2
					firstNode = result
				} else {
					(*result).Next = l2
					result = l2
				}
				l2 = l2.Next
			}
		}
	}
	if result != nil {
		result.Next = nil
	}
	return firstNode
}
