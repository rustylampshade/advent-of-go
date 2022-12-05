package main

import "testing"

func Test04(t *testing.T) {
	want1, want2 := "RLFNRTNFB", "MHQTLJRLB"
	real1, real2 := solve()
	if real1 != want1 || real2 != want2 {
		t.Fatal()
	}
}
