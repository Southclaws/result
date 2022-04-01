package result

// Ternary simply returns `a` or `b` based on `cond`.
func Ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	} else {
		return b
	}
}

// Ternary executes either `a` or `b` based on `cond` and returns its value.
func TernaryFn[T any](cond bool, a, b func() T) T {
	if cond {
		return a()
	} else {
		return b()
	}
}

// TernaryResult returns either `a` or `b` as a result type based on `cond`.
// This is useful for picking function calls that return (T, error) like so:
//
//     result := TernaryResult(
//         x > 1,
//         Wrap(A()),
//         Wrap(B()),
//     )
//
// It's cleaner compared to the idiomatic mutable variable approach:
//
//     var result T
//     var err error
//     if x > 1 {
//         result, err = A()
//     } else {
//         result, err = B()
//     }
//
func TernaryResult[T any](cond bool, a, b Result[T]) Result[T] {
	if cond {
		return a
	} else {
		return b
	}
}
