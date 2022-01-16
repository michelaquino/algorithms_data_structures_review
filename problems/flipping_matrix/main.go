package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int32{
		{112, 42, 83, 119},
		{56, 125, 56, 49},
		{15, 78, 101, 43},
		{62, 98, 114, 108}}

	fmt.Println(flippingMatrix(matrix))
}

func flippingMatrix(matrix [][]int32) int32 {
	length := len(matrix)
	n := int(length / 2)

	/*
		Primeiro olha para o ultimo quadrante e troca os maiores numeros de lugar com os menores
	*/
	for i := n; i < length; i++ {
		for j := n; j < length; j++ {

			reverseRow1 := i
			reverseColumn1 := 2*n - j - 1

			reverseRow2 := 2*n - i - 1
			reverseColumn2 := j

			value := float64(matrix[i][j])
			reverseValue1 := float64(matrix[reverseRow1][reverseColumn1])
			reverseValue2 := float64(matrix[reverseRow2][reverseColumn2])

			if value <= reverseValue1 && value <= reverseValue2 {
				continue
			}

			if reverseValue1 < reverseValue2 {
				matrix[i][j] = int32(reverseValue1)
				matrix[reverseRow1][reverseColumn1] = int32(value)
				continue
			}

			matrix[i][j] = int32(reverseValue2)
			matrix[reverseRow2][reverseColumn2] = int32(value)
		}
	}

	// Olha somente para o primeiro quadrante, adicionando à soma
	// o maior numero entre o valor atual e os que podem ser trocados
	var sum int32 = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			reverseRow1 := i
			reverseColumn1 := 2*n - j - 1
			reverseValue1 := float64(matrix[reverseRow1][reverseColumn1])

			reverseRow2 := 2*n - i - 1
			reverseColumn2 := j
			reverseValue2 := float64(matrix[reverseRow2][reverseColumn2])

			value := float64(matrix[i][j])

			max1 := math.Max(value, reverseValue1)
			max2 := math.Max(value, reverseValue2)
			max3 := math.Max(max1, max2)

			sum += int32(max3)
		}
	}

	return sum
}
