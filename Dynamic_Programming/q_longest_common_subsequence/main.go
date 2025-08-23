package longestcommonsubsequence

/*
Problem: Longest Common Subsequence (LCS)

Description:
Given two strings text1 and text2, return the length of their longest common subsequence.
If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters 
(can be none) deleted without changing the relative order of the remaining characters.

Example 1:
Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.

Example 2:
Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.

Example 3:
Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.

Constraints:
- 1 <= text1.length, text2.length <= 1000
- text1 and text2 consist of only lowercase English characters

DP Recurrence:
- If characters match: dp[i][j] = 1 + dp[i-1][j-1]
- If characters don't match: dp[i][j] = max(dp[i-1][j], dp[i][j-1])

Time Complexity: O(m * n) where m and n are lengths of strings
Space Complexity: O(m * n)

Companies: Google, Amazon, Microsoft, Facebook, Bloomberg
*/

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
