package main

/*
Problem: Course Schedule

Description:
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

Return true if you can finish all courses. Otherwise, return false.

Example 1:
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take. 
To take course 1 you should have finished course 0. So it is possible.

Example 2:
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take. 
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.

Constraints:
- 1 <= numCourses <= 2000
- 0 <= prerequisites.length <= 5000
- prerequisites[i].length == 2
- 0 <= ai, bi < numCourses
- All the pairs prerequisites[i] are unique

Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Adobe, Uber

Time Complexity: O(V + E) where V is number of courses and E is number of prerequisites
Space Complexity: O(V + E)

Related Problems:
- Course Schedule II
- Graph Valid Tree
- Minimum Height Trees
- Parallel Courses

Note: This is a cycle detection problem in a directed graph.
Two approaches:
1. DFS with recursion stack to detect back edges
2. Topological Sort using Kahn's algorithm (BFS with indegree)

Key insight: If there's a cycle in the dependency graph, it's impossible to complete all courses.
*/

func main() {
	// TODO: Implement the solution
}