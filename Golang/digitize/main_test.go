package main

import (
	"reflect"
	"testing"
)

func TestDigitize(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"first", 12345, []int{5, 4, 3, 2, 1}},
		{"second", 35231, []int{1, 3, 2, 5, 3}},
		{"zero", 0, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Digitize(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Digitize() = %v, want %v", got, tt.want)
			}
		})
	}
}
