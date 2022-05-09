package args

var KeyPairFlags = map[string]bool{
	"-q": false,
}

func ParseKeyPairFlags() {
	parseFlags(KeyPairFlags)
}

func parseFlags(flags map[string]bool) {
	for _, arg := range args {
		if _, ok := flags[arg]; ok {
			flags[arg] = true
		}
	}
}
