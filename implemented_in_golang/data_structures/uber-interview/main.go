package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/plates-between-candles/
func main() {
	testCases := []struct {
		schedules     []schedule
		startRange    int64
		endRange      int64
		expectdOutput bool
	}{
		{
			schedules: []schedule{
				{
					start: 1,
					end:   2,
				},
				{
					start: 3,
					end:   4,
				},
				{
					start: 7,
					end:   9,
				},
				{
					start: 10,
					end:   11,
				},
			},
			startRange:    5,
			endRange:      7,
			expectdOutput: true,
		},
		{
			schedules: []schedule{
				{
					start: 1,
					end:   2,
				},
				{
					start: 3,
					end:   4,
				},
				{
					start: 7,
					end:   9,
				},
				{
					start: 10,
					end:   11,
				},
			},
			startRange:    5,
			endRange:      8,
			expectdOutput: false,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := isAvailable(t.schedules, t.startRange, t.endRange)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

/*
s5-7

[s1-2, s3-4, s7-9, s10-11]
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
