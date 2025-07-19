package main

/*
Problem: Combination Sum

Description:
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

Example 1:
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.

Example 2:
Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

Example 3:
Input: candidates = [2], target = 1
Output: []

Constraints:
- 1 <= candidates.length <= 30
- 2 <= candidates[i] <= 40
- All elements of candidates are distinct
- 1 <= target <= 40

LeetCode Link: https://leetcode.com/problems/combination-sum/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(N^(T/M)) where N is number of candidates, T is target, M is minimal candidate value
Space Complexity: O(T/M) for recursion stack

Related Problems:
- Combination Sum II
- Combination Sum III
- Combination Sum IV
- Subsets

Note: Combination saving approach with backtracking:
1. Sort candidates array for optimization
2. Use backtracking with start index to avoid duplicates
3. Can reuse same element multiple times
4. Prune when current sum exceeds target

Key insight: Use start index to avoid generating duplicate combinations in different orders.
*/

func main() {
	// TODO: Implement the solution
}