package main

/*
Problem: Rotate Image

Description:
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]

Example 2:
Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

Constraints:
- n == matrix.length == matrix[i].length
- 1 <= n <= 20
- -1000 <= matrix[i][j] <= 1000

LeetCode Link: https://leetcode.com/problems/rotate-image/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n^2)
Space Complexity: O(1)

Related Problems:
- Rotate Array
- Spiral Matrix
- Transpose Matrix

Note: Two approaches:
1. Transpose then reverse each row
2. Rotate groups of four cells directly

Key insight: 90-degree clockwise rotation = transpose + reverse each row
- Transpose: matrix[i][j] = matrix[j][i]
- Reverse rows: reverse each row individually
*/

func main() {
	// TODO: Implement the solution
}