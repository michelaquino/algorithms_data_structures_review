package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func manipulateArray(array []string, queries [][]int) (string, string) {
	for _, query := range queries {
		// fmt.Println()
		// fmt.Println("query: ", query)
		// fmt.Println("array: ", array)

		queryType := query[0]
		start := query[1] - 1
		end := query[2] - 1

		var subArray = make([]string, len(array[start:end+1]))
		copy(subArray, array[start:end+1])

		var prefix = make([]string, len(array[:start]))
		copy(prefix, array[:start])

		var suffix = make([]string, len(array[end+1:]))
		copy(suffix, array[end+1:])

		// fmt.Println("subArray: ", subArray)
		// fmt.Println("prefix: ", prefix)
		// fmt.Println("suffix: ", suffix)

		if queryType == 1 {
			array = append(subArray, append(prefix, suffix...)...)
			continue
		}

		array = append(append(prefix, suffix...), subArray...)
	}

	firstPosition, err := strconv.ParseInt(array[0], 10, 64)
	checkError(err)
	lastPosition, err := strconv.ParseInt(array[len(array)-1], 10, 64)
	difference := int64(math.Abs(float64(firstPosition - lastPosition)))

	return strconv.FormatInt(difference, 10), strings.Join(array, " ")
}

func main() {
	file, err := os.Open("testcase0.txt")
	defer file.Close()
	checkError(err)
	reader := bufio.NewReader(file)
	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)
	// defer stdout.Close()
	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	line, _, err := reader.ReadLine()
	checkError(err)

	lineSplited := strings.Split(string(line), " ")
	// n, err := strconv.Atoi(lineSplited[0]) // array lengh
	// checkError(err)

	m, err := strconv.Atoi(lineSplited[1]) // number of queries
	checkError(err)

	arrayLine, _, err := reader.ReadLine()
	checkError(err)

	array := strings.Split(string(arrayLine), " ")

	queries := [][]int{}
	for i := 0; i < m; i++ {
		queryLine, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		checkError(err)

		queryLineSplited := strings.Split(string(queryLine), " ")

		value0, err := strconv.Atoi(queryLineSplited[0])
		checkError(err)
		value1, err := strconv.Atoi(queryLineSplited[1])
		checkError(err)
		value2, err := strconv.Atoi(queryLineSplited[2])
		checkError(err)

		query := make([]int, 3)
		query[0] = value0
		query[1] = value1
		query[2] = value2

		queries = append(queries, query)
	}

	difference, arrayResult := manipulateArray(array, queries)
	fmt.Fprintf(writer, "%s\n", difference)
	fmt.Fprintf(writer, "%s\n", arrayResult)
	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
