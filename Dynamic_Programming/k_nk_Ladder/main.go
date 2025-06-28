package nkladder

func CuntWaysTD(n int, k int, dp []int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if dp[n] != 0 {
		return dp[n]
	}

	ans := 0
	for jump := range k {
		ans += CuntWaysTD(n-jump, k, dp)
	}

	dp[n] = ans
	return ans
}

// O(n*k)
func CountWaysBU(n int, k int) int {
	dp := []int{}

	dp[0] = 1

	for i := 1; i <= n; i++ {
		for jump := 1; jump <= k; jump++ {
			if i-jump >= 0 {
				dp[i] += dp[i-jump]
			}
		}
	}

	return dp[n]
}
