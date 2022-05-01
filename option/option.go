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

// Ideally we'd have some syntactic sugar for chaining
// Bind calls.
// Since Go doesn't allow type parameters in methods,
// we can't just do this:
//   func (m M[T]) Bind[T, U any](func(T) M[U]) M[U]
// which would then allow us to do this:
//   m.Bind(f).Bind(g).Bind(h)...
// which would basically be enough.
// So instead we need to do this:
//   Bind(Bind(Bind(m, f), g) h)
// Consider a function Pipe that composes two functions,
// where Pipe f g = g f. Then we could potentially have:
//   Bind(m, Pipe(Pipe(f, g), h))
// which at least decouples writing the functional pipeline
// from calling Bind, but is still rather verbose. If Go
// had a syncatic sugar for function composition like +,
// then we could have something like:
//   Bind(m, h + g + f)
// Alternatively, if Go had a syntactic sugar like => for
// Bind, then we could do something like:
//   m => f => g => h
// which gives a clear pipeline.

// What if we had something that took a func and returned
// an input channel and an output channel, where the input
// side was M[T] and the output M[U]? How far can that
// get us?

type Pipe[T, U any] struct {
	In  chan<- Option[T]
	Out <-chan Option[U]
}

func NewPipe[T, U any](
	f func(T) U,
) Pipe[T, U] {
	in := make(chan Option[T])
	out := make(chan Option[U])
	run := func() {
		for o := range in {
			out <- Map(f, o)
		}
		close(out)
	}
	go run()
	return Pipe[T, U]{in, out}
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

// TODO: move these to a separate package.

// Id is the identity func.
func Id[T any](v T) T {
	return v
}

// Zero is the zero val func.
func Zero[T any]() T {
	var zero T
	return zero
}
