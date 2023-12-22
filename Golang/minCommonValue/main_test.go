package main

import "testing"

func Test_getCommon(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first success",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{2, 4},
			},
			want: 2,
		},
		{
			name: "second success",
			args: args{
				nums1: []int{1, 2, 3, 6},
				nums2: []int{2, 3, 4, 5},
			},
			want: 2,
		},
		{
			name: "three success",
			args: args{
				nums1: []int{1, 2, 8, 12, 32, 34, 43, 52, 57, 62, 64, 67, 71, 71, 79, 81, 86, 91, 93, 94},
				nums2: []int{9, 11, 12, 14, 19, 25, 29, 31, 38, 39, 41, 42, 58, 63, 66, 70, 71, 73, 91, 91},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCommon(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("getCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}
