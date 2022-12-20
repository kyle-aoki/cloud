package args

func PollOrEmpty() string {
	if len(programArgs) == 0 {
		return ""
	}
	arg := programArgs[0]
	programArgs = programArgs[1:]
	return arg
}

func PollOrPanic() string {
	if len(programArgs) == 0 {
		panic("not enough arguments. try lab --help")
	}
	arg := programArgs[0]
	programArgs = programArgs[1:]
	return arg
}

func CollectOrEmpty() []string {
	collection := programArgs[:]
	programArgs = []string{}
	return collection
}

func CollectOrPanic() []string {
	if len(programArgs) == 0 {
		panic("not enough arguments. try lab --help")
	}
	remainingArgs := programArgs[:]
	programArgs = []string{}
	return remainingArgs
}
