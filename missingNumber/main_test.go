package main

import "testing"

func Test_missingNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "first",
			nums: []int{3, 0, 1},
			want: 2,
		},
		{
			name: "second",
			nums: []int{0, 1},
			want: 2,
		},
		{
			name: "third",
			nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			want: 8,
		},
		{
			name: "four",
			nums: []int{3, 2, 1},
			want: 0,
		},
		{
			name: "five",
			nums: []int{0, 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
