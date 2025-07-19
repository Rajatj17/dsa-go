package main

/*
Problem: Subsets

Description:
Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

Example 1:
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:
Input: nums = [0]
Output: [[],[0]]

Constraints:
- 1 <= nums.length <= 10
- -10 <= nums[i] <= 10
- All the numbers of nums are unique

LeetCode Link: https://leetcode.com/problems/subsets/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(2^n * n) where n is the length of nums
Space Complexity: O(2^n * n) for storing all subsets

Related Problems:
- Subsets II
- Permutations
- Combination Sum
- Generate Parentheses

Note: Classic backtracking problem with multiple approaches:
1. Backtracking with include/exclude decisions
2. Bit manipulation (each bit represents include/exclude)
3. Iterative approach building subsets incrementally

Key insight: For each element, we have two choices - include it or exclude it.
*/

func main() {
	// TODO: Implement the solution
}