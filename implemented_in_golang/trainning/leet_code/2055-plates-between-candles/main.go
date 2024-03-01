package main

import (
	"fmt"
	"reflect"
	"sort"
)

// https://leetcode.com/problems/plates-between-candles/
func main() {
	testCases := []struct {
		s             string
		queries       [][]int
		expectdOutput []int
	}{
		{
			s: "**|**|***|",
			queries: [][]int{
				// {2, 5},
				// {5, 9},
				{2, 9},
			},
			expectdOutput: []int{2, 3},
		},
		// {
		// 	s: "***|**|*****|**||**|*",
		// 	queries: [][]int{
		// 		{1, 17},
		// 		{4, 5},
		// 		{14, 17},
		// 		{5, 11},
		// 		{15, 16},
		// 	},
		// 	expectdOutput: []int{9, 0, 0, 0, 0},
		// },
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := platesBetweenCandles(t.s, t.queries)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func platesBetweenCandles(s string, queries [][]int) []int {
	candles := make([]int, 0, len(s))

	for i, ch := range s {
		if ch == '|' {
			candles = append(candles, i)
		}
	}

	fmt.Println("candles: ", candles)
	awnser := make([]int, 0, len(queries))

	for _, query := range queries {
		fmt.Println("---")
		fmt.Println("query: ", query)
		startIndex := query[0]
		endIndex := query[1]

		leftCandleIndex := sort.Search(len(candles), func(i int) bool {
			return candles[i] >= startIndex
		})

		rightCandleIndex := sort.Search(len(candles), func(i int) bool {
			return candles[i] >= endIndex
		})

		if rightCandleIndex <= leftCandleIndex {
			awnser = append(awnser, 0)
			continue
		}

		fmt.Println("leftCandleIndex: ", leftCandleIndex)
		fmt.Println("rightCandleIndex: ", rightCandleIndex)
		fmt.Println("candles[rightCandleIndex]: ", candles[rightCandleIndex])
		fmt.Println("candles[leftCandleIndex]: ", candles[leftCandleIndex])
		fmt.Println()
		fmt.Println("candles[rightCandleIndex] - candles[leftCandleIndex]: ", candles[rightCandleIndex]-candles[leftCandleIndex])
		fmt.Println("rightCandleIndex - leftCandleIndex: ", rightCandleIndex-leftCandleIndex)

		count := (candles[rightCandleIndex] - candles[leftCandleIndex]) - (rightCandleIndex - leftCandleIndex)
		awnser = append(awnser, count)

	}

	return awnser
}

func platesBetweenCandlesWrong(s string, queries [][]int) []int {
	awnser := []int{}
	for _, q := range queries {
		startIndex, endIndex := q[0], q[1]
		awnser = append(awnser, numberOfItemsOn(s, startIndex, endIndex))
	}

	return awnser
}

func numberOfItemsOn(s string, startIndex, endIndex int) int {
	if startIndex > int(len(s)) || endIndex > int(len(s)) {
		return 0
	}

	var count int = 0
	var firstPipeIndex int = -1
	var endPipeIndex int = -1

	var i int
	for i = startIndex; i <= endIndex; i++ {
		if firstPipeIndex != -1 && endPipeIndex != -1 {
			count += endPipeIndex - firstPipeIndex - 1
			firstPipeIndex = endPipeIndex
			endPipeIndex = -1
		}

		if i >= len(s) || s[i] == '*' {
			continue
		}

		if firstPipeIndex == -1 {
			firstPipeIndex = i
			continue
		}

		if endPipeIndex == -1 {
			endPipeIndex = i
			continue
		}
	}

	if firstPipeIndex != -1 && endPipeIndex != -1 {
		count += endPipeIndex - firstPipeIndex - 1
	}

	return count
}

// func numberOfItems(s string, startIndices []int32, endIndices []int32) []int32 {
// 	awnser := []int32{}
// 	for i := range startIndices {
// 		startIndex := startIndices[i]
// 		endIndex := endIndices[i]

// 		awnser = append(awnser, numberOfItemsOn(s, startIndex, endIndex))
// 	}

// 	return awnser
// }

// func numberOfItemsOn(s string, startIndex, endIndex int32) int32 {
// 	if startIndex > int32(len(s)) || endIndex > int32(len(s)) {
// 		return 0
// 	}

// 	var count int32 = 0
// 	var firstPipeIndex int32 = -1
// 	var endPipeIndex int32 = -1

// 	var i int32
// 	for i = startIndex - 1; i < endIndex; i++ {
// 		if firstPipeIndex != -1 && endPipeIndex != -1 {
// 			count += endPipeIndex - firstPipeIndex - 1
// 			startIndex = endIndex
// 			endIndex = -1
// 		}

// 		if s[i] == '*' {
// 			continue
// 		}

// 		if firstPipeIndex == -1 {
// 			firstPipeIndex = i
// 			continue
// 		}

// 		if endPipeIndex == -1 {
// 			endPipeIndex = i
// 			continue
// 		}
// 	}

// 	if firstPipeIndex != -1 && endPipeIndex != -1 {
// 		count += endPipeIndex - firstPipeIndex - 1
// 	}

// 	return int32(count)
// }

/*
Build a meeting reservation system
Given a time range, the system needs to reserve a room if exists a n available room on that range or returns error
- Input
    - Time range
        - Could be timestamp
        - monday 10AM - friday 6PM

- Output
    - Returns the ID of the room
    - Error

Given a time range
    - If there are available room in that range
        - returns the ID
    - Else
        - Returns an error

    getAnAvailabilityRoom(startTimestamp, endTimestamp int64) string
        rooms := [room1, room2, room4, room5]

    reserveTheRoom(roomID string, startTimestamp, endTimestamp int64)


type  Room struct  { // room1
    ID string
    ReservedTime [Schedule2, Schedule4, Schedule5]
}

type Schedule struct {
    startTime int64
    endTime int64
}


middle.start >= newSchedule.end -- focus on the sstart->middle array
middle.start < newSchedule.end
    middle.end < newSchedule.start
    middle.end > newSchedule.start -- focus on the middle->end array
    middle.end == newSchedule.start





[s1-2, s5-7 s3-4, s7-9, s10-11]
[s1-2, s3-4]


*/

type schedule struct {
	start int64
	end   int64
}

func isAvailable(schedules []schedule, startRange, endRange int64) bool {
	if len(schedules) == 0 {
		return true
	}

	if startRange >= endRange {
		return false
	}

	middleScheduleIndex := len(schedules) / 2
	middleSchedule := schedules[middleScheduleIndex]
	if middleSchedule.start >= endRange {
		return isAvailable(schedules[:middleScheduleIndex], startRange, endRange)
	}

	if middleSchedule.end < startRange {
		return isAvailable(schedules[:middleScheduleIndex], startRange, endRange)
	}

	return false
}
