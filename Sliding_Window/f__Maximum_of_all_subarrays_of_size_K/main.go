package main

import "fmt"

// Problem: Maximum of all subarrays of size K
// Given an array and a window size K, find the maximum element in each window
// Time Complexity: O(n) using deque approach
// Space Complexity: O(k) for storing indices in deque

// Approach 1: Brute Force - O(n*k)
func maxInWindowBruteForce(arr []int, k int) []int {
	if len(arr) < k {
		return []int{}
	}
	
	var result []int
	
	// Check each window
	for i := 0; i <= len(arr)-k; i++ {
		maxVal := arr[i]
		for j := i; j < i+k; j++ {
			maxVal = max(maxVal, arr[j])
		}
		result = append(result, maxVal)
	}
	
	return result
}

// Approach 2: Using Deque (Double-ended queue) - O(n)
func maxInWindowOptimized(arr []int, k int) []int {
	if len(arr) < k {
		return []int{}
	}
	
	var result []int
	var deque []int // Store indices
	
	// Process first window
	for i := 0; i < k; i++ {
		// Remove elements smaller than current element from back
		for len(deque) > 0 && arr[deque[len(deque)-1]] <= arr[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	
	// The front of deque contains the index of maximum element
	result = append(result, arr[deque[0]])
	
	// Process remaining elements
	for i := k; i < len(arr); i++ {
		// Remove elements outside current window from front
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		
		// Remove elements smaller than current element from back
		for len(deque) > 0 && arr[deque[len(deque)-1]] <= arr[i] {
			deque = deque[:len(deque)-1]
		}
		
		deque = append(deque, i)
		result = append(result, arr[deque[0]])
	}
	
	return result
}

// Approach 3: Using Monotonic Stack approach
func maxInWindowStack(arr []int, k int) []int {
	if len(arr) < k {
		return []int{}
	}
	
	var result []int
	
	// Stack to maintain indices in decreasing order of values
	var stack []int
	
	for i := 0; i < len(arr); i++ {
		// Remove elements outside current window
		for len(stack) > 0 && stack[0] <= i-k {
			stack = stack[1:]
		}
		
		// Maintain decreasing order
		for len(stack) > 0 && arr[stack[len(stack)-1]] <= arr[i] {
			stack = stack[:len(stack)-1]
		}
		
		stack = append(stack, i)
		
		// Add maximum to result if we have processed at least k elements
		if i >= k-1 {
			result = append(result, arr[stack[0]])
		}
	}
	
	return result
}

// Approach 4: Using segment tree (for educational purposes)
type SegmentTree struct {
	tree []int
	n    int
}

func buildSegmentTree(arr []int) *SegmentTree {
	n := len(arr)
	tree := make([]int, 4*n)
	st := &SegmentTree{tree: tree, n: n}
	st.build(arr, 0, 0, n-1)
	return st
}

func (st *SegmentTree) build(arr []int, node, start, end int) {
	if start == end {
		st.tree[node] = arr[start]
	} else {
		mid := (start + end) / 2
		st.build(arr, 2*node+1, start, mid)
		st.build(arr, 2*node+2, mid+1, end)
		st.tree[node] = max(st.tree[2*node+1], st.tree[2*node+2])
	}
}

func (st *SegmentTree) query(node, start, end, l, r int) int {
	if r < start || end < l {
		return -1000000 // Minimum value
	}
	if l <= start && end <= r {
		return st.tree[node]
	}
	mid := (start + end) / 2
	left := st.query(2*node+1, start, mid, l, r)
	right := st.query(2*node+2, mid+1, end, l, r)
	return max(left, right)
}

func maxInWindowSegmentTree(arr []int, k int) []int {
	if len(arr) < k {
		return []int{}
	}
	
	st := buildSegmentTree(arr)
	var result []int
	
	for i := 0; i <= len(arr)-k; i++ {
		maxVal := st.query(0, 0, len(arr)-1, i, i+k-1)
		result = append(result, maxVal)
	}
	
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function to print array with indices
func printArrayWithIndices(arr []int) {
	fmt.Print("Indices: ")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%2d ", i)
	}
	fmt.Println()
	
	fmt.Print("Values:  ")
	for _, val := range arr {
		fmt.Printf("%2d ", val)
	}
	fmt.Println()
}

func main() {
	fmt.Println("=== Maximum of all subarrays of size K ===")
	
	// Test case 1
	arr1 := []int{1, 2, 3, 1, 4, 5, 2, 3, 6}
	k1 := 3
	
	fmt.Printf("\nTest Case 1:\n")
	fmt.Printf("Array: %v\n", arr1)
	fmt.Printf("Window size: %d\n", k1)
	printArrayWithIndices(arr1)
	
	fmt.Printf("\nApproach 1 (Brute Force): %v\n", maxInWindowBruteForce(arr1, k1))
	fmt.Printf("Approach 2 (Deque): %v\n", maxInWindowOptimized(arr1, k1))
	fmt.Printf("Approach 3 (Stack): %v\n", maxInWindowStack(arr1, k1))
	fmt.Printf("Approach 4 (Segment Tree): %v\n", maxInWindowSegmentTree(arr1, k1))
	
	// Test case 2
	arr2 := []int{8, 5, 10, 7, 9, 4, 15, 12, 90, 13}
	k2 := 4
	
	fmt.Printf("\nTest Case 2:\n")
	fmt.Printf("Array: %v\n", arr2)
	fmt.Printf("Window size: %d\n", k2)
	printArrayWithIndices(arr2)
	
	fmt.Printf("\nApproach 1 (Brute Force): %v\n", maxInWindowBruteForce(arr2, k2))
	fmt.Printf("Approach 2 (Deque): %v\n", maxInWindowOptimized(arr2, k2))
	fmt.Printf("Approach 3 (Stack): %v\n", maxInWindowStack(arr2, k2))
	fmt.Printf("Approach 4 (Segment Tree): %v\n", maxInWindowSegmentTree(arr2, k2))
	
	// Test case 3: Edge cases
	fmt.Printf("\nEdge Cases:\n")
	
	// Single element
	single := []int{5}
	fmt.Printf("Single element [5], k=1: %v\n", maxInWindowOptimized(single, 1))
	
	// K equals array length
	equal := []int{1, 3, 2, 4}
	fmt.Printf("K equals array length [1,3,2,4], k=4: %v\n", maxInWindowOptimized(equal, 4))
	
	// Decreasing array
	decreasing := []int{5, 4, 3, 2, 1}
	fmt.Printf("Decreasing array [5,4,3,2,1], k=3: %v\n", maxInWindowOptimized(decreasing, 3))
	
	// Increasing array
	increasing := []int{1, 2, 3, 4, 5}
	fmt.Printf("Increasing array [1,2,3,4,5], k=3: %v\n", maxInWindowOptimized(increasing, 3))
	
	// All same elements
	same := []int{3, 3, 3, 3, 3}
	fmt.Printf("All same elements [3,3,3,3,3], k=3: %v\n", maxInWindowOptimized(same, 3))
}