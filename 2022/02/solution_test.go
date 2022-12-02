package main

import "testing"

func Test02(t *testing.T) {
	want1, want2 := "12156", "10835"
	real1, real2 := solve()
	if real1 != want1 || real2 != want2 {
		t.Fatal()
	}

}
