package main

/*
Problem: Word Break

Description:
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into 
a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.

Example 1:
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.

Example 3:
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false
Explanation: Cannot be segmented into dictionary words.

Constraints:
- 1 <= s.length <= 300
- 1 <= wordDict.length <= 1000
- 1 <= wordDict[i].length <= 20
- s and wordDict[i] consist of only lowercase English letters
- All the strings of wordDict are unique

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n^2 * m) where n is length of string and m is average length of words
Space Complexity: O(n)

Related Problems:
- Word Break II
- Concatenated Words
- Extra Characters in a String
- Sentence Screen Fitting

Note: This is a classic 1D DP problem where we check if we can break the string at each position.
dp[i] represents whether the substring s[0:i] can be segmented using the dictionary.
*/

func main() {
	// TODO: Implement the solution
}