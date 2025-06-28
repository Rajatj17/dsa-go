package main

import "math"

/**
Rod Cutting Problem
 Given a rod of length n inches and an array of prices that contains prices of all pieces of size smaller than n.
 Determine the  maximum value obtainable by cutting up the rod and selling the pieces.
Example:
 If length of the rod is 8 and the values of different pieces are given as following,
 then the maximum obtainable value is 22 (by cutting in two pieces of lengths 2 and 6)


length   | 1   2   3   4   5   6   7   8
--------------------------------------------
price    | 1   5   8   9  10  17  17  20

**/

func RodCutting(price []int, length []int, n int, w int) int {
	t := [][]int{{}}

	for i := range n + 1 {
		for j := range w + 1 {

			if length[i-1] <= j {

				t[i][j] = max(
					price[i-1]+t[i][j-length[i-1]],
					t[i-1][j],
				)
			} else {

				t[i][j] = t[i-1][j]
			}

		}
	}

	return t[n][w]
}

func RodCuttingRecursive(prices []int, n int, dp []int) int {
	if dp[n] != -1 {
		return dp[n]
	}

	if n <= 0 {
		return 0
	}

	ans := math.MinInt32
	for i := range n {
		cut := i + 1

		current_ans := prices[i] + RodCuttingRecursive(prices, n-cut, dp)

		ans = max(current_ans, ans)
	}

	dp[n] = ans
	return ans
}

func RodCuttingBottomUp(prices []int, n int) {
	dp := make([]int, n)

	dp[0] = 0

	for length := range n {

		ans := math.MinInt32
		for i := range length {
			cut := i + 1

			current_ans := prices[i] + dp[n-cut]
			ans = max(current_ans, ans)
		}
	}
}

func main() {
	var price []int
	var length []int

	RodCutting(price, length, len(price), len(length))
}
