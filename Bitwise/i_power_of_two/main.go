package poweroftwo

func IsPowerOfTwo(n int) bool {
	// Property for pwoer of 2
	return (n & (n - 1)) == 0
}
