package main

import "fmt"

func WordSearch(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return false
	}

	takenCoord := make(map[string]bool)
	firstChar := word[0]
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == firstChar {
				if len(word) == 1 {
					return true
				}
				takenCoord[fmt.Sprintf("%d:%d", i, j)] = true
				found := CheckNext(board, word[1:len(word)], i, j, takenCoord)
				if found {
					return true
				}
			}
		}
	}

	return false
}

func CheckNext(board [][]byte, nextWord string, i, j int, takenCoord map[string]bool) bool {
	if len(nextWord) == 0 {
		return false
	}

	possibleMoves := fetchPossibleMoves(board, i, j, takenCoord)
	firstChar := nextWord[0]
	for _, possibleMove := range possibleMoves {
		if board[possibleMove[0]][possibleMove[1]] == firstChar {
			if len(nextWord) == 1 {
				return true
			}
			takenCoord[fmt.Sprintf("%d:%d", possibleMove[0], possibleMove[1])] = true
			found := CheckNext(board, nextWord[1:len(nextWord)], possibleMove[0], possibleMove[1], takenCoord)
			if found {
				return true
			}
		}
	}

	return false
}

func fetchPossibleMoves(board [][]byte, i, j int, takenCoord map[string]bool) (result [][]int) {
	if i > 0 {
		if _, ok := takenCoord[fmt.Sprintf("%d:%d", i-1, j)]; !ok {
			result = append(result, []int{i - 1, j})
		}
	}

	if i < len(board)-1 {
		if _, ok := takenCoord[fmt.Sprintf("%d:%d", i+1, j)]; !ok {
			result = append(result, []int{i + 1, j})
		}
	}

	if j > 0 {
		if _, ok := takenCoord[fmt.Sprintf("%d:%d", i, j-1)]; !ok {
			result = append(result, []int{i, j - 1})
		}
	}

	if j < len(board[0])-1 {
		if _, ok := takenCoord[fmt.Sprintf("%d:%d", i, j+1)]; !ok {
			result = append(result, []int{i, j + 1})
		}
	}

	return result
}
