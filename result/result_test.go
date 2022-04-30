package result

import (
	"errors"
	"fmt"
	"testing"
)

func TestResult(t *testing.T) {
	success := Ok(6)
	err := Error[int](errors.New("oops!"))

	// Pattern matching style might be a bit nicer if Go had
	// named parameters in function calls, since the names
	// would show inside Use call, but this still does the job
	// and with practice you'd get used to knowing that the
	// match for Ok goes in the second arg and the match for
	// Error goes in the third. Plus the function signatures
	// make it pretty clear inline anyway.
	ok := func(v int) string {
		return fmt.Sprintf("%d", v)
	}
	e := func(err error) string {
		return err.Error()
	}

	if s := Match(success, ok, e); s != "6" {
		t.Errorf(`Ok: got %q, want "6"`, s)
	}

	if s := Match(err, ok, e); s != "oops!" {
		t.Errorf(`Error: got %q, want oops!`, s)
	}
}
