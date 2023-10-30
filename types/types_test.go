package mytypes_test

import (
	mytypes "happy-fun-book/types"
	"testing"
)

func TestMyInt_Twice(t *testing.T) {
	t.Parallel()
	input := mytypes.MyInt(5)
	want := mytypes.MyInt(10)
	got := input.Twice()
	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestMyStringLen(t *testing.T) {
	t.Parallel()
	input := mytypes.MyStringLen("Hello")
	want := 5
	got := input.Len()
	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}
