package main

/*
Problem: Word Search

Description:
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example 1:
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true

Example 2:
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true

Example 3:
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false

Constraints:
- m == board.length
- n = board[i].length
- 1 <= m, n <= 6
- 1 <= word.length <= 15
- board and word consists of only lowercase and uppercase English letters

LeetCode Link: https://leetcode.com/problems/word-search/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(N * 4^L) where N is number of cells and L is length of word
Space Complexity: O(L) for recursion stack

Related Problems:
- Word Search II
- Boggle Game
- Exist Word in Board

Note: Backtracking in matrix with index optimization:
1. For each cell, try to start word search from that position
2. Use DFS with backtracking to explore all 4 directions
3. Mark visited cells temporarily and unmark when backtracking
4. Prune early if current path can't lead to solution

Key insight: Use backtracking with visited marking to avoid using same cell twice in one path.
*/

func main() {
	// TODO: Implement the solution
}