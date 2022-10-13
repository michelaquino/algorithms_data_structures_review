package main

import (
	"testing"
)

func TestScheduleCourse(t *testing.T) {
	type testCase struct {
		input    [][]int
		expected int
		desc     string
	}

	testCases := []testCase{
		{
			input:    [][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}},
			expected: 3,
			desc:     "test 1",
		},
		{
			input:    [][]int{{1, 2}},
			expected: 1,
			desc:     "test 2",
		},
		{
			input:    [][]int{{3, 2}, {4, 3}},
			expected: 0,
			desc:     "test 3",
		},
		{
			input:    [][]int{{5, 5}, {4, 6}, {2, 6}},
			expected: 2,
			desc:     "test 4",
		},
		{
			input:    [][]int{{5, 15}, {3, 19}, {6, 7}, {2, 10}, {5, 16}, {8, 14}, {10, 11}, {2, 19}},
			expected: 5,
			desc:     "test 5",
		},
		{
			input:    [][]int{{5, 7}, {3, 5}, {10, 18}, {4, 16}, {10, 14}},
			expected: 3,
			desc:     "test 6",
		},
	}

	for _, tc := range testCases {
		result := scheduleCourse(tc.input)

		if result != tc.expected {
			t.Errorf("%s failed. want %d, got %d", tc.desc, tc.expected, result)
		}

	}
}

// func TestAbs(t *testing.T) {
// 	got := Abs(-1)
// 	if got != 1 {
// 		t.Errorf("Abs(-1) = %d; want 1", got)
// 	}
// }
