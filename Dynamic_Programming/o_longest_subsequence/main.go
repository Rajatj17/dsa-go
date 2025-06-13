package longestsubsequence

func LongestIncreasingSubsequence(arr []int) int {
	n := len(arr)
	dp := make([]int, n)

	len := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = max(dp[i], 1+dp[j])
				len = max(dp[i], len)
			}
		}
	}

	return len
}
