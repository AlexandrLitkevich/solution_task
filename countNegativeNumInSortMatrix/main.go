package main

import "fmt"

func main() {
	fmt.Println("run")
}

func countNegatives(grid [][]int) int {
	count := 0
	for _, ints := range grid {
		for j := len(ints) - 1; j >= 0; j-- {
			if ints[j] >= 0 {
				break
			}
			count++
		}

	}
	return count
}

/*
func countNegatives(grid [][]int) int {
	var n, m, negatives_counter int = len(grid), len(grid[0]), 0

	for i := 0; i < n; i++ {
		var left, right, middle int = 0, m - 1, 0
		for right >= left {
			middle = (left + right) / 2
			if grid[i][middle] >= 0 {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
		if grid[i][middle] >= 0 {
			middle += 1
		}
		if middle == 0 {
			return negatives_counter + m*(n-i)
		}
		negatives_counter += m - middle
		left, right, middle = 0, middle, 0
	}
	return negatives_counter
}



*/
