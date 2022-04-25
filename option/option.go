// package option implements an Option type that allows
// a similar style to "pattern-matching" features in other
// languages, which require cases to be exhaustively handled
// in match exprs.
package option

// Option represents the possibility of a value.
// If constructed with Some, it represents a real value.
// If constructed with None, it represents the absence of
// a value.
type Option[T any] struct {
	v *T
}

// Some constructs an Option from a real value.
func Some[T any](v T) Option[T] {
	return Option[T]{&v}
}

// None constructs an Option that represents the absence
// of a value.
func None[T any]() Option[T] {
	return Option[T]{nil}
}

// Use unpacks the Option o and executes some if the Option
// was constructed with Some or executes none if the Option
// was constructed with None. The return value of whichever
// function executes will be returned from Use.
func Use[T, U any](
	o Option[T],
	some func(v T) U,
	none func() U,
) U {
	if o.v == nil {
		return none()
	}
	return some(*o.v)
}
