package main

import "testing"

func Test_heightChecker(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{"first", []int{1, 1, 4, 2, 1, 3}, 3},
		{"second", []int{5, 1, 2, 3, 4}, 5},
		{"three", []int{1, 2, 3, 4, 5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heightChecker(tt.heights); got != tt.want {
				t.Errorf("heightChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
