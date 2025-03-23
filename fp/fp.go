package fp

// Map applies a function to each element of a slice and returns a new slice.
func Map[T any, R any](input []T, mapper func(T) R) []R {
	result := make([]R, len(input))
	for i, v := range input {
		result[i] = mapper(v)
	}
	return result
}

// Filter returns a new slice containing only the elements for which predicate returns true.
func Filter[T any](input []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range input {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce aggregates elements of a slice using an accumulator function and an initial value.
func Reduce[T any, R any](input []T, initial R, reducer func(R, T) R) R {
	result := initial
	for _, v := range input {
		result = reducer(result, v)
	}
	return result
}
