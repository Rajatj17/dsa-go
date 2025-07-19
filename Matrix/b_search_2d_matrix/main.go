package main

/*
Problem: Search a 2D Matrix

Description:
Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:
- Integers in each row are sorted from left to right.
- The first integer of each row is greater than the last integer of the previous row.

Example 1:
Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
Output: true

Example 2:
Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 14
Output: false

Constraints:
- m == matrix.length
- n == matrix[i].length
- 1 <= m, n <= 100
- -10^4 <= matrix[i][j], target <= 10^4

LeetCode Link: https://leetcode.com/problems/search-a-2d-matrix/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(log(m * n))
Space Complexity: O(1)

Related Problems:
- Search a 2D Matrix II
- Find Peak Element II

Note: Multiple approaches:
1. Treat as 1D array and use binary search
2. Two-step binary search: find row, then find column
3. Start from top-right corner and eliminate row/col

Key insight: Since matrix is sorted like a 1D array, we can use binary search with coordinate conversion.
*/

func main() {
	// TODO: Implement the solution
}