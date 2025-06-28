package main

import "math"

// Give an array of coin denominations & integer M representing target money.
// We need to find the min coins required to make the change.
// Input
// coins = [1,3,7,10]
// M = 15
// Output
// 3 (7 + 7 + 1)

func CoinChange(coin []int) {
	// Initialize matrix
	// x-axis: sum
	// y-axis: length of array
}

func CoinChangeTD(coins []int, sum int) int {
	if sum < 0 {
		return -1
	}
	if sum == 0 {
		return 0
	}

	minCoins := math.MaxInt32

	for _, coin := range coins {
		result := CoinChangeTD(coins, sum-coin)
		if result != -1 {
			minCoins = min(minCoins, result+1)
		}
	}

	if minCoins == math.MaxInt32 {
		return -1
	}

	return minCoins
}

func CoinChangeBU(coins []int, amount int) int {
	dp := make([]int, amount+1)

	// Initialize with max value (infinity)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0 // 0 coins needed to make amount 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// coins = [10, 20, 40]
	// M = 90
	// dp[0] = 0
	// dp[1] = -1
	// dp[2] = -1
	// dp[10] = 1
	// dp[20] = 1 + dp[10], 1 + dp[0]
	// dp[30] = 1 + dp[20], 1 + dp[10]

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
