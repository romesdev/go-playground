package functional

type reduceFunc[T any] func(T, T) T

func Reduce[T any](list []T, init T, fn reduceFunc[T]) T {
	current := init
	for _, value := range list {
		current = fn(current, value)
	}
	return current
}
