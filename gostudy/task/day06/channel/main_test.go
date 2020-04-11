package main

import "testing"

func Test_makeInt(t *testing.T) {
	type args struct {
		c chan int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			makeInt(tt.args.c)
		})
	}
}
