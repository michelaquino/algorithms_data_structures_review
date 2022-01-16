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
 * Complete the 'arrayManipulation' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

/*
10, 4

2 6 8
3 5 7
1 8 1
5 9 15


*/

func arrayManipulation1(n int32, queries [][]int32) int64 {
	array := make([]int64, n)

	// fmt.Println("array: ", array)

	var max int64 = -1
	for _, query := range queries {
		// fmt.Println("query[0]: ", query[0])
		// fmt.Println("query[1]: ", query[1])
		// fmt.Println("query[2]: ", query[2])

		for i := query[0] - 1; i < query[1]; i++ {
			array[i] += int64(query[2])
			// fmt.Println("array: ", array)

			if array[i] > max {
				max = array[i]
			}
		}
	}

	// fmt.Println("array: ", array)
	return max
}

func arrayManipulation2(n int32, queries [][]int32) int64 {
	mapColumns := make(map[int32]int64, n)
	for i := 0; int32(i) < n; i++ {
		mapColumns[int32(i)] = 0
	}

	for i := 0; i < len(queries); i++ {
		// fmt.Println()
		// fmt.Println("mapColumns: ", mapColumns)

		leftIndex := queries[i][0] - 1
		// fmt.Println("leftIndex: ", leftIndex)
		rightIndex := queries[i][1] - 1
		// fmt.Println("rightIndex: ", rightIndex)
		value := int64(queries[i][2])
		// fmt.Println("value: ", value)

		// fmt.Println("len(mapColumns): ", len(mapColumns))
		for key := range mapColumns {
			// fmt.Println("key: ", key)

			// se ta contido
			if key >= leftIndex && key <= rightIndex {
				mapColumns[key] += value
			}
		}
	}

	// fmt.Println("mapColumns: ", mapColumns)
	var maxSum int64 = -1
	for _, value := range mapColumns {
		if value > maxSum {
			maxSum = value
		}
	}

	return maxSum
}

// Final solution
func arrayManipulation(n int32, queries [][]int32) int64 {
	array := make([]int64, n)

	for i := 0; i < len(queries); i++ {
		leftIndex := int(queries[i][0] - 1)
		rightIndex := int(queries[i][1] - 1)
		value := int64(queries[i][2])

		array[leftIndex] += value
		if rightIndex+1 < len(array) {
			array[rightIndex+1] -= value
		}
	}

	var maxSum int64 = -1
	var sum int64 = 0
	for _, value := range array {
		sum += value

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	file, err := os.Open("input_test.txt")
	checkError(err)
	reader := bufio.NewReaderSize(file, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 3 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

	fmt.Fprintf(writer, "%d\n", result)

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
