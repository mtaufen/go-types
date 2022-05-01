package option

import (
	"fmt"
	"testing"
)

func TestOption(t *testing.T) {
	value := Some(1)
	nothing := None[int]()

	// Pattern matching style might be a bit nicer if Go had
	// named parameters in function calls, since the names
	// would show inside Use call, but this still does the job
	// and with practice you'd get used to knowing that the
	// match for Some goes in the second arg and the match for
	// None goes in the third. Plus the function signatures
	// make it pretty clear inline anyway.
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

func TestMap(t *testing.T) {
	intToString := func(v int) string {
		return fmt.Sprintf("%d", v)
	}

	value := Map(intToString, Some(1))
	nothing := Map(intToString, None[int]())

	some := func(v string) string {
		return v
	}
	none := func() string {
		return ""
	}

	if v := Match(value, some, none); v != "1" {
		t.Errorf(`Some: got %s, want "1"`, v)
	}

	if v := Match(nothing, some, none); v != "" {
		t.Errorf(`None: got %s, want empty string ""`, v)
	}
}

func TestApply(t *testing.T) {
	// TODO
}

func TestBind(t *testing.T) {
	// TODO
}

func TestPipe(t *testing.T) {
	value := Some(1)
	nothing := None[int]()

	ib := NewPipe(func(i int) bool {
		return i != 0
	})
	bs := NewPipe(func(b bool) string {
		return fmt.Sprintf("%t", b)
	})

	pipeline := func(i Option[int]) Option[string] {
		ib.In <- i
		bs.In <- <-ib.Out
		return <-bs.Out
	}

	if v := Match(pipeline(value),
		Id[string], Zero[string]); v != "true" {
		t.Errorf(`Some: got %s, want "true"`, v)
	}

	if v := Match(pipeline(nothing),
		Id[string], Zero[string]); v != "" {
		t.Errorf(`None: got %s, want ""`, v)
	}
}
