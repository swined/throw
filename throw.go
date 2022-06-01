package throw

func E(e error) {
	if e != nil {
		panic(e)
	}
}

func VE[T any](v T, e error) T {
	if e == nil {
		return v
	} else {
		panic(e)
	}
}
