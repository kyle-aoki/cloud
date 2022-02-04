package util

func StrPtr(s string) *string {
	return &s
}

func IntToInt64Ptr(i int) *int64 {
	i64 := int64(i)
	return &i64
}

func StrSlicePtr(s []string) (sptrs []*string) {
	for i := range s {
		sptrs = append(sptrs, &s[i])
	}
	return sptrs
}
