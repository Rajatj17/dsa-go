package main

/*
Problem: Split a String Into the Max Number of Unique Substrings

Description:
Given a string s, return the maximum number of unique substrings that the given string can be split into.

You can split string s into any list of non-empty substrings, where the concatenation of the substrings forms the original string. However, you must split the substrings such that all of them are unique.

A substring is a contiguous sequence of characters within a string.

Example 1:
Input: s = "ababccc"
Output: 5
Explanation: One way to split maximally is ['a', 'b', 'ab', 'c', 'cc']. Splitting like ['a', 'b', 'a', 'b', 'c', 'cc'] is not valid as two substrings are the same.

Example 2:
Input: s = "aba"
Output: 2
Explanation: One way to split maximally is ['a', 'ba'].

Example 3:
Input: s = "aa"
Output: 1
Explanation: It is impossible to split the string any further.

Constraints:
- 1 <= s.length <= 16
- s contains only lower case English letters

LeetCode Link: https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/

Companies: Amazon, Microsoft, Google, Apple, Facebook

Time Complexity: O(2^n * n) where n is length of string
Space Complexity: O(n) for recursion stack and set

Related Problems:
- Palindrome Partitioning
- Word Break
- Partition Labels

Note: Similar to palindrome partitioning approach:
1. Use backtracking to try all possible splits
2. Maintain a set of used substrings to ensure uniqueness
3. For each position, try all possible substrings starting from that position
4. Track maximum number of unique substrings found

Key insight: Use backtracking with a set to track used substrings and maximize the count.
*/

func main() {
	// TODO: Implement the solution
}