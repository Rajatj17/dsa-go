package ithbitcrud

func GetIthBit(n int, i int) int {
	mask := (1 << i)

	if n&mask > 0 {
		return 1
	}

	return 0
}

func SetIthBit(n int, i int) int {
	mask := (1 << i)

	n = n | mask

	return n
}

func ClearIthBit(n int, i int) int {
	mask := ^(1 << i)

	n = n & mask

	return n
}

func UpdateIthBit(n int, i int, v int) int {
	n = ClearIthBit(n, i)

	mask := (v << i)

	n = n | mask

	return n
}

func ClearLastIBits(n int, i int) int {
	mask := (-1 << i)

	n = n & mask

	return n
}
