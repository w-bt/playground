package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	nums := getSlice(head)
	var newList *ListNode
	index := len(nums) - n
	isFirst := true
	for i := len(nums) - 1; i >= 0; i-- {
		if i == index {
			continue
		}
		list := ListNode{}
		if isFirst {
			list = ListNode{
				Val: nums[i],
			}
			isFirst = false
		} else {
			copyList := *newList
			list = ListNode{
				Val:  nums[i],
				Next: &copyList,
			}
		}
		newList = &list
	}
	return newList
}

func getSlice(head *ListNode) []int {
	result := []int{}
	for true {
		if head == nil {
			break
		}
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
