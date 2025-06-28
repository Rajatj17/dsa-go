package countingsubsequence

func CountSubTD(a string, b string, m int, n int) int {
	if (m == -1 && n == -1) || n == -1 {
		return 1
	}

	if m == -1 {
		return 0
	}

	if a[m] == b[n] {
		return CountSubTD(a, b, m-1, n-1) + CountSubTD(a, b, m-1, n)
	}

	return CountSubTD(a, b, m-1, n)
}

// func CountSubBU(a string, b string, m int, n int) int {

// }
