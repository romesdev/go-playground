package functional

type filterFunc[T any] func(T) bool

func Filter[T any](list []T, fn filterFunc[T]) []T {
	result := []T{}
	for _, value := range list {
		if fn(value) {
			result = append(result, value)
		}
	}
	return result
}
