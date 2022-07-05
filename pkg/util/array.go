package util

func Contains(elem string, arr []string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			return true
		}
	}
	return false
}
