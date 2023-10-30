package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("start  solution")
}
func Digitize(n int) []int {
	convertStr := strconv.Itoa(n)

	var res []int
	for i := len(convertStr) - 1; i >= 0; i-- {
		in, _ := strconv.Atoi(string(convertStr[i]))
		res = append(res, in)

	}
	return res
}
