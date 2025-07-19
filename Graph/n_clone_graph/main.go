package main

/*
Problem: Clone Graph

Description:
Given a reference of a node in a connected undirected graph.

Return a deep copy (clone) of the graph.

Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

Example 1:
Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
Output: [[2,4],[1,3],[2,4],[1,3]]
Explanation: There are 4 nodes in the graph.
1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).

Example 2:
Input: adjList = [[]]
Output: [[]]

Example 3:
Input: adjList = []
Output: []

Constraints:
- The number of nodes in the graph is in the range [0, 100]
- 1 <= Node.val <= 100
- Node.val is unique for each node
- There are no repeated edges and no self-loops in the graph
- The Graph is connected and all nodes can be visited starting from the given node

LeetCode Link: https://leetcode.com/problems/clone-graph/

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(V + E) where V is vertices and E is edges
Space Complexity: O(V) for the hash map

Related Problems:
- Copy List with Random Pointer
- Clone Binary Tree With Random Pointer
- Clone N-ary Tree

Note: BFS and DFS approaches:
1. Use hash map to store mapping from original to cloned nodes
2. DFS: Recursively clone neighbors
3. BFS: Use queue to clone level by level

Key insight: Use hash map to avoid infinite loops and ensure each node is cloned only once.
*/

func main() {
	// TODO: Implement the solution
}