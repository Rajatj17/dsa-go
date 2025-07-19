package main

/*
Problem: Unique Paths

Description:
There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]).
The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]).
The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take 
to reach the bottom-right corner.

Example 1:
Input: m = 3, n = 7
Output: 28
Explanation: There are 28 unique paths from top-left to bottom-right.

Example 2:
Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down

Example 3:
Input: m = 1, n = 1
Output: 1
Explanation: There is only one way to reach the destination.

Constraints:
- 1 <= m, n <= 100

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(m * n)
Space Complexity: O(m * n), can be optimized to O(min(m, n))

Related Problems:
- Unique Paths II (with obstacles)
- Minimum Path Sum
- Paths with Maximum Score
- Number of Paths with Max Score

Note: This is a classic 2D DP problem. 
dp[i][j] = dp[i-1][j] + dp[i][j-1] represents the number of ways to reach cell (i,j).
*/

func main() {
	// TODO: Implement the solution
}