package main

import "testing"

func TestGetPlanetName(t *testing.T) {

	tests := []struct {
		name string
		ID   int
		want string
	}{
		{"Mercury", 1, "Mercury"},
		{"Venus", 2, "Venus"},
		{"Uranus", 7, "Uranus"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPlanetName(tt.ID); got != tt.want {
				t.Fatalf("GetPlanetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
