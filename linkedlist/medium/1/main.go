package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	numberSlice1 := convertNodeToSlice(l1)
	numberSlice2 := convertNodeToSlice(l2)
	standardizeSize(&numberSlice1, &numberSlice2)

	slice := newSlice(numberSlice1, numberSlice2)

	return convertToList(slice)
}

func standardizeSize(slice1, slice2 *[]int) {
	if len(*slice1) == len(*slice2) {
		return
	}

	var slice []int
	if len(*slice1) > len(*slice2) {
		slice = make([]int, len(*slice1))
		for index := range *slice2 {
			slice[index] = (*slice2)[index]
		}
		*slice2 = slice
	} else {
		slice = make([]int, len(*slice2))
		for index := range *slice1 {
			slice[index] = (*slice1)[index]
		}
		*slice1 = slice
	}
	return
}

func newSlice(slice1, slice2 []int) []int {
	var slice []int
	var isCarried bool
	for index := range slice1 {
		carried := 0
		if isCarried {
			carried = 1
		}
		result := slice1[index] + slice2[index] + carried
		if result > 9 {
			result = result - 10
			isCarried = true
		} else {
			isCarried = false
		}
		slice = append(slice, result)
	}
	if isCarried {
		slice = append(slice, 1)
	}
	return slice
}

func convertNodeToSlice(list *ListNode) []int {
	node := *list
	number := []int{}
	for true {
		if node.Next == nil {
			number = append(number, node.Val)
			break
		}
		number = append(number, node.Val)
		node = *node.Next
	}

	return number
}

func convertToList(slice []int) *ListNode {
	var previousNode *ListNode
	for i := len(slice) - 1; i >= 0; i-- {
		node := &ListNode{
			Val:  slice[i],
			Next: previousNode,
		}
		previousNode = node
	}
	return previousNode
}
