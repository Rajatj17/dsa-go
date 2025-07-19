package main

/*
Problem: Set Matrix Zeroes

Description:
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.

Example 1:
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Example 2:
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

Constraints:
- m == matrix.length
- n == matrix[0].length
- 1 <= m, n <= 200
- -2^31 <= matrix[i][j] <= 2^31 - 1

Follow up:
- A straightforward solution using O(mn) space is probably a bad idea.
- A simple improvement uses O(m + n) space, but still not the best solution.
- Could you devise a constant space solution?

LeetCode Link: https://leetcode.com/problems/set-matrix-zeroes/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(m * n)
Space Complexity: O(1)

Related Problems:
- Game of Life
- Spiral Matrix
- Rotate Image

Note: Multiple approaches:
1. O(m + n) space: Use arrays to track rows/cols to zero
2. O(1) space: Use first row/col as markers
3. Handle edge cases for first row/col separately

Key insight: Use the matrix itself to store which rows/columns need to be zeroed.
*/

func main() {
	// TODO: Implement the solution
}