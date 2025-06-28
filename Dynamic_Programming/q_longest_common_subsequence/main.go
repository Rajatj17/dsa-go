package longestcommonsubsequence

func LCSBottomUp(s1 string, s2 string) {
	n1 := len(s1)
	n2 := len(s2)

	dp := [][]int{{}}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Print the match letters

	// result := ""

	// for i, j := n1, n2; i > 0 && j > 0; {
	// 	if d[i][j] = dp[j][i]
	// }
}
