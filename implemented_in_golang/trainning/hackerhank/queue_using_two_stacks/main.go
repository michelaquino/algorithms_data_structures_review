package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writter := bufio.NewWriter(os.Stdout)
	firstLine := readLine(reader)

	queue := []int{}

	queryQuantity := parseInt(firstLine)
	for query := 0; query < int(queryQuantity); query++ {
		line := readLine(reader)
		querySplited := strings.Split(line, " ")

		queryType, value := 0, 0
		if len(querySplited) > 1 {
			queryType, value = parseInt(querySplited[0]), parseInt(querySplited[1])
		} else {
			queryType = parseInt(querySplited[0])
		}

		switch queryType {
		case 1:
			queue = append(queue, value)
		case 2:
			if len(queue) == 0 {
				continue
			}

			queue = queue[1:]
		case 3:
			if len(queue) == 0 {
				continue
			}

			firstValue := queue[0]
			writter.WriteString(fmt.Sprintf("%d\n", firstValue))
		}
	}

	writter.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err != nil {
		return ""
	}

	return strings.Trim(string(str), "\n\r")
}

func parseInt(numberStr string) int {
	numberInt, err := strconv.Atoi(numberStr)
	checkError(err)

	return numberInt
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
