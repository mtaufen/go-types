package option

import "testing"

func TestOption(t *testing.T) {
	value := Some(1)
	nothing := None[int]()

	some := func(v int) int {
		return v
	}
	none := func() int {
		return -1
	}

	if v := Match(value, some, none); v != 1 {
		t.Errorf("Some: got %d, want 1", v)
	}

	if v := Match(nothing, some, none); v != -1 {
		t.Errorf("None: got %d, want -1", v)
	}
}
