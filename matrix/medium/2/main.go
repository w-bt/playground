package main

func SpiralOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0 {
		return result
	}

	mode := 1
	indexMode1 := 0
	indexMode2 := len(matrix[0]) - 1
	indexMode3 := len(matrix) - 1
	indexMode4 := 0
	flag := true
	for flag {
		switch mode {
		case 1:
			if indexMode1 > indexMode3 {
				flag = false
				break
			}
			for j := indexMode4; j <= indexMode2; j++ {
				result = append(result, matrix[indexMode1][j])
			}
			mode += 1
			indexMode1 += 1
		case 2:
			if indexMode2 < indexMode4 {
				flag = false
				break
			}
			for j := indexMode1; j <= indexMode3; j++ {
				result = append(result, matrix[j][indexMode2])
			}
			mode += 1
			indexMode2 -= 1
		case 3:
			if indexMode1 > indexMode3 {
				flag = false
				break
			}
			for j := indexMode2; j >= indexMode4; j-- {
				result = append(result, matrix[indexMode3][j])
			}
			mode += 1
			indexMode3 -= 1
		case 4:
			if indexMode2 < indexMode4 {
				flag = false
				break
			}
			for j := indexMode3; j >= indexMode1; j-- {
				result = append(result, matrix[j][indexMode4])
			}
			mode = 1
			indexMode4 += 1
		}
	}
	return result
}
