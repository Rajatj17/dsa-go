package checkodd

func IsOdd(x int) bool {
	return x&1 == 1
}

func IsEvent(x int) bool {
	return x&1 == 0
}
