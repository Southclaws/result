package result

// Result describes a type that may either be some value or an error.
type Result[T any] interface {
	Valid() bool
	Value() T
	Error() error
}

// result is a hidden type that implements the Result interface. It's basically
// a discriminated union where `e != nil` is the discriminator.
type result[T any] struct {
	v T
	e error
}

// Valid implements Result
func (r result[T]) Valid() bool { return r.e == nil }

// Valid implements Result
func (r result[T]) Value() T { return r.v }

// Valid implements Result
func (r result[T]) Error() error { return r.e }

// Wrap takes a typical Go function that returns (T, error) and wraps it in a
// result type. This is useful for converting from the normal Go pattern of
// (T, error) to a wrapped result type. It is used like so:
//
//     r := Wrap(SomeFunction())
//     if r.Valid() {
//         // do something with r.Value()
//     }
//
func Wrap[T any](t T, err error) Result[T] {
	if err != nil {
		return result[T]{e: err}
	} else {
		return result[T]{t, nil}
	}
}
