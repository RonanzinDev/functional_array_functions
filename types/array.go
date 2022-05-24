package types

type Array[T any] struct {
	Data []T
}

func NewArray[M any](data []M) Array[M] {
	arr := Array[M]{data}
	return arr
}

func (a *Array[T]) Map(f func(T) T) []T {
	n := make([]T, len(a.Data))
	for i, e := range a.Data {
		n[i] = f(e)
	}
	return n
}

func (a *Array[T]) Filter(f func(T) bool) []T {
	var n []T
	for i := range a.Data {
		if f(a.Data[i]) {
			n = append(n, a.Data[i])
		}
	}
	return n
}
func (a *Array[T]) FilterStructs(f func(T) bool) T {
	var n T
	for i := range a.Data {
		if f(a.Data[i]) {
			n = a.Data[i]
		}
	}
	return n
}

func (a *Array[T]) Append(v T) {
	n := make([]T, len(a.Data)+1)
	copy(n, a.Data)
	// for i := range a.Data {
	// 	n[i] = a.Data[i]
	// }
	n[len(n)-1] = v
	a.Data = n
}
