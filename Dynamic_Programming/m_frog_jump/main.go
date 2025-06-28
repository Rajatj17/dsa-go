package frogjump

import "math"

// There are N stones given in the form an array,
// each element in array representing the height of the ith stone.
// This is a frog who is initially on Stone 1.
// Frog will repeat the following action some number of times to each the
// stone N
// If the frog is on Stone I, jump to Stone i + 1 or Stone i + 2
// Find the min possible cost

// Input
// 2
// [10, 10]
// O/P: 0

// 6
// [30, 10, 60, 10, 60, 50]
// O/P: 40 1->3->5->6

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
