package main

/*
Problem: Container With Most Water

Description:
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container that can hold the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Example 1:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:
Input: height = [1,1]
Output: 1

Constraints:
- n == height.length
- 2 <= n <= 10^5
- 0 <= height[i] <= 10^4

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n) using two-pointer approach
Space Complexity: O(1)

Related Problems:
- Trapping Rain Water
- Largest Rectangle in Histogram
- Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts

Note: Classic two-pointer problem. Start from both ends and move the pointer with smaller height.
Key insight: Moving the shorter line might give us a larger area, but moving the taller line will definitely not.
*/

func main() {
	// TODO: Implement the solution
}