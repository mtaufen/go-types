package result

import (
	"errors"
	"fmt"
	"testing"
)

func TestResult(t *testing.T) {
	success := Ok(6)
	err := Error[int](errors.New("oops!"))

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
