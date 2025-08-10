package main

func Generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	result := [][]int{[]int{1}}
	if numRows == 1 {
		return result
	}
	return RecursiveGenerate(numRows, 2, result)
}

func RecursiveGenerate(numRows, currentRow int, result [][]int) [][]int {
	resultRow := make([]int, currentRow)
	for i := 0; i < currentRow; i++ {
		if i == 0 {
			resultRow[i] = 1
		} else if i == currentRow-1 {
			resultRow[i] = 1
		} else {
			resultRow[i] = result[currentRow-2][i-1] + result[currentRow-2][i]
		}
	}
	result = append(result, resultRow)
	if numRows == currentRow {
		return result
	} else {
		return RecursiveGenerate(numRows, currentRow+1, result)
	}
}
