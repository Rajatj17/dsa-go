package main

/*
Problem: Hamming Distance

Description:
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, return the Hamming distance between them.

Example 1:
Input: x = 1, y = 4
Output: 2
Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
The above arrows point to positions where the corresponding bits are different.

Example 2:
Input: x = 3, y = 1
Output: 1

Constraints:
- 0 <= x, y <= 2^31 - 1

LeetCode Link: https://leetcode.com/problems/hamming-distance/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(log n) where n is the larger of x and y
Space Complexity: O(1)

Related Problems:
- Total Hamming Distance
- Number of 1 Bits
- Single Number

Note: Multiple approaches:
1. XOR then count bits: x ^ y gives different bits, count 1s
2. Bit by bit comparison: Check each bit position
3. Brian Kernighan's algorithm: Remove rightmost set bit

Key insight: XOR of two numbers gives 1 where bits differ, then count 1s in result.
*/

func main() {
	// TODO: Implement the solution
}