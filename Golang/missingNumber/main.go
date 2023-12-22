package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Run main")
}

func missingNumber(nums []int) int {
	slices.Sort(nums)
	if nums[0] != 0 {
		return 0
	}

	n := len(nums)
	if nums[n-1] != n {
		return n
	}
	res := 0

	for i := 1; i < len(nums); i++ {
		if (nums[i] - nums[i-1]) != 1 {
			res = nums[i-1] + 1
		}
	}
	return res
}
