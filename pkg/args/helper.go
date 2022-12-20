package args

func isLastArg(args []string, index int) bool {
	return len(args)-1 == index
}
