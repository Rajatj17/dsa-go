package main

import "fmt"

// Problem: Longest Substring with K Unique Characters
// Given a string, find the length of the longest substring with exactly K unique characters
// Time Complexity: O(n) where n is length of string
// Space Complexity: O(k) for storing character counts

// Approach 1: Sliding Window with HashMap
func longestSubstringWithKUniqueChars(s string, k int) int {
	if len(s) == 0 || k == 0 {
		return 0
	}
	
	charCount := make(map[byte]int)
	left := 0
	maxLength := 0
	
	for right := 0; right < len(s); right++ {
		// Expand window
		charCount[s[right]]++
		
		// Contract window if more than k unique characters
		for len(charCount) > k {
			charCount[s[left]]--
			if charCount[s[left]] == 0 {
				delete(charCount, s[left])
			}
			left++
		}
		
		// Update max length only if we have exactly k unique characters
		if len(charCount) == k {
			maxLength = max(maxLength, right-left+1)
		}
	}
	
	return maxLength
}

// Approach 2: Return the actual substring
func longestSubstringWithKUniqueCharsString(s string, k int) string {
	if len(s) == 0 || k == 0 {
		return ""
	}
	
	charCount := make(map[byte]int)
	left := 0
	maxLength := 0
	resultStart := 0
	
	for right := 0; right < len(s); right++ {
		// Expand window
		charCount[s[right]]++
		
		// Contract window if more than k unique characters
		for len(charCount) > k {
			charCount[s[left]]--
			if charCount[s[left]] == 0 {
				delete(charCount, s[left])
			}
			left++
		}
		
		// Update max length only if we have exactly k unique characters
		if len(charCount) == k && right-left+1 > maxLength {
			maxLength = right - left + 1
			resultStart = left
		}
	}
	
	if maxLength == 0 {
		return ""
	}
	return s[resultStart : resultStart+maxLength]
}

// Approach 3: Find all substrings with exactly k unique characters
func allSubstringsWithKUniqueChars(s string, k int) []string {
	if len(s) == 0 || k == 0 {
		return []string{}
	}
	
	var result []string
	seen := make(map[string]bool)
	
	for i := 0; i < len(s); i++ {
		charCount := make(map[byte]int)
		for j := i; j < len(s); j++ {
			charCount[s[j]]++
			
			if len(charCount) == k {
				substring := s[i : j+1]
				if !seen[substring] {
					result = append(result, substring)
					seen[substring] = true
				}
			} else if len(charCount) > k {
				break
			}
		}
	}
	
	return result
}

// Approach 4: Using sliding window with character frequency array (for lowercase letters)
func longestSubstringWithKUniqueCharsArray(s string, k int) int {
	if len(s) == 0 || k == 0 {
		return 0
	}
	
	charCount := make([]int, 26) // For lowercase a-z
	left := 0
	maxLength := 0
	uniqueCount := 0
	
	for right := 0; right < len(s); right++ {
		// Expand window
		rightChar := s[right] - 'a'
		if charCount[rightChar] == 0 {
			uniqueCount++
		}
		charCount[rightChar]++
		
		// Contract window if more than k unique characters
		for uniqueCount > k {
			leftChar := s[left] - 'a'
			charCount[leftChar]--
			if charCount[leftChar] == 0 {
				uniqueCount--
			}
			left++
		}
		
		// Update max length only if we have exactly k unique characters
		if uniqueCount == k {
			maxLength = max(maxLength, right-left+1)
		}
	}
	
	return maxLength
}

// Helper function to validate if string has exactly k unique characters
func hasExactlyKUniqueChars(s string, k int) bool {
	charSet := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		charSet[s[i]] = true
	}
	return len(charSet) == k
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("=== Longest Substring with K Unique Characters ===")
	
	// Test case 1
	s1 := "aabbcc"
	k1 := 2
	fmt.Printf("\nTest Case 1:\n")
	fmt.Printf("String: %s, K: %d\n", s1, k1)
	fmt.Printf("Longest length: %d\n", longestSubstringWithKUniqueChars(s1, k1))
	fmt.Printf("Longest substring: %s\n", longestSubstringWithKUniqueCharsString(s1, k1))
	
	// Test case 2
	s2 := "aabbccddeeffgg"
	k2 := 3
	fmt.Printf("\nTest Case 2:\n")
	fmt.Printf("String: %s, K: %d\n", s2, k2)
	fmt.Printf("Longest length: %d\n", longestSubstringWithKUniqueChars(s2, k2))
	fmt.Printf("Longest substring: %s\n", longestSubstringWithKUniqueCharsString(s2, k2))
	
	// Test case 3
	s3 := "abcba"
	k3 := 2
	fmt.Printf("\nTest Case 3:\n")
	fmt.Printf("String: %s, K: %d\n", s3, k3)
	fmt.Printf("Longest length: %d\n", longestSubstringWithKUniqueChars(s3, k3))
	fmt.Printf("Longest substring: %s\n", longestSubstringWithKUniqueCharsString(s3, k3))
	
	// Test case 4: All substrings
	s4 := "aabacbebebe"
	k4 := 3
	fmt.Printf("\nTest Case 4 - All substrings:\n")
	fmt.Printf("String: %s, K: %d\n", s4, k4)
	fmt.Printf("Longest length: %d\n", longestSubstringWithKUniqueChars(s4, k4))
	fmt.Printf("Longest substring: %s\n", longestSubstringWithKUniqueCharsString(s4, k4))
	
	allSubs := allSubstringsWithKUniqueChars(s4, k4)
	fmt.Printf("All substrings with exactly %d unique chars:\n", k4)
	for i, sub := range allSubs {
		fmt.Printf("  %d: %s (length: %d, valid: %v)\n", i+1, sub, len(sub), hasExactlyKUniqueChars(sub, k4))
	}
	
	// Test case 5: Edge cases
	fmt.Printf("\nEdge Cases:\n")
	
	// Empty string
	fmt.Printf("Empty string, k=2: %d\n", longestSubstringWithKUniqueChars("", 2))
	
	// K = 0
	fmt.Printf("String 'abc', k=0: %d\n", longestSubstringWithKUniqueChars("abc", 0))
	
	// K > unique characters in string
	fmt.Printf("String 'aab', k=5: %d\n", longestSubstringWithKUniqueChars("aab", 5))
	
	// Single character repeated
	fmt.Printf("String 'aaaa', k=1: %d\n", longestSubstringWithKUniqueChars("aaaa", 1))
	
	// K = 1
	fmt.Printf("String 'abcdef', k=1: %d\n", longestSubstringWithKUniqueChars("abcdef", 1))
	
	// All unique characters
	fmt.Printf("String 'abcdef', k=6: %d\n", longestSubstringWithKUniqueChars("abcdef", 6))
	
	// Performance comparison
	fmt.Printf("\nPerformance Comparison:\n")
	longString := "abcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabc"
	fmt.Printf("Long string length: %d\n", len(longString))
	fmt.Printf("HashMap approach: %d\n", longestSubstringWithKUniqueChars(longString, 3))
	fmt.Printf("Array approach: %d\n", longestSubstringWithKUniqueCharsArray(longString, 3))
}