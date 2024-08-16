package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)



func TestAddFuncCarry(t *testing.T) {

	tests := []struct {
		name    string
		argFirst int
		argSecond int
		want    int
	}{
		{
			name:    "first",
			argFirst: 2,
			argSecond: 2,
			want:    4,
		}, {
			name:    "second",
			argFirst: 22,
			argSecond: 2,
			want:    24,
		
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			
			 got := AddCarry(tt.argFirst)(tt.argSecond); 

			require.Equal(t, tt.want, got)
	})
	}
}