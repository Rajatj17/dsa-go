package main

var dp = [][]int{{}}

/*
*

	We will mostly be given the following:
	Price Array: [2, 3, 5, 7, 4]
	Weight Array: [2. 4, 5, 6, 7]

*
*/
func knapsackProblem(price []int, weight []int, n int, maxWeight int) int {
	if n == 0 || maxWeight == 0 {
		return 0
	}

	if dp[n][maxWeight] != -1 {
		return dp[n][maxWeight]
	}

	if weight[n-1] > maxWeight {
		dp[n][maxWeight] = knapsackProblem(price, weight, n-1, maxWeight)

		return dp[n][maxWeight]
	}

	dp[n][maxWeight] = max(
		price[n-1]+knapsackProblem(price, weight, n-1, maxWeight-weight[n-1]),
		knapsackProblem(price, weight, n-1, maxWeight),
	)

	return dp[n][maxWeight]
}

func main() {
	var price []int
	var weight []int
	var maxWeight int

	knapsackProblem(price, weight, len(price), maxWeight)
}
