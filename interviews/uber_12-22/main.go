package main

import (
	"fmt"
	"reflect"
)

type class struct {
	start int64
	end   int64
}

func main() {
	testCases := []struct {
		input         []class
		newClass      class
		expectdOutput bool
	}{
		{
			input: []class{
				{start: 1, end: 2},
				{start: 2, end: 3},
				{start: 5, end: 6},
				{start: 7, end: 8},
				{start: 9, end: 10},
			},
			newClass:      class{start: 4, end: 5},
			expectdOutput: false,
		},
		{
			input: []class{
				{start: 1, end: 2},
				{start: 2, end: 3},
				{start: 5, end: 6},
				{start: 7, end: 8},
				{start: 9, end: 10},
			},
			newClass:      class{start: 5, end: 7},
			expectdOutput: true,
		},
		{
			input: []class{
				{start: 1, end: 2},
				{start: 2, end: 3},
				{start: 5, end: 6},
				{start: 7, end: 8},
				{start: 9, end: 10},
			},
			newClass:      class{start: 0, end: 1},
			expectdOutput: false,
		},
		{
			input: []class{
				{start: 1, end: 2},
				{start: 2, end: 3},
				{start: 5, end: 6},
				{start: 7, end: 8},
				{start: 9, end: 10},
			},
			newClass:      class{start: 10, end: 11},
			expectdOutput: false,
		},
		{
			input: []class{
				{start: 1, end: 2},
				{start: 2, end: 3},
				{start: 5, end: 6},
				{start: 7, end: 8},
				{start: 9, end: 10},
			},
			newClass:      class{start: 4, end: 5},
			expectdOutput: false,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := willOverlap(t.input, 0, len(t.input)-1, t.newClass)
		fmt.Printf("output: '%v'\n", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted '%v', got '%v'", t.expectdOutput, output))
		}
	}
}

func willOverlap(classes []class, startIndex, endIndex int, newClass class) bool {
	if len(classes) == 0 || startIndex > endIndex {
		return false
	}

	pivotIndex := (startIndex + endIndex) / 2
	classPivot := classes[pivotIndex]

	if newClass.start >= classPivot.end {
		return willOverlap(classes, pivotIndex+1, endIndex, newClass)
	} else if newClass.end <= classPivot.start {
		return willOverlap(classes, startIndex, pivotIndex-1, newClass)
	}

	return true
}
