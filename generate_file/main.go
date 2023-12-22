package main

import "fmt"

func main() {
	w, err := MakePasswordInFile()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("this w", w)

	str, err := ReadPasswordInFile()
	if err != nil {
		fmt.Println("this error", err)
	}
	fmt.Println("this str", str)
}
