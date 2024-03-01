package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func (n Node) valAsInt() int {
	if n.Val {
		return 1
	}

	return 0
}

func (n Node) isLeafAsInt() int {
	if n.IsLeaf {
		return 1
	}

	return 0
}

func (n Node) string() string {
	return fmt.Sprintf("[%d %d]", n.isLeafAsInt(), n.valAsInt())
}

func construct(grid [][]int) *Node {
	if len(grid) == 0 {
		return nil
	}

	root := Node{
		Val: true,
	}

	if isLeaf(grid) {
		value := grid[0][0]
		if value == 0 {
			root.Val = false
		} else {
			root.Val = true
		}

		root.IsLeaf = true
		return &root
	}

	subGrid1, subGrid2, subGrid3, subGrid4 := splitQuadrants(grid)
	root.TopLeft = construct(subGrid1)
	root.TopRight = construct(subGrid2)
	root.BottomLeft = construct(subGrid3)
	root.BottomRight = construct(subGrid4)

	return &root
}

func splitQuadrants(grid [][]int) ([][]int, [][]int, [][]int, [][]int) {
	subGrid1 := [][]int{}
	subGrid2 := [][]int{}
	subGrid3 := [][]int{}
	subGrid4 := [][]int{}

	for i := 0; i < len(grid); i++ {
		row1 := []int{}
		row2 := []int{}
		row3 := []int{}
		row4 := []int{}

		for j := 0; j < len(grid[i]); j++ {

			if i < len(grid)/2 && j < len(grid)/2 {
				row1 = append(row1, grid[i][j])
				continue
			}

			if i < len(grid)/2 && j >= len(grid)/2 {
				row2 = append(row2, grid[i][j])
				continue
			}

			if i >= len(grid)/2 && j < len(grid)/2 {
				row3 = append(row3, grid[i][j])
				continue
			}

			if i >= len(grid)/2 && j >= len(grid)/2 {
				row4 = append(row4, grid[i][j])
				continue
			}
		}

		if len(row1) > 0 {
			subGrid1 = append(subGrid1, row1)
		}

		if len(row2) > 0 {
			subGrid2 = append(subGrid2, row2)
		}

		if len(row3) > 0 {
			subGrid3 = append(subGrid3, row3)
		}

		if len(row4) > 0 {
			subGrid4 = append(subGrid4, row4)
		}
	}

	return subGrid1, subGrid2, subGrid3, subGrid4
}

func isLeaf(grid [][]int) bool {
	if len(grid) == 1 {
		return true
	}

	if len(grid) == 0 {
		return false
	}

	valueToCompare := grid[0][0]
	for _, row := range grid {
		for _, value := range row {
			if value != valueToCompare {
				return false
			}
		}
	}

	return true
}

func printLevel(node *Node, levelToPrint, currentLevel int) string {
	if node == nil || currentLevel < 1 {
		return "null"
	}

	if currentLevel == levelToPrint {
		return node.string()
	}

	array := []string{}
	tl := printLevel(node.TopLeft, levelToPrint, currentLevel-1)
	if tl != "" {
		array = append(array, tl)
	}

	tr := printLevel(node.TopRight, levelToPrint, currentLevel-1)
	if tr != "" {
		array = append(array, tr)
	}

	bl := printLevel(node.BottomLeft, levelToPrint, currentLevel-1)
	if bl != "" {
		array = append(array, bl)
	}

	br := printLevel(node.BottomRight, levelToPrint, currentLevel-1)
	if br != "" {
		array = append(array, br)
	}

	return strings.Join(array, ",")
}

func print(root *Node) {
	if root == nil {
		return
	}

	output := strings.Builder{}
	output.WriteString("[")
	levels := height(root)
	for level := levels; level > 0; level-- {
		output.WriteString(printLevel(root, level, levels))
		if level > 1 {
			output.WriteString(",")
		}
	}

	output.WriteString("]")
}

func height(node *Node) int {
	if node == nil {
		return 0
	}

	calcMax := func(numbers []int) int {
		max := -1
		for _, n := range numbers {
			if n > max {
				max = n
			}
		}

		return max
	}

	topLeftHeight := height(node.TopLeft)
	topRight := height(node.TopRight)
	bottomLeft := height(node.BottomLeft)
	bottomRight := height(node.BottomRight)

	return calcMax([]int{topLeftHeight, topRight, bottomLeft, bottomRight}) + 1
}

func main() {
	grid := [][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}
	// grid := [][]int{
	// 	{0, 1},
	// 	{1, 0},
	// }
	// construct(grid)

	root := construct(grid)
	// levels := height(root)
	print(root)
}
