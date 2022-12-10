package main

import "testing"

func Test10(t *testing.T) {
	want1, want2 := "6098", "2597"
	real1, real2 := solve()
	if real1 != want1 || real2 != want2 {
		t.Fatal()
	}
}
