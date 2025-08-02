package utils

// SumSlice get slice of any type and return sum of its elements using provided func.
func SumSlice[T any](s []T, f func(t T) int) int {
	result := 0
	for _, element := range s {
		result += f(element)
	}
	return result
}
