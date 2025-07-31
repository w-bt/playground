package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	level := 0
	result := [][]int{[]int{root.Val}}
	result = getList(root.Left, level+1, result)
	result = getList(root.Right, level+1, result)
	return result
}

func getList(t *TreeNode, level int, result [][]int) [][]int {
	if t == nil {
		return result
	}
	if level > len(result)-1 {
		result = append(result, []int{t.Val})
		result = getList(t.Left, level+1, result)
		result = getList(t.Right, level+1, result)
	} else {
		result[level] = append(result[level], t.Val)
		result = getList(t.Left, level+1, result)
		result = getList(t.Right, level+1, result)
	}
	return result
}
