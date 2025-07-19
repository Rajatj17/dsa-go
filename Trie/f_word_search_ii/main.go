package main

/*
Problem: Word Search II

Description:
Given an m x n board of characters and a list of strings words, return all words on the board.

Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

Example 1:
Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]

Example 2:
Input: board = [["a","b"],["c","d"]], words = ["abcb"]
Output: []

Constraints:
- m == board.length
- n == board[i].length
- 1 <= m, n <= 12
- board[i][j] is a lowercase English letter
- 1 <= words.length <= 3 * 10^4
- 1 <= words[i].length <= 10
- words[i] consists of lowercase English letters
- All the strings of words are unique

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(M * N * 4^L) where M*N is board size and L is max word length
Space Complexity: O(K * L) where K is number of words

Related Problems:
- Word Search
- Boggle Game
- Stream of Characters

Note: Trie + DFS backtracking:
1. Build a trie from all words
2. For each cell in board, start DFS if it matches trie root
3. During DFS, traverse trie and board simultaneously
4. When we reach end of a word in trie, add to result
5. Use backtracking to explore all paths

Key insight: Trie allows us to prune search space efficiently when current path doesn't match any word prefix.
*/

func main() {
	// TODO: Implement the solution
}