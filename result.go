package result

type Result[T any] interface {
	Valid() bool
	Value() T
	Error() error
}

type result[T any] struct {
	v T
	e error
}

func (r result[T]) Valid() bool  { return r.e == nil }
func (r result[T]) Value() T     { return r.v }
func (r result[T]) Error() error { return r.e }

func Wrap[T any](t T, err error) Result[T] {
	if err != nil {
		return result[T]{e: err}
	} else {
		return result[T]{t, nil}
	}
}

func TernaryResult[T Result[T]](cond bool, a, b T) Result[T] {
	if cond {
		return a
	} else {
		return b
	}
}

func Ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	} else {
		return b
	}
}
