package main

import "testing"

func TestMakePasswordInFile(t *testing.T) {
	tests := []struct {
		name    string
		want    bool
		wantErr bool
	}{
		{"first", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakePasswordInFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("MakePasswordInFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MakePasswordInFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
