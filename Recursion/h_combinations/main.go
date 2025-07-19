package main

/*
Problem: Combinations

Description:
Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].

You may return the answer in any order.

Example 1:
Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Explanation: There are 4 choose 2 = 6 total combinations.

Example 2:
Input: n = 1, k = 1
Output: [[1]]
Explanation: There is 1 choose 1 = 1 total combination.

Constraints:
- 1 <= n <= 20
- 1 <= k <= n

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe

Time Complexity: O(C(n,k) * k) where C(n,k) is n choose k
Space Complexity: O(k) for recursion stack

Related Problems:
- Permutations
- Combination Sum
- Combination Sum II
- Letter Combinations of a Phone Number

Note: Backtracking approach:
1. For each position, try numbers from current to n
2. Add current number to combination and recurse
3. Remove current number (backtrack) and try next
4. Base case: when combination size equals k

Key insight: Use backtracking with start index to avoid duplicates and ensure lexicographic order.
*/

func main() {
	// TODO: Implement the solution
}