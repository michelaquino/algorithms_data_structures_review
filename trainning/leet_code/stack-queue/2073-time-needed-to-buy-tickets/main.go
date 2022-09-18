package main

import "fmt"

// https://leetcode.com/problems/time-needed-to-buy-tickets/
func main() {

	// tickets := []int{2, 3, 2}
	// k := 2
	// fmt.Println(timeRequiredToBuy(tickets, k))

	// tickets = []int{5, 1, 1, 1}
	// k = 0
	// fmt.Println(timeRequiredToBuy(tickets, k))

	tickets := []int{84, 49, 5, 24, 70, 77, 87, 8}
	k := 3
	fmt.Println(timeRequiredToBuy(tickets, k))
}

func timeRequiredToBuy(tickets []int, k int) int {
	timeTaken := 0

	for {
		for i := 0; i < len(tickets); i++ {
			if tickets[i] == 0 {
				continue
			}

			timeTaken++
			tickets[i]--

			if tickets[k] == 0 {
				break
			}
		}

		if tickets[k] == 0 {
			break
		}
	}

	return timeTaken
}
