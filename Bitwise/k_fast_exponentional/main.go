package fastexponentional

func FastExpo(a int, n int) {
	ans := 1

	for {
		if n <= 0 {
			break
		}

		last_bit := (n & 1)
		if last_bit == 1 {
			ans = ans * a
		}

		a = a * a

		n = n >> 1
	}
}
