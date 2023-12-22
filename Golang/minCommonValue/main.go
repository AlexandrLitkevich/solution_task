package main

import (
	"slices"
)

func main() {
	f := []int{1, 2, 3}
	f2 := []int{2, 4}
	getCommon(f, f2)
}

func getCommon(nums1 []int, nums2 []int) int {
	minCount := -1
	for i := 0; i < len(nums1); i++ {
		n, found := slices.BinarySearch(nums2, nums1[i])
		if found && minCount < 0 {
			minCount = nums2[n]
			continue
		}

		if found && minCount > 0 {
			if nums2[n] < minCount {
				minCount = nums2[n]
			}
		}
	}

	return minCount
}
