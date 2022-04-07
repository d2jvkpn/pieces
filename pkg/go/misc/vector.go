package misc

type Vector[T any] []T

func (v Vector[T]) First() *T {
	if len(v) == 0 {
		return nil
	}
	return &v[0]
}

func (v Vector[T]) Last() *T {
	var n = len(v)

	if n == 0 {
		return nil
	}
	return &v[n-1]
}
