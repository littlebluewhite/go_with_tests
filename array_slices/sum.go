package array_slices

import (
	"fmt"
)

func Sum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func SumAll(n1 ...[]int) []int {
	fmt.Println(n1)
	var sums []int
	for _, numbers := range n1 {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, numbers := range numbers {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return sums
}
