package main

import "testing"

func Test08(t *testing.T) {
	want1, want2 := "1647", "392080"
	real1, real2 := solve()
	if real1 != want1 || real2 != want2 {
		t.Fatal()
	}
}
