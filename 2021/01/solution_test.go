package main

import "testing"

func Test01(t *testing.T) {
	want1, want2 := "1121", "1065"
	real1, real2 := solve()
	if want1 != real1 || want2 != real2 {
		t.Fatal()
	}
}
