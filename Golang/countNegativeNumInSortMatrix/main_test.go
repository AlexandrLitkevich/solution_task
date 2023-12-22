package main

import "testing"

func Test_countNegatives(t *testing.T) {

	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "first",
			grid: [][]int{
				[]int{4, 3, 2, -1},
				[]int{3, 2, 1, -1},
				[]int{1, 1, -1, -2},
				[]int{-1, -1, -2, -3},
			},
			want: 8,
		},
		{
			name: "second",
			grid: [][]int{
				[]int{3, 2},
				[]int{1, 0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNegatives(tt.grid); got != tt.want {
				t.Errorf("countNegatives() = %v, want %v", got, tt.want)
			}
		})
	}
}
