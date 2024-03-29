package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'swapNodes' function below.
 *
 * The function is expected to return a 2D_INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY indexes
 *  2. INTEGER_ARRAY queries
 */

func swapNodes(writer *bufio.Writer, indexes [][]int32, queries []int32) [][]int32 {
	nodeDepthMap := map[int]int{
		1: 1,
	}

	for i, index := range indexes {
		left := index[0]
		if left != -1 {
			nodeDepthMap[int(left)] = nodeDepthMap[i+1] + 1
		}

		right := index[1]
		if right != -1 {
			nodeDepthMap[int(right)] = nodeDepthMap[i+1] + 1
		}
	}

	for _, query := range queries {
		for node, _ := range indexes {
			depth := nodeDepthMap[node+1]

			// if node isn't multiple of query
			if (int32(depth))%query != 0 {
				continue
			}

			newIndex := make([]int32, 2)
			copy(newIndex, indexes[node])

			newIndex[0], newIndex[1] = newIndex[1], newIndex[0]
			indexes[node] = newIndex
		}

		printTree(writer, indexes, 1)

		writer.Flush()
		writer.WriteString("\n")
	}

	return indexes
}

func printTree(writer *bufio.Writer, tree [][]int32, depth int32) {
	if int(depth) > len(tree) {
		return
	}

	indexDepth := tree[depth-1]
	left := indexDepth[0]
	if left > 0 {
		printTree(writer, tree, left)
	}

	writer.WriteString(fmt.Sprintf("%d ", depth))

	right := indexDepth[1]
	if right > 0 {
		printTree(writer, tree, right)
	}
}

func main() {
	file, err := os.Open("input1.txt")
	defer file.Close()
	checkError(err)
	reader := bufio.NewReader(file)

	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var indexes [][]int32
	for i := 0; i < int(n); i++ {
		indexesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var indexesRow []int32
		for _, indexesRowItem := range indexesRowTemp {
			indexesItemTemp, err := strconv.ParseInt(indexesRowItem, 10, 64)
			checkError(err)
			indexesItem := int32(indexesItemTemp)
			indexesRow = append(indexesRow, indexesItem)
		}

		if len(indexesRow) != 2 {
			panic("Bad input")
		}

		indexes = append(indexes, indexesRow)
	}

	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []int32

	for i := 0; i < int(queriesCount); i++ {
		queriesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	swapNodes(writer, indexes, queries)
	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
