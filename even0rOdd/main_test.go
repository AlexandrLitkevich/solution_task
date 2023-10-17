package main

import "testing"

func TestEvenOrOdd(t *testing.T) {

	tests := []struct {
		name   string
		number int
		want   string
	}{
		{
			name:   "first even",
			number: 2,
			want:   "Even",
		}, {
			name:   "first odd",
			number: 3,
			want:   "Odd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvenOrOdd(tt.number); got != tt.want {
				t.Errorf("EvenOrOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
