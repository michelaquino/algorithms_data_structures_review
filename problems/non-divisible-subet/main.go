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

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("input_simple.txt")
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

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var s []int32

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int32(sItemTemp)
		s = append(s, sItem)
	}

	result := nonDivisibleSubset(k, s)

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

func nonDivisibleSubset(k int32, s []int32) int32 {
	if k == 0 || k == 1 {
		return k
	}

	mapRemainderByKCount := make(map[int32]int32, k)
	for i := 0; int32(i) < k; i++ {
		mapRemainderByKCount[int32(i)] = 0
	}

	remaindersByK := []int32{}
	for _, number := range s {
		remainderByK := number % k
		mapRemainderByKCount[remainderByK]++

		remaindersByK = append(remaindersByK, remainderByK)
	}

	var numbersAllowed int32 = int32(math.Min(float64(mapRemainderByKCount[0]), 1.0))
	var i int32
	for i = 1; i < (k/2)+1; i++ {
		if i == k-i {
			continue
		}

		max := math.Max(float64(mapRemainderByKCount[i]), float64(mapRemainderByKCount[k-i]))
		numbersAllowed += int32(max)
	}

	if k%2 == 0 {
		numbersAllowed++
	}

	return numbersAllowed
}
