package main

import "fmt"

func main() {	
	fmt.Println("run func")
}

func AddCarry(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}