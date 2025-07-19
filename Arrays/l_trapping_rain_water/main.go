package main

/*
Problem: Trapping Rain Water

Description:
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

Example 1:
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

Example 2:
Input: height = [4,2,0,3,2,5]
Output: 9

Constraints:
- n == height.length
- 1 <= n <= 2 * 10^4
- 0 <= height[i] <= 3 * 10^4

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n) using two-pointer approach
Space Complexity: O(1)

Related Problems:
- Container With Most Water
- Product of Array Except Self
- Candy

Note: Can be solved using multiple approaches:
1. Brute Force: O(n^2) - For each element, find max height to left and right
2. Dynamic Programming: O(n) time, O(n) space - Precompute max heights
3. Two Pointers: O(n) time, O(1) space - Optimal solution
4. Stack: O(n) time, O(n) space - Alternative approach

Key insight for two pointers: Water level at any point is determined by the minimum of max heights on both sides.
*/

func main() {
	// TODO: Implement the solution
}