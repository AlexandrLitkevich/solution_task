package main

import (
	"fmt"
)

func main() {
	fmt.Println(-30 + -30)
}

func CountPositivesSumNegatives(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	store := map[string]int{
		"positive": 0,
		"negative": 0,
	}

	for _, number := range numbers {
		if number > 0 {
			val, _ := store["positive"]
			store["positive"] = val + 1

		} else {
			val, _ := store["negative"]
			store["negative"] = val + number
		}

	}

	negative, _ := store["negative"]
	positive, _ := store["positive"]

	res := []int{positive, negative}
	return res
}
