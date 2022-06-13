package main

import (
	"fmt"
	"sort"
)

func findMaxSum(numbers []int) int {
	maxIndex := len(numbers) - 1
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	return numbers[maxIndex] + numbers[maxIndex-1]
}

func main() {
	list := []int{5, 9, 7, 11}
	fmt.Println(findMaxSum(list))
}
