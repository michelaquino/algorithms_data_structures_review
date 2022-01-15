package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// kmpSearch("a", "abcdabeabf")
	// computePrefix("aabcadaabe")
	// fmt.Println(kmpSearch("aaa", "aa"))
	// return

	file, err := os.Open("input.txt")
	checkError(err)
	reader := bufio.NewReaderSize(file, 16*1024*1024)

	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	genesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var genes []string

	for i := 0; i < int(n); i++ {
		genesItem := genesTemp[i]
		genes = append(genes, genesItem)
	}

	healthTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var health []int32

	for i := 0; i < int(n); i++ {
		healthItemTemp, err := strconv.ParseInt(healthTemp[i], 10, 64)
		checkError(err)
		healthItem := int32(healthItemTemp)
		health = append(health, healthItem)
	}

	sTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	s := int32(sTemp)

	start := time.Now()
	min := math.Inf(1)
	max := math.Inf(-1)
	for sItr := 0; sItr < int(s); sItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		firstTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		first := int32(firstTemp)

		lastTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		last := int32(lastTemp)

		d := firstMultipleInput[2]

		geneHealthSum := dnaHealth(genes, health, d, first, last)
		if float64(geneHealthSum) < min {
			min = float64(geneHealthSum)
		}

		if float64(geneHealthSum) > max {
			max = float64(geneHealthSum)
		}
	}

	fmt.Println("main:", time.Since(start))
	fmt.Printf("%d %d", int64(min), int64(max))
}

func dnaHealth(genes []string, geneHealth []int32, d string, first, last int32) int64 {
	var sum int64 = 0

	for i := first; i <= last; i++ {
		gene := genes[i]
		// count := naiveSearch(d, gene)
		count := kmpSearch(d, gene)
		geneHealthSum := int64(count * geneHealth[i])
		sum += geneHealthSum
	}

	return sum
}

func naiveSearch(str, pattern string) int32 {
	lenPattern := len(pattern)
	j := 0
	var quantityFound int32 = 0

	startSearchIndex := 0
	for i := 0; i < len(str); i++ {

		if string(str[i]) != string(pattern[j]) {
			startSearchIndex++
			i = startSearchIndex - 1
			j = 0
			continue
		}

		j++
		// Not found
		if j < lenPattern {
			continue
		}

		quantityFound++
		j = 0
		startSearchIndex = i

		if i+1 >= len(str) || i == 0 {
			continue
		}

		if len(pattern) > 1 {
			i--
		}
	}

	return quantityFound
}

// https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm
func kmpSearch(str, pattern string) int32 {
	// start := time.Now()

	// LPS: longest proper prefix which is also suffix
	lpsTableNumber := []int{}
	for range pattern {
		lpsTableNumber = append(lpsTableNumber, 0)
	}

	for i := 0; i < len(pattern); i++ {
		subPattern := pattern[:i+1]
		max := maxLenOfPropperPrefixAndSuffixThatMatch(subPattern)
		lpsTableNumber[i] = max
	}

	// fmt.Println("str: ", str)
	// fmt.Println("pattern: ", pattern)
	// fmt.Println("lpsTableNumber: ", lpsTableNumber)

	lenPattern := len(pattern)
	j := 0
	var quantityFound int32 = 0

	for i := 0; i < len(str); i++ {
		// fmt.Println("-----------")
		// fmt.Println("i: ", i)
		// fmt.Println("j: ", j)

		// fmt.Println("string(str[i]): ", string(str[i]))
		// fmt.Println("string(pattern[j]): ", string(pattern[j]))

		// fmt.Println("str: ", str)
		// fmt.Println("pattern: ", string(pattern))
		// fmt.Println("j: ", j)
		// fmt.Println("pattern[j]: ", string(pattern[j]))
		if string(str[i]) != string(pattern[j]) {
			// fmt.Println("Mismatch")
			// Mismatch

			returnTo := lpsTableNumber[j] - 1
			if returnTo < 0 {
				continue
			}

			j = lpsTableNumber[j] - 1
			continue
		}

		// fmt.Println("Match")
		// fmt.Println("lenPattern: ", lenPattern)
		// fmt.Println("j: ", j)
		// Match
		j++
		if j < lenPattern {
			// fmt.Println("Not found yet")
			// Not found
			continue
		}

		quantityFound++
		j = 0
		if i+1 >= len(str) || i == 0 {
			continue
		}

		if len(pattern) > 1 {
			i--
		}
	}

	// fmt.Println("kmpSearch: ", time.Since(start))
	return quantityFound
}

func maxLenOfPropperPrefixAndSuffixThatMatch(str string) int {
	// start := time.Now()
	mapProperMatch := map[string]int{}

	max := 0
	for i, j := 0, len(str)-1; i <= len(str)-2 || j > 1; i, j = i+1, j-1 {
		prefix := str[:i+1]
		suffix := str[j:]

		mapProperMatch[prefix] += 1
		mapProperMatch[suffix] += 1

	}

	for subStr, value := range mapProperMatch {
		if value <= 1 {
			// No match
			continue
		}

		if len(subStr) > max {
			max = len(subStr)
		}
	}

	// fmt.Println("maxLenOfPropperPrefixAndSuffixThatMatch: ", time.Since(start))
	return max
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
