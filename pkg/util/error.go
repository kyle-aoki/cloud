package util

func MustExec(err error) {
	if err != nil {
		panic(err)
	}
}
