package main

/*
Problem: Minimum Number of Days to Disconnect Island

Description:
You are given an m x n binary grid grid where 1 represents land and 0 represents water. An island is a maximal 4-directionally (horizontal or vertical) connected group of 1's.

The grid is said to be connected if we have exactly one island, otherwise is said disconnected.

In one day, we are allowed to change any single land cell (1) into a water cell (0).

Return the minimum number of days to disconnect the grid.

Example 1:
Input: grid = [[0,1,1,0],[0,1,1,0],[0,0,0,0]]
Output: 2
Explanation: We need at least 2 days to get a disconnected grid.
Change land grid[1][1] and grid[0][2] to water and get 2 separated islands.

Example 2:
Input: grid = [[1,1]]
Output: 2
Explanation: Grid of full water is also disconnected ([[0,0]]), so you can return 2.

Constraints:
- m == grid.length
- n == grid[i].length
- 1 <= m, n <= 30
- grid[i][j] is either 0 or 1

LeetCode Link: https://leetcode.com/problems/minimum-number-of-days-to-disconnect-island/

Companies: Amazon, Microsoft, Google, Apple, Facebook

Time Complexity: O(m^2 * n^2)
Space Complexity: O(m * n)

Related Problems:
- Number of Islands
- Critical Connections in a Network
- Articulation Points and Bridges

Note: Graph connectivity analysis:
1. Check if grid is already disconnected (0 days)
2. Try removing each land cell and check if disconnected (1 day)
3. Otherwise, answer is 2 (can always disconnect by removing 2 cells)

Key insight: The answer is at most 2 - we can always disconnect by removing any corner land cell and its neighbor.
*/

func main() {
	// TODO: Implement the solution
}