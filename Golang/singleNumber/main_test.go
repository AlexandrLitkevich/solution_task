package main

import "testing"

func Test_singleNumber(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"1", []int{2, 2, 1}, 1},
		{"2", []int{4, 1, 2, 1, 2}, 4},
		{"once", []int{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
