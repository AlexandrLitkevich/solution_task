package main

import (
	"sort"
)

func main() {
	heightChecker([]int{1, 1, 4, 2, 1, 3})
}

func heightChecker(heights []int) int {
	sortArr := make([]int, len(heights))
	count := 0
	copy(sortArr, heights)

	sort.Slice(sortArr, func(i, j int) bool {
		return sortArr[i] < sortArr[j]
	})

	for i := range heights {
		if heights[i] != sortArr[i] {
			count++
			continue
		}

	}

	return count
}
