package main

import "testing"

func Test07(t *testing.T) {
	want1, want2 := "1317827", "1117448"
	real1, real2 := solve()
	if real1 != want1 || real2 != want2 {
		t.Fatal()
	}
}
