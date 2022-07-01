package util

func AtLeastOneTrue(bools []bool) bool {
	for i := range bools {
		if bools[i] {
			return true
		}
	}
	return false
}