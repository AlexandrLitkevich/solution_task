package main

import (
	"reflect"
	"testing"
)

func TestFindMultiples(t *testing.T) {
	type args struct {
		integer int
		limit   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{
				2, 10,
			},
			want: []int{2, 4, 6, 8, 10},
		},
		{
			name: "second",
			args: args{
				5, 25,
			},
			want: []int{5, 10, 15, 20, 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMultiples(tt.args.integer, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMultiples() = %v, want %v", got, tt.want)
			}
		})
	}
}
