package main

func Rotate(matrix [][]int) {
	dupMatrix := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dupMatrix[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dupMatrix[i][j] = matrix[i][j]
		}
	}
	for i := 0; i < len(dupMatrix); i++ {
		for j := 0; j < len(dupMatrix[i]); j++ {
			matrix[j][len(dupMatrix)-1-i] = dupMatrix[i][j]
		}
	}
}
