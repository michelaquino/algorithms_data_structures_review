package main

import (
	"fmt"
	"math"
)

func main() {
	pastMovingRating := map[string][]int{
		"person-1": {3, 4, 4, 1, 4},
		"person-2": {4, 3, 5, 1, 5},
		"person-3": {2, 5, 1, 3, 1},
	}

	person := "person-1"
	closestPerson := personClosestOf(pastMovingRating, person, pastMovingRating[person])
	fmt.Println("closestPerson: ", closestPerson)
}

func personClosestOf(movingRating map[string][]int, person string, ratings []int) string {
	distances := map[string]float64{}

	closestDistance := math.Inf(1)
	closestPerson := ""
	for person2, ratings2 := range movingRating {
		if person == person2 {
			continue
		}

		distance, ok := distances[fmt.Sprintf("%s-%s", person, person2)]
		if !ok {
			var sum float64 = 0
			for i := 0; i < len(ratings); i++ {
				sum += math.Pow(float64(ratings[i])-float64(ratings2[i]), 2)
			}

			distance = math.Sqrt(sum)
			distances[fmt.Sprintf("%s-%s", person, person2)] = distance
			distances[fmt.Sprintf("%s-%s", person2, person)] = distance
		}

		if distance < closestDistance {
			closestDistance = distance
			closestPerson = person2
		}
	}

	return closestPerson
}
