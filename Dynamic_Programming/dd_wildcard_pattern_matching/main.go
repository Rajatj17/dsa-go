package main

/*
Problem: Wildcard Pattern Matching

Description:
Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:

- '?' Matches any single character.
- '*' Matches any sequence of characters (including the empty sequence).

The matching should cover the entire input string (not partial).

Example 1:
Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".

Example 2:
Input: s = "aa", p = "*"
Output: true
Explanation: '*' matches any sequence.

Example 3:
Input: s = "cb", p = "?a"
Output: false
Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.

Example 4:
Input: s = "adceb", p = "*a*b*"
Output: true
Explanation: The first '*' matches the empty sequence, while the second '*' matches the substring "dce".

Example 5:
Input: s = "acdcb", p = "a*c?b"
Output: false

Constraints:
- 0 <= s.length, p.length <= 2000
- s contains only lowercase English letters
- p contains only lowercase English letters, '?' or '*'

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe

Time Complexity: O(m * n) where m is length of string and n is length of pattern
Space Complexity: O(m * n), can be optimized to O(n)

Related Problems:
- Regular Expression Matching
- Edit Distance
- Longest Common Subsequence

Note: This is a classic 2D DP problem.
dp[i][j] represents whether s[0:i] matches p[0:j].

Transitions:
- If p[j-1] == s[i-1] or p[j-1] == '?': dp[i][j] = dp[i-1][j-1]
- If p[j-1] == '*': dp[i][j] = dp[i-1][j] || dp[i][j-1] || dp[i-1][j-1]
  - dp[i-1][j]: '*' matches s[i-1]
  - dp[i][j-1]: '*' matches empty sequence
  - dp[i-1][j-1]: '*' matches s[i-1] and we're done with this '*'
*/

func main() {
	// TODO: Implement the solution
}