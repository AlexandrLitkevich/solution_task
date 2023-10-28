package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, -1, 2, -2, 3}
	Solve2(arr)
	Solve(arr)

}

func Solve2(arr []int) int {
	resp := 0
	result := make(map[int]struct{}, len(arr))

	for i := 0; i < len(arr); i++ {
		if _, ok := result[arr[i]]; ok {
			continue
		}
		result[arr[i]] = struct{}{}
	}
	fmt.Println("this result map 1", result)

	for key := range result {
		if key < 0 {
			positiveKey := int(math.Abs(float64(key)))
			if _, ok := result[positiveKey]; ok {
				delete(result, key)
				delete(result, positiveKey)
			}
		}
	}

	for key := range result {
		resp = key
		break
	}

	return resp
}

func Solve(arr []int) int {
	hash := make(map[int]bool)
	ans := 0
	for _, entry := range arr {
		if _, ok := hash[entry]; !ok {
			hash[entry] = true
			ans += entry // складываем все которых уникальные числа
		}
	}
	return ans
}
