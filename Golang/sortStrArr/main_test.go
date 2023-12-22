package sortStrArr

import (
	"reflect"
	"testing"
)

func TestSortByLength(t *testing.T) {

	tests := []struct {
		name string
		arr  []string
		want []string
	}{
		{"first", []string{"Telescopes", "Glasses", "Eyes", "Monocles"}, []string{"Eyes", "Glasses", "Monocles", "Telescopes"}},
		{"Second", []string{"Telescopes"}, []string{"Telescopes"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortByLength(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortByLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
