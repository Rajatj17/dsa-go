package main

/*
Problem: Longest Substring Without Repeating Characters

Description:
Given a string s, find the length of the longest substring without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:
- 0 <= s.length <= 5 * 10^4
- s consists of English letters, digits, symbols and spaces

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n)
Space Complexity: O(min(m, n)) where m is the size of the charset

Related Problems:
- Longest Substring with At Most Two Distinct Characters
- Longest Substring with At Most K Distinct Characters
- Subarrays with K Different Integers

Note: Sliding window with hashmap:
1. Use two pointers (left and right) to maintain a window
2. Expand window by moving right pointer
3. When duplicate found, shrink window from left until no duplicates
4. Keep track of maximum window size

Key insight: Use a hashmap to store the most recent index of each character.
*/

func main() {
	// TODO: Implement the solution
}