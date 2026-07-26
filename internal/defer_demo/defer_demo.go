package deferdemo

func DeferLifo() (stack string) {
	defer func() { stack += "1" }()
	defer func() { stack += "2" }()
	defer func() { stack += "3" }()

	return
}

func DeferDouble(args int) (result int) {
	defer func() { result *= 2 }()
	return args
}

func DeferSnapshot() (captured int) {
	i := 10

	defer func(snapshot int) {
		captured = snapshot
	}(i)

	i = 99

	return
}
