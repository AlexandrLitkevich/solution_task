package sortStrArr

import (
	"fmt"
	"sort"
)

func main() {
	arr := []string{"Telescopes", "Glasses", "Eyes", "Monocles"}
	fmt.Println(SortByLength(arr))
}

func SortByLength(arr []string) []string {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []string
	for i := 1; i < len(arr); i++ {
		if len(pivot) > len(arr[i]) {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}

	}
	result := append(SortByLength(left), pivot)
	result = append(result, SortByLength(right)...)
	return result
}

func SortByLengthV2(arr []string) []string {
	sort.SliceStable(arr, func(i, j int) bool { return len(arr[i]) < len(arr[j]) })
	return arr
}
