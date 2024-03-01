package main

import (
	"fmt"
	"math"
)

func main() {
	breadsMadeCharacteristiscMap := map[int][]int{
		300: {5, 1, 0},
		225: {3, 1, 1},
		75:  {1, 1, 0},
		200: {4, 0, 1},
		150: {4, 0, 0},
		50:  {2, 0, 0},
	}

	howManyBreads := howManyBreadsNeedToMadeFor(
		breadsMadeCharacteristiscMap,
		[]int{4, 1, 0},
	)

	fmt.Println("howManyBreads: ", howManyBreads)
}

func howManyBreadsNeedToMadeFor(breadsMadeCharacteristiscMap map[int][]int, newCharacteristisc []int) float64 {
	var maxClosestDistance float64 = 2.0
	var howManyBreadsSum float64 = 0
	var closestCount float64 = 0

	for howManyBreadsMade, characteristisc := range breadsMadeCharacteristiscMap {
		var sum float64 = 0
		for i := 0; i < len(characteristisc); i++ {
			sum += math.Pow(float64(characteristisc[i])-float64(newCharacteristisc[i]), 2)
		}

		distance := math.Sqrt(sum)
		if distance <= maxClosestDistance {
			howManyBreadsSum += float64(howManyBreadsMade)
			closestCount++
		}
	}

	return howManyBreadsSum / closestCount
}
