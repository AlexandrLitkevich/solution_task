package main

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "1",
			arr:  []int{1, -1, 2, -2, 3},
			want: 3,
		}, {
			name: "2",
			arr:  []int{-3, 1, 2, 3, -1, -4, -2},
			want: -4,
		}, {
			name: "3",
			arr:  []int{1, -1, 2, -2, 3, 3},
			want: 3,
		}, {
			name: "4",
			arr:  []int{-110, 110, -38, -38, -62, 62, -38, -38, -38},
			want: -38,
		}, {
			name: "5",
			arr:  []int{-9, -105, -9, -9, -9, -9, 105},
			want: -9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.arr); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
