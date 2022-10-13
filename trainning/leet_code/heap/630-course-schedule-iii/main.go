package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/course-schedule-iii/
// courses[i] = [durationi, lastDayi]
func main() {
	courses := [][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}
	// courses := [][]int{{1, 2}}
	// courses := [][]int{{3, 2}, {4, 3}}
	// courses := [][]int{{5, 5}, {4, 6}, {2, 6}}
	fmt.Println("result: ", scheduleCourse(courses))
}

func scheduleCourse(courses [][]int) int {
	sort.SliceStable(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	count := 0
	time := 0
	for i := 0; i < len(courses); i++ {
		if courses[i][0]+time <= courses[i][1] {
			time += courses[i][0]
			count++
		} else {
			maxIndex := i
			for j := 0; j < i; j++ {
				if courses[j][0] > courses[maxIndex][0] {
					maxIndex = j
				}
			}

			if courses[maxIndex][0] > courses[i][0] {
				time -= courses[maxIndex][0] - courses[i][0]
			}

			courses[maxIndex][0] = -1
		}
	}

	return count
}
