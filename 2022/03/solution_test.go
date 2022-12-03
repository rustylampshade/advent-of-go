package main

import "testing"

func Test03(t *testing.T) {
	want1, want2 := "8105", "2363"
	real1, real2 := solve()
	if real1 != want1 || real2 != want2 {
		t.Fatal()
	}

}
