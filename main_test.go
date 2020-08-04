package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 string
	}{
		// TODO: Add test cases.
		{
			name:  "1",
			want:  0,
			want1: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := A()
			if got != tt.want {
				t.Errorf("A() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("A() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}