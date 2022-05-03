package misc

type Result[T any] struct {
	Data *T
	Err  error
}

func Ok[T any](d *T) Result[T] {
	return Result[T]{Data: d}
}

func Err[T any](err error) Result[T] {
	return Result[T]{Err: err}
}

func (result *Result[T]) Ok() bool {
	return result.Err == nil
}
