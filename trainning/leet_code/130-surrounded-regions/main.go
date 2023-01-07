package main

import (
	"fmt"
	"reflect"
)

var (
	notProcessed              = 0
	pieceOfIslandNotProcessed = 1
	alreadyProcessed          = -1
	islandAlreadyFound        = 2
)

// https://leetcode.com/problems/surrounded-regions
func main() {
	testCases := []struct {
		board         [][]byte
		expectdOutput [][]byte
	}{
		{
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			expectdOutput: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			board: [][]byte{
				{'X', 'O', 'X', 'X'},
				{'O', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X'},
			},
			expectdOutput: [][]byte{
				{'X', 'O', 'X', 'X'},
				{'O', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'O'},
				{'O', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'O'},
				{'O', 'X', 'O', 'X'},
			},
		},
	}

	for _, t := range testCases {
		// alreadyProcessed := make([][]bool, len(t.board))
		// for i := 0; i < len(t.board); i++ {
		// 	alreadyProcessed[i] = make([]bool, len(t.board[i]))
		// }

		// if canPaint(t.board, alreadyProcessed, 3, 2) {
		// 	paint(t.board, 3, 2)
		// }

		// printGrid(t.board)
		// continue

		// fmt.Println("=========================================")
		solve(t.board)

		fmt.Println()
		printGrid(t.board)

		if !reflect.DeepEqual(t.board, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, t.board))
		}
	}
}

func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			// not edge
			if i > 0 && i < len(board)-1 && j > 0 && j < len(board[i])-1 {
				continue
			}

			if board[i][j] == 'O' {
				paintE(board, i, j)
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case 'E':
				board[i][j] = 'O'
			}
		}
	}
}

func paintE(board [][]byte, i, j int) {
	// out of limits
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[i])-1 {
		return
	}

	if board[i][j] == 'E' || board[i][j] == 'X' {
		return
	}

	board[i][j] = 'E'
	paintE(board, i-1, j)
	paintE(board, i+1, j)
	paintE(board, i, j-1)
	paintE(board, i, j+1)
	return
}

func solve_not_optimized(board [][]byte) {
	alreadyProcessed := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		alreadyProcessed[i] = make([]bool, len(board[i]))
	}

	for i := 1; i < len(board); i++ {
		for j := 1; j < len(board[i]); j++ {
			if board[i][j] == 'O' && !alreadyProcessed[i][j] && canPaint(board, alreadyProcessed, i, j) {
				paint(board, i, j)
			}
		}
	}
}

func canPaint(board [][]byte, alreadyProcessed [][]bool, i, j int) bool {
	// out of limits
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[i])-1 {
		return false
	}

	if board[i][j] == 'X' {
		return true
	}

	// board
	if i == 0 || j == 0 || i == len(board)-1 || j == len(board[i])-1 {
		return false
	}

	// if already processed
	if alreadyProcessed[i][j] {
		return true
	}

	alreadyProcessed[i][j] = true

	//up
	upOk := canPaint(board, alreadyProcessed, i-1, j)

	// down
	downOk := canPaint(board, alreadyProcessed, i+1, j)

	// left
	leftOk := canPaint(board, alreadyProcessed, i, j-1)

	// right
	rightOk := canPaint(board, alreadyProcessed, i, j+1)
	return upOk && downOk && leftOk && rightOk
}

func paint(board [][]byte, i, j int) {
	// out of limits
	if i < 0 || i > len(board) || j < 0 || j > len(board[i]) {
		return
	}

	if board[i][j] == 'X' {
		return
	}

	board[i][j] = 'X'
	paint(board, i-1, j)
	paint(board, i+1, j)
	paint(board, i, j-1)
	paint(board, i, j+1)
	return
}

func printGrid(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf("%s ", string(board[i][j]))
		}

		fmt.Println()
	}
}
