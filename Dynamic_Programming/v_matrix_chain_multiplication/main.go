package main

/*
Problem: Matrix Chain Multiplication

Description:
Given a sequence of matrices, find the most efficient way to multiply these matrices together.
The problem is not actually to perform the multiplications, but merely to decide in which order 
to perform the multiplications.

Given an array p[] which represents the chain of matrices such that the ith matrix Ai is of 
dimension p[i-1] x p[i]. We need to write a function that should return the minimum number 
of multiplications needed to multiply the chain.

Example 1:
Input: p = [1, 2, 3, 4]
Output: 18
Explanation: There are 3 matrices of dimensions 1x2, 2x3, 3x4
The minimum number of multiplications are obtained by putting parenthesis as ((AB)C) or (A(BC))
If we put parenthesis as ((AB)C), then cost is 1*2*3 + 1*3*4 = 6 + 12 = 18
If we put parenthesis as (A(BC)), then cost is 2*3*4 + 1*2*4 = 24 + 8 = 32

Example 2:
Input: p = [40, 20, 30, 10, 30]
Output: 26000
Explanation: There are 4 matrices of dimensions 40x20, 20x30, 30x10, 10x30
Optimal parenthesization is ((40x20)(20x30))((30x10)(10x30))
Cost = 40*20*30 + 30*10*30 + 40*30*30 = 24000 + 9000 + 36000 = 69000

Constraints:
- 1 <= p.length <= 200
- 1 <= p[i] <= 1000

Companies: Amazon, Microsoft, Google, Adobe, Oracle

Time Complexity: O(n^3) where n is the number of matrices
Space Complexity: O(n^2)

Related Problems:
- Burst Balloons
- Minimum Score Triangulation of Polygon
- Boolean Parenthesization
*/

func main() {
	// TODO: Implement the solution
}