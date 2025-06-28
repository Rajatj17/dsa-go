package noadjacentsum

func MaxNonAdjancentSum(arr []int) int {
	n := len(arr)
	dp := make([]int, n)

	dp[0] = arr[0]
	dp[1] = max(arr[0], arr[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+arr[i])
	}

	return dp[n-1]
}
