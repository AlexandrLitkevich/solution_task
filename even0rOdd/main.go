package main

func main() {

}

func EvenOrOdd(number int) string {
	res := number % 2
	if res == 0 {
		return "Even"
	}
	return "Odd"
}
