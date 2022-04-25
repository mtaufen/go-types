package option

import "testing"

func TestOption(t *testing.T) {
	value := Some(6)
	nothing := None[int]()

	// Pattern matching style might be a bit nicer if Go had
	// named parameters in function calls, since the names
	// would show inside Get, but this still does the job
	// and with practice you'd get used to knowing that
	// the match for Some goes in the second arg and the
	// match for None goes in the third.
	some := func(v int) bool {
		return v > 5
	}
	none := func() bool {
		return false
	}

	if Get(value, some, none) == false {
		t.Errorf("Get Some: got false, want true")
	}

	if Get(nothing, some, none) == true {
		t.Errorf("Get None: got true, want false")
	}
}
