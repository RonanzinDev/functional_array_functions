package types

type Array[T any] struct {
	data []T
}

func NewArray[M any](data []M) Array[M] {
	arr := Array[M]{data}
	return arr
}

func (a *Array[T]) Map(f func(T) T) []T {
	n := make([]T, len(a.data))
	for i, e := range a.data {
		n[i] = f(e)
	}
	return n
}

func (a *Array[T]) Filter(f func(T) bool) []T {
	var n []T
	for i := range a.data {
		if f(a.data[i]) {
			n = append(n, a.data[i])
		}
	}
	return n
}
func (a *Array[T]) FilterStructs(f func(T) bool) T {
	var n T
	for i := range a.data {
		if f(a.data[i]) {
			n = a.data[i]
		}
	}
	return n
}
