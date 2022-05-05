package util

func Contains[T comparable](array []T, element T) bool {
	for i := range array {
		if array[i] == element {
			return true
		}
	}
	return false
}
