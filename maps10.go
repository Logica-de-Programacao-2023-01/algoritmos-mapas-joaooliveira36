package main

import (
	"fmt"
)

func countPairs(nums []int) map[[2]int]int {
	pairMap := make(map[[2]int]int)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			pair := [2]int{nums[i], nums[j]}
			pairMap[pair]++
		}
	}

	return pairMap
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 2, 4, 1, 3}
	pairMap := countPairs(nums)

	for pair, count := range pairMap {
		fmt.Printf("%v: %d\n", pair, count)
	}
}
