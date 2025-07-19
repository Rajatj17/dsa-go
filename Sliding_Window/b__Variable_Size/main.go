package main

import "fmt"

// Problem: Variable Size Sliding Window Template
// This is a template for solving variable-size sliding window problems
// Time Complexity: O(n) where n is the length of the array
// Space Complexity: O(1) for basic problems

// Example 1: Longest substring with at most k distinct characters
func longestSubstringWithKDistinct(s string, k int) int {
	if len(s) == 0 || k == 0 {
		return 0
	}
	
	charCount := make(map[rune]int)
	left := 0
	maxLength := 0
	
	for right := 0; right < len(s); right++ {
		// Expand window
		rightChar := rune(s[right])
		charCount[rightChar]++
		
		// Contract window until we have at most k distinct characters
		for len(charCount) > k {
			leftChar := rune(s[left])
			charCount[leftChar]--
			if charCount[leftChar] == 0 {
				delete(charCount, leftChar)
			}
			left++
		}
		
		// Update maximum length
		maxLength = max(maxLength, right-left+1)
	}
	
	return maxLength
}

// Example 2: Minimum window substring
func minWindowSubstring(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	
	// Count characters in t
	tCount := make(map[rune]int)
	for _, char := range t {
		tCount[char]++
	}
	
	required := len(tCount)
	formed := 0
	
	windowCounts := make(map[rune]int)
	left := 0
	minLen := len(s) + 1
	minLeft := 0
	
	for right := 0; right < len(s); right++ {
		// Expand window
		rightChar := rune(s[right])
		windowCounts[rightChar]++
		
		if count, exists := tCount[rightChar]; exists && windowCounts[rightChar] == count {
			formed++
		}
		
		// Contract window
		for left <= right && formed == required {
			// Update minimum window
			if right-left+1 < minLen {
				minLen = right - left + 1
				minLeft = left
			}
			
			leftChar := rune(s[left])
			windowCounts[leftChar]--
			if count, exists := tCount[leftChar]; exists && windowCounts[leftChar] < count {
				formed--
			}
			left++
		}
	}
	
	if minLen == len(s)+1 {
		return ""
	}
	return s[minLeft : minLeft+minLen]
}

// Example 3: Longest subarray with sum at most k
func longestSubarrayWithSumAtMostK(arr []int, k int) int {
	left := 0
	sum := 0
	maxLength := 0
	
	for right := 0; right < len(arr); right++ {
		// Expand window
		sum += arr[right]
		
		// Contract window
		for sum > k {
			sum -= arr[left]
			left++
		}
		
		// Update maximum length
		maxLength = max(maxLength, right-left+1)
	}
	
	return maxLength
}

// Example 4: Smallest subarray with sum greater than k
func smallestSubarrayWithSumGreaterThanK(arr []int, k int) int {
	left := 0
	sum := 0
	minLength := len(arr) + 1
	
	for right := 0; right < len(arr); right++ {
		// Expand window
		sum += arr[right]
		
		// Contract window
		for sum > k {
			minLength = min(minLength, right-left+1)
			sum -= arr[left]
			left++
		}
	}
	
	if minLength == len(arr)+1 {
		return 0 // No valid subarray found
	}
	return minLength
}

// Example 5: Fruits into baskets (at most 2 types)
func fruitsIntoBaskets(fruits []int) int {
	fruitCount := make(map[int]int)
	left := 0
	maxFruits := 0
	
	for right := 0; right < len(fruits); right++ {
		// Expand window
		fruitCount[fruits[right]]++
		
		// Contract window to have at most 2 types
		for len(fruitCount) > 2 {
			fruitCount[fruits[left]]--
			if fruitCount[fruits[left]] == 0 {
				delete(fruitCount, fruits[left])
			}
			left++
		}
		
		// Update maximum fruits
		maxFruits = max(maxFruits, right-left+1)
	}
	
	return maxFruits
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("=== Variable Size Sliding Window Examples ===")
	
	// Example 1: Longest substring with k distinct characters
	fmt.Println("\n1. Longest substring with at most K distinct characters:")
	s1 := "aabacbebebe"
	k1 := 3
	fmt.Printf("String: %s, K: %d\n", s1, k1)
	fmt.Printf("Result: %d\n", longestSubstringWithKDistinct(s1, k1))
	
	// Example 2: Minimum window substring
	fmt.Println("\n2. Minimum window substring:")
	s2 := "ADOBECODEBANC"
	t2 := "ABC"
	fmt.Printf("String: %s, Target: %s\n", s2, t2)
	fmt.Printf("Result: %s\n", minWindowSubstring(s2, t2))
	
	// Example 3: Longest subarray with sum at most k
	fmt.Println("\n3. Longest subarray with sum at most K:")
	arr3 := []int{1, 4, 2, 9, 5, 10, 7}
	k3 := 10
	fmt.Printf("Array: %v, K: %d\n", arr3, k3)
	fmt.Printf("Result: %d\n", longestSubarrayWithSumAtMostK(arr3, k3))
	
	// Example 4: Smallest subarray with sum greater than k
	fmt.Println("\n4. Smallest subarray with sum greater than K:")
	arr4 := []int{2, 1, 2, 4, 3, 1}
	k4 := 7
	fmt.Printf("Array: %v, K: %d\n", arr4, k4)
	fmt.Printf("Result: %d\n", smallestSubarrayWithSumGreaterThanK(arr4, k4))
	
	// Example 5: Fruits into baskets
	fmt.Println("\n5. Fruits into baskets (at most 2 types):")
	fruits := []int{1, 2, 1, 2, 3, 1, 2}
	fmt.Printf("Fruits: %v\n", fruits)
	fmt.Printf("Maximum fruits: %d\n", fruitsIntoBaskets(fruits))
}