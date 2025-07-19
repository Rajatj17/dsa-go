package main

/*
Problem: Minimum Window Substring

Description:
Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character in t (including duplicates) is included in the window. If there is no such window, return the empty string "".

The testcases will be generated such that the answer is unique.

Example 1:
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

Example 2:
Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.

Example 3:
Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.

Constraints:
- m == s.length
- n == t.length
- 1 <= m, n <= 10^5
- s and t consist of uppercase and lowercase English letters

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(|s| + |t|)
Space Complexity: O(|s| + |t|)

Related Problems:
- Substring with Concatenation of All Words
- Minimum Size Subarray Sum
- Sliding Window Maximum

Note: Sliding window with two pointers:
1. Expand right pointer until window contains all characters from t
2. Once valid window found, try to shrink from left while maintaining validity
3. Keep track of minimum valid window
4. Use hashmaps to track character frequencies

Key insight: Use two hashmaps - one for target characters and one for current window characters.
*/

func main() {
	// TODO: Implement the solution
}