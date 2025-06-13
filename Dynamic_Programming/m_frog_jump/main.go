package frogjump

import "math"

func getCost(jump1, jump2 int) int {
	cost := jump2 - jump1
	if cost < 1 {
		return -1 * cost
	}

	return cost
}

func FrogJumpBottomUp(arr []int, index int) int {
	if index > len(arr) {
		return math.MaxInt32
	}

	min_cost := math.MaxInt32

	cost :=
		min(
			min_cost,
			getCost(arr[index], arr[index+1])+FrogJumpBottomUp(arr, index+1),
			getCost(arr[index], arr[index+2])+FrogJumpBottomUp(arr, index+2),
		)

	return cost
}

func FrogJumpTopDown(arr []int) int {
	dp := []int{}

	dp[0] = 0
	dp[1] = getCost(arr[1], arr[0])

	for i := 2; i < len(arr); i++ {
		op1 := getCost(arr[i], arr[i-1]) + dp[i-1]
		op2 := getCost(arr[i], arr[i-2]) + dp[i-2]

		dp[i] = min(op1, op2)
	}

	return dp[len(arr)-1]
}
