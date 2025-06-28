package countbits

func CountBits(n int) int {
	count := 0

	for {
		if n <= 0 {
			break
		}

		if n&1 == 1 {
			count++
		}

		n = n >> 1
	}

	return count
}
