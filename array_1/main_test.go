package main

import (
	"reflect"
	"testing"
)

func TestCountPositivesSumNegatives(t *testing.T) {

	tests := []struct {
		name    string
		numbers []int
		want    []int
	}{
		{
			name:    "first",
			numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15},
			want:    []int{10, -65},
		}, {
			name:    "second",
			numbers: []int{},
			want:    []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := CountPositivesSumNegatives(tt.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountPositivesSumNegatives() = %v, want %v", got, tt.want)
			}
		})
	}
}
