package main

import "testing"

func Test06(t *testing.T) {
	want1, want2 := "1034", "2472"
	real1, real2 := solve()
	if real1 != want1 || real2 != want2 {
		t.Fatal()
	}
}
