package main

/*
Problem: Number of Islands

Description:
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

Example 2:
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

Constraints:
- m == grid.length
- n == grid[i].length
- 1 <= m, n <= 300
- grid[i][j] is '0' or '1'

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(m * n)
Space Complexity: O(m * n) in worst case for recursion stack

Related Problems:
- Max Area of Island
- Number of Distinct Islands
- Number of Closed Islands
- Surrounded Regions

Note: Classic DFS/BFS problem:
1. Iterate through each cell in the grid
2. If cell is '1', start DFS/BFS and increment island count
3. During DFS/BFS, mark all connected '1's as visited
4. Continue until all cells are processed

Key insight: Each DFS/BFS call explores one complete island.
*/

func main() {
	// TODO: Implement the solution
}