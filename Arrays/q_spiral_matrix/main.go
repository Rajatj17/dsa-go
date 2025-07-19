package main

/*
Problem: Spiral Matrix

Description:
Given an m x n matrix, return all elements of the matrix in spiral order.

Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Example 2:
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

Constraints:
- m == matrix.length
- n == matrix[i].length
- 1 <= m, n <= 10
- -100 <= matrix[i][j] <= 100

LeetCode Link: https://leetcode.com/problems/spiral-matrix/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(m * n)
Space Complexity: O(1) not counting output array

Related Problems:
- Spiral Matrix II
- Spiral Matrix III
- Spiral Matrix IV

Note: Layer by layer traversal:
1. Define boundaries: top, bottom, left, right
2. Traverse right, then down, then left, then up
3. Update boundaries after each direction
4. Continue until all elements are visited

Key insight: Use four boundaries and traverse in spiral order, updating boundaries after each direction.
*/

func main() {
	// TODO: Implement the solution
}