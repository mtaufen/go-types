package option

import "testing"

func TestOption(t *testing.T) {
	value := Some(6)
	nothing := None[int]()

	// Pattern matching style might be a bit nicer if Go had
	// named parameters in function calls, since the names
	// would show inside Use call, but this still does the job
	// and with practice you'd get used to knowing that the
	// match for Some goes in the second arg and the match for
	// None goes in the third. Plus the function signatures
	// make it pretty clear inline anyway.
	some := func(v int) bool {
		return v > 5
	}
	none := func() bool {
		return false
	}

	if Use(value, some, none) == false {
		t.Errorf("Use Some: got false, want true")
	}

	if Use(nothing, some, none) == true {
		t.Errorf("Use None: got true, want false")
	}
}
