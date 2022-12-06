package shared

import (
	"testing"
)

func TestMax(t *testing.T) {
	give := []int{1, 2, 3}
	want := 3
	_, real := Max(give)
	if want != real {
		t.Fatal()
	}
}

func TestMin(t *testing.T) {
	give := []int{1, 2, 3}
	want := 1
	_, real := Min(give)
	if want != real {
		t.Fatal()
	}
}

func TestFindAllInt(t *testing.T) {
	give := []int{1, 0, 0, 1}
	want := []int{0, 3}
	real := FindAll(give, 1)
	for i, v := range real {
		if v != want[i] {
			t.Fatal()
		}
	}
}

func TestFindAllString(t *testing.T) {
	give := []string{"a", "", "b", ""}
	want := []int{1, 3}
	real := FindAll(give, "")
	for i, v := range real {
		if v != want[i] {
			t.Fatal()
		}
	}
}

func TestFindAllMissing(t *testing.T) {
	give := []string{"a", "", "b", ""}
	defer func() {
		if r := recover(); r == nil {
			t.Fatal()
		}
	}()

	FindAll(give, "missing")
}

func TestMapInt(t *testing.T) {
	give := []int{1, 2, 3}
	want := []int{2, 3, 4}
	f := func(elem int) int {
		return elem + 1
	}
	real := Map(give, f)
	for i, v := range real {
		if v != want[i] {
			t.Fatal()
		}
	}
}

func TestMapString(t *testing.T) {
	give := []string{"ab", "bc", "cd"}
	want := []string{"b", "c", "d"}
	f := func(elem string) string {
		return elem[1:]
	}
	real := Map(give, f)
	for i, v := range real {
		if v != want[i] {
			t.Fatal()
		}
	}
}
