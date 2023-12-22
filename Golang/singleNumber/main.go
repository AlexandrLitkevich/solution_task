package main

import (
	"fmt"
)

func main() {
	fmt.Println("run single number")
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := 0

	m := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = struct{}{}
			continue
		}
		delete(m, num)
	}

	for key := range m {
		res = key

	}

	return res
}
