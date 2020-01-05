package main

import "testing"

func Test_bonuses(t *testing.T) {

	tests := []struct {
		name string
		sales int
		want int
	}{
		// TODO: Add test cases.
		{"Have bonus",12_000,100},
		{"Have bonus",15_000,250},
		{"No bonus",8_000,0},


	}
	for _, tt := range tests {
		got :=bonuses(tt.sales)
		if got!=tt.want{
			t.Error("for bonus test: ",tt.name,"got: ",got,"want: ",tt.want)
		}

	}
}