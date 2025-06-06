package shortestgridpath

// Shortest Grid Path
// Given a two dimensional grid,
// each cell of which contains integer cost which represents a cost to traverse through that cell,
// we need to find a path from top left cell to bottom right cell by which total cost incurred is minimum.

// Note : It is assumed that negative cost cycles do not exist in input matrix.

// Input

// Grid as shown above (input given as vector<vector<int> > grid ).
// Hint : Use grid.size() to get rows and grid[0].size() to get columns.

// 31 100 64 12 18
// 10 13 47 157 6
// 100 113 174 11 33
// 88 124 41 20 140
// 99 32 111 41 20
// Output

// An integer denoting the minimum cost.

// 327
// Explanation

// Cells in green are the cells which are visited to complete this route
// 327 (= 31 + 10 + 13 + 47 + 65 + 12 + 18 +
// 6 + 33 + 11 + 20 + 41 + 20)
