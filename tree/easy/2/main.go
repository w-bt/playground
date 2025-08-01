package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindMode(root *TreeNode) []int {
	total := make(map[int]int)
	buildMap(root, total)
	greatest := 0
	for _, value := range total {
		if value >= greatest {
			greatest = value
		}
	}

	result := []int{}
	for key, value := range total {
		if value == greatest {
			result = append(result, key)
		}
	}

	return result
}

func buildMap(root *TreeNode, total map[int]int) {
	total[root.Val] += 1
	if root.Left != nil {
		buildMap(root.Left, total)
	}

	if root.Right != nil {
		buildMap(root.Right, total)
	}
}
