package main

import (
	"testing"
)

func TestMaxPerformance(t *testing.T) {
	type testCase struct {
		speed      []int
		efficiency []int
		n          int
		k          int
		expected   int
		desc       string
	}

	testCases := []testCase{
		{
			speed:      []int{2, 10, 3, 1, 5, 8},
			efficiency: []int{5, 4, 3, 9, 7, 2},
			n:          6,
			k:          2,
			expected:   60,
			desc:       "test 1",
		},
	}

	for _, tc := range testCases {
		result := maxPerformance(tc.n, tc.speed, tc.efficiency, tc.k)

		if result != tc.expected {
			t.Errorf("%s failed. want %d, got %d", tc.desc, tc.expected, result)
		}

	}
}
