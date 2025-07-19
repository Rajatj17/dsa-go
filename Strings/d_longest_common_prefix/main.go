package main

/*
Problem: Longest Common Prefix

Description:
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Constraints:
- 1 <= strs.length <= 200
- 0 <= strs[i].length <= 200
- strs[i] consists of only lowercase English letters

LeetCode Link: https://leetcode.com/problems/longest-common-prefix/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(S) where S is sum of all characters in all strings
Space Complexity: O(1)

Related Problems:
- Longest Common Subsequence
- Find the Index of the First Occurrence in a String

Note: Multiple approaches:
1. Vertical scanning: Compare characters at each position
2. Horizontal scanning: Compare strings pair by pair
3. Divide and conquer: Recursively find prefix
4. Trie-based approach

Key insight: The longest common prefix cannot be longer than the shortest string.
*/

func main() {
	// TODO: Implement the solution
}