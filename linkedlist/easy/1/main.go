package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	hashMap := make(map[interface{}]bool)
	for true {
		if head == nil {
			break
		}
		_, ok := hashMap[head]
		if ok {
			return true
		}
		hashMap[head] = true
		head = head.Next
	}
	return false
}
