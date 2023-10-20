package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		str    string
		ending string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"first", args{"abc", "bc"}, true},
		{"second", args{"abc", "d"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.str, tt.args.ending); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
