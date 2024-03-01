package main

import (
	"fmt"
)

func main() {
	// minimumBribes([]int32{2, 1, 5, 3, 4})
	minimumBribes([]int32{2, 5, 1, 3, 4})
}

func minimumBribes2(q []int32) {
	sumBribes := 0
	for i, actualNumber := range q {
		q2 := q[i+1:]
		bribes := 0

		for _, next := range q2 {
			if actualNumber <= next {
				continue
			}

			bribes++
			if bribes > 2 {
				fmt.Println("Too chaotic")
				return
			}
		}

		sumBribes += bribes
	}

	fmt.Println(sumBribes)
}

func minimumBribes(q []int32) {
	expectedWindow := []int32{1, 2, 3}

	sumBribes := 0
	for _, ticket := range q {
		expectedWindow = append(expectedWindow, expectedWindow[2]+1)

		if ticket == expectedWindow[0] {
			expectedWindow = expectedWindow[1:]
			continue
		}

		if ticket == expectedWindow[1] {
			sumBribes += 1
			expectedWindow = append(expectedWindow[0:1], expectedWindow[2:]...)
			continue
		}

		if ticket == expectedWindow[2] {
			sumBribes += 2
			expectedWindow = append(expectedWindow[:2], expectedWindow[3:]...)
			continue
		}

		fmt.Println("Too chaotic")
		return

	}

	fmt.Println(sumBribes)
}

/*
2, 1, 5, 3, 4

2 -> 1
2 ->


2, 5, 1, 3, 4
*/
