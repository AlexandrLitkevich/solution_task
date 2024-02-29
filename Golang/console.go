package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Fscan(os.Stdin, name)

	fmt.Println("this input console", name)
}
