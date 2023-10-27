package main

func main() {

}

func FindMultiples(integer, limit int) []int {
	result := make([]int, 0, limit/integer)
	for i := integer; i <= limit; i += integer {
		if i%integer == 0 {
			result = append(result, i)
		}
	}
	return result
}
