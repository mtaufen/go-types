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

// Match unpacks the Option o and executes some if the Option
// was constructed with Some or executes none if the Option
// was constructed with None. The return value of whichever
// function executes will be returned from Match.
func Match[T, U any](
	o Option[T],
	some func(T) U,
	none func() U,
) U {
	if o.v == nil {
		return none()
	}
	return some(*o.v)
}

// TODO: Consider manually currying these helpers. Would
// make it possible to create convenient, reusable thunks
// that can then be applied to different Options.

// Map applies the func f to the value in the Option o,
// and returns a new Option containing the result. If o
// is None, then the returned Option is also None.
func Map[T, U any](f func(T) U, o Option[T]) Option[U] {
	return Match(o,
		// Some v
		func(v T) Option[U] {
			return Some(f(v))
		},
		None[U])
}

// Apply applies the func in the Option f to the value in
// the Option o, and returns a new Option containing the
// result. If either Option is None, the returned Option
// is also None.
func Apply[T, U any](
	f Option[func(T) U],
	o Option[T],
) Option[U] {
	return Match(f,
		func(f func(T) U) Option[U] {
			return Match(o,
				func(v T) Option[U] {
					return Some(f(v))
				},
				None[U])
		},
		None[U])
}

// Bind applies the func f to the value in the Option o,
// and returns a new Option containing the result. If o
// is None, the returned Option is also None.
func Bind[T, U any](
	o Option[T],
	f func(T) Option[U],
) Option[U] {
	return Match(o, f, None[U])
}

// TODO: Consider manually currying Do
// Kind of an experiment. Not quite a monad (return types
// can't change across the pipeline) but maybe still useful
// for sequences of transformations across the same type.
// If we make it a package func instead of a method, then
// we can allow the type to change once, if fs remains
// variadic, or N times, where N is the number of function
// arguments (and we'd implement Do1, Do2, Do3, Do4, etc,
// which would be ugly but potentially useful).
func (o Option[T]) Do(
	fs ...func(v T) Option[T],
) Option[T] {
	var done bool
	for _, f := range fs {
		o = Match(o,
			f,
			func() Option[T] {
				done = true
				return None[T]()
			})
		if done {
			return o
		}
	}
	return o
}
