package boardgame

// # Word Search in 2D Grid - Original Coding Problem

// ## Problem Statement

// You are given an **m × n** 2D grid (matrix) where each cell contains a single character, and a list of words. Your task is to find all words from the given list that can be formed by connecting adjacent characters in the grid.

// ## Rules:
// 1. **Starting Point**: You can start from any cell in the grid
// 2. **Movement**: From any cell, you can move to any of the 8 adjacent cells (horizontal, vertical, and diagonal neighbors)
// 3. **No Revisiting**: Each cell can only be used once per word formation (no backtracking to previously visited cells)
// 4. **Path Formation**: Words are formed by following a continuous path of adjacent characters

// ## Input:
// - **grid**: A 2D character matrix of size m × n
// - **words**: A list/array of strings to search for

// ## Output:
// - Return a list of all words from the input list that can be found in the grid
// - The order of words in the output doesn't matter

// ## Example:

// **Grid:**
// ```
// [
//   ['S', 'N', 'A', 'K'],
//   ['A', 'T', 'E', 'E'],
//   ['M', 'O', 'C', 'K'],
//   ['R', 'A', 'S', 'H']
// ]
// ```

// **Word List:** `["SNAKE", "FOR", "QUIT", "SNACK", "SNACKS", "TEAM", "CASH"]`

// **Expected Output:** `["SNAKE", "SNACK", "SNACKS", "TEAM", "CASH"]`

// **Explanation:**
// - **SNAKE**: S(0,0) → N(0,1) → A(0,2) → K(0,3) → E(1,3) ✓
// - **FOR**: F not found in grid ✗
// - **QUIT**: Q not found in grid ✗
// - **SNACK**: S(0,0) → N(0,1) → A(0,2) → C(2,2) → K(0,3) ✓
// - **SNACKS**: S(0,0) → N(0,1) → A(0,2) → C(2,2) → K(0,3) → S(3,2) ✓
// - **TEAM**: T(1,1) → E(1,2) → A(1,0) → M(2,0) ✓
// - **CASH**: C(2,2) → A(1,0) → S(3,2) → H(3,3) ✓
