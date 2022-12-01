package main

import "testing"

func Test01(t *testing.T) {
	want1, want2 := "71502", "208191"
	real1, real2 := solve()
	if real1 != want1 || real2 != want2 {
		t.Fatal()
	}

}
