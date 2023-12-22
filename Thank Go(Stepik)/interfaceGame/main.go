package main

func main() {

}

type S1 struct {
	F1 int
	F2 string
	F3 int
}

type S2 struct {
	F1 int
	F2 string
	F3 S1
}

type S2slice []S2

func (a S2slice) Len() int {
	return len(a)
}
