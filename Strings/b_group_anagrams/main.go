package main

/*
Problem: Group Anagrams

Description:
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]

Constraints:
- 1 <= strs.length <= 10^4
- 0 <= strs[i].length <= 100
- strs[i] consists of lowercase English letters

LeetCode Link: https://leetcode.com/problems/group-anagrams/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(n * k log k) where n is number of strings and k is max length
Space Complexity: O(n * k)

Related Problems:
- Valid Anagram
- Find All Anagrams in a String
- Anagram Mappings

Note: Hash map approach:
1. For each string, create a key (sorted string or character count)
2. Group strings with same key
3. Use hash map to store groups

Key insight: Anagrams will have the same sorted string or character frequency pattern.
*/

func main() {
	// TODO: Implement the solution
}