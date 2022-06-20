package util

import (
	"fmt"
	"strconv"
)

func StrPtr(s string) *string {
	return &s
}

func BoolPtr(b bool) *bool {
	return &b
}

func IntToInt64Ptr(i int) *int64 {
	i64 := int64(i)
	return &i64
}

func IntToString(i int) string {
	return fmt.Sprintf("%d", i)
}

func StringToInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)
	MustExec(err)
	return int(i)
}

func StrSlicePtr(s []string) (sptrs []*string) {
	for i := range s {
		sptrs = append(sptrs, &s[i])
	}
	return sptrs
}
