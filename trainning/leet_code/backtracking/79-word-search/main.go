package main

import "fmt"

// https://leetcode.com/problems/word-search

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	// word := "SEE"
	// word := "ABCB"
	// board := [][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}
	// word := "AAB"

	anwser := exist(board, word)
	fmt.Println("anwser: ", anwser)
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			found := search(board, i, j, 0, word, "")
			if found {
				return true
			}
		}
	}

	return false
}

func search(board [][]byte, i, j, letterPos int, word, alreadyFound string) bool {
	if alreadyFound == word {
		return true
	}

	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[i])-1 {
		return false
	}

	letterToSearch := word[letterPos]
	if board[i][j] != letterToSearch {
		return false
	}

	aux := board[i][j]
	board[i][j] = 0

	alreadyFound = fmt.Sprintf("%s%s", alreadyFound, string(letterToSearch))

	letterPos++
	found := false
	found = search(board, i+1, j, letterPos, word, alreadyFound)
	if found {
		return found
	}

	found = search(board, i-1, j, letterPos, word, alreadyFound)
	if found {
		return found
	}

	found = search(board, i, j+1, letterPos, word, alreadyFound)
	if found {
		return found
	}

	found = search(board, i, j-1, letterPos, word, alreadyFound)
	if found {
		return found
	}

	board[i][j] = aux
	return false
}
