// package result implements a Result type that allows
// a similar style to "pattern-matching" features in other
// languages, which require cases to be exhaustively handled
// in match exprs.
package result

import "errors"

// Result represents the result of a computation.
// If constructed with Ok, it represents a success.
// If constructed with Error, it represents a failure.
type Result[T any] struct {
	v   *T
	err error
}

// Ok constructs a result from a value.
func Ok[T any](v T) Result[T] {
	return Result[T]{v: &v}
}

// Error constructs a Result from an error.
func Error[T any](err error) Result[T] {
	if err == nil {
		err = errors.New("")
	}
	return Result[T]{err: err}
}

// Match unpacks the Result r and executes ok if the Result
// was constructed with Ok or executes e if the Result was
// constructed with Error. The return value of whichever
// function executes will be returned from Match.
func Match[T, U any](
	r Result[T],
	ok func(v T) U,
	e func(err error) U,
) U {
	if r.err != nil {
		return e(r.err)
	}
	return ok(*r.v)
}
