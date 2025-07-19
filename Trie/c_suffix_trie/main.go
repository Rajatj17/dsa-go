package main

import (
	"fmt"
	"strings"
)

// Problem: Suffix Trie Implementation
// Build a trie of all suffixes of a string for efficient substring search
// Time Complexity: O(n²) for construction, O(m) for search where n is string length, m is pattern length
// Space Complexity: O(n²) in worst case

type SuffixTrieNode struct {
	children map[rune]*SuffixTrieNode
	isEnd    bool
	indexes  []int // Store all starting positions of suffixes ending at this node
}

type SuffixTrie struct {
	root *SuffixTrieNode
	text string
}

func NewSuffixTrie(text string) *SuffixTrie {
	trie := &SuffixTrie{
		root: &SuffixTrieNode{
			children: make(map[rune]*SuffixTrieNode),
			isEnd:    false,
			indexes:  []int{},
		},
		text: text,
	}
	
	// Build the suffix trie
	trie.buildSuffixTrie()
	return trie
}

// Build suffix trie by inserting all suffixes
func (st *SuffixTrie) buildSuffixTrie() {
	text := st.text
	for i := 0; i < len(text); i++ {
		st.insertSuffix(text[i:], i)
	}
}

// Insert a suffix starting at given index
func (st *SuffixTrie) insertSuffix(suffix string, startIndex int) {
	node := st.root
	for _, char := range suffix {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &SuffixTrieNode{
				children: make(map[rune]*SuffixTrieNode),
				isEnd:    false,
				indexes:  []int{},
			}
		}
		node = node.children[char]
		node.indexes = append(node.indexes, startIndex)
	}
	node.isEnd = true
}

// Search for a pattern in the text
func (st *SuffixTrie) Search(pattern string) bool {
	node := st.root
	for _, char := range pattern {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return true
}

// Find all occurrences of a pattern
func (st *SuffixTrie) FindAllOccurrences(pattern string) []int {
	node := st.root
	for _, char := range pattern {
		if _, exists := node.children[char]; !exists {
			return []int{}
		}
		node = node.children[char]
	}
	
	// Return all indexes where this pattern occurs
	var result []int
	for _, index := range node.indexes {
		if index+len(pattern) <= len(st.text) {
			result = append(result, index)
		}
	}
	return result
}

// Count occurrences of a pattern
func (st *SuffixTrie) CountOccurrences(pattern string) int {
	return len(st.FindAllOccurrences(pattern))
}

// Find the longest repeated substring
func (st *SuffixTrie) FindLongestRepeatedSubstring() string {
	longestSubstring := ""
	st.findLongestRepeatedHelper(st.root, "", &longestSubstring)
	return longestSubstring
}

func (st *SuffixTrie) findLongestRepeatedHelper(node *SuffixTrieNode, current string, longest *string) {
	// If this node has more than one index, it means this substring appears multiple times
	if len(node.indexes) > 1 && len(current) > len(*longest) {
		*longest = current
	}
	
	for char, child := range node.children {
		st.findLongestRepeatedHelper(child, current+string(char), longest)
	}
}

// Find all substrings that occur more than once
func (st *SuffixTrie) FindRepeatedSubstrings() []string {
	var result []string
	st.findRepeatedHelper(st.root, "", &result)
	return result
}

func (st *SuffixTrie) findRepeatedHelper(node *SuffixTrieNode, current string, result *[]string) {
	if len(node.indexes) > 1 && current != "" {
		*result = append(*result, current)
	}
	
	for char, child := range node.children {
		st.findRepeatedHelper(child, current+string(char), result)
	}
}

// Check if a string is a substring of the text
func (st *SuffixTrie) IsSubstring(substring string) bool {
	return st.Search(substring)
}

// Find the longest common substring between two strings
func FindLongestCommonSubstring(str1, str2 string) string {
	// Create suffix trie for first string
	trie := NewSuffixTrie(str1)
	
	longestCommon := ""
	
	// Check all suffixes of second string
	for i := 0; i < len(str2); i++ {
		for j := i + 1; j <= len(str2); j++ {
			substring := str2[i:j]
			if trie.IsSubstring(substring) && len(substring) > len(longestCommon) {
				longestCommon = substring
			}
		}
	}
	
	return longestCommon
}

// Get all suffixes of the text
func (st *SuffixTrie) GetAllSuffixes() []string {
	var suffixes []string
	text := st.text
	for i := 0; i < len(text); i++ {
		suffixes = append(suffixes, text[i:])
	}
	return suffixes
}

// Print the suffix trie structure
func (st *SuffixTrie) PrintTrie() {
	fmt.Println("Suffix Trie Structure:")
	st.printHelper(st.root, "", "")
}

func (st *SuffixTrie) printHelper(node *SuffixTrieNode, prefix, indent string) {
	if len(node.indexes) > 0 {
		fmt.Printf("%s%s [Indexes: %v]\n", indent, prefix, node.indexes)
	}
	
	for char, child := range node.children {
		st.printHelper(child, prefix+string(char), indent+"  ")
	}
}

// Find all palindromic substrings using suffix trie
func (st *SuffixTrie) FindPalindromicSubstrings() []string {
	var palindromes []string
	text := st.text
	
	// Check all possible substrings
	for i := 0; i < len(text); i++ {
		for j := i + 1; j <= len(text); j++ {
			substring := text[i:j]
			if st.isPalindrome(substring) {
				palindromes = append(palindromes, substring)
			}
		}
	}
	
	return palindromes
}

func (st *SuffixTrie) isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Find the shortest unique substring that contains a given character
func (st *SuffixTrie) FindShortestUniqueSubstring(char rune) string {
	text := st.text
	shortest := ""
	
	for i := 0; i < len(text); i++ {
		if rune(text[i]) == char {
			for j := i + 1; j <= len(text); j++ {
				substring := text[i:j]
				if st.CountOccurrences(substring) == 1 {
					if shortest == "" || len(substring) < len(shortest) {
						shortest = substring
					}
					break
				}
			}
		}
	}
	
	return shortest
}

func main() {
	fmt.Println("=== Suffix Trie Implementation ===")
	
	// Test Case 1: Basic operations
	text1 := "banana"
	fmt.Printf("\n1. Building suffix trie for: %s\n", text1)
	trie1 := NewSuffixTrie(text1)
	
	fmt.Printf("All suffixes: %v\n", trie1.GetAllSuffixes())
	
	// Search for patterns
	fmt.Println("\n2. Pattern searching:")
	patterns := []string{"an", "ana", "nana", "ban", "xyz"}
	for _, pattern := range patterns {
		found := trie1.Search(pattern)
		occurrences := trie1.FindAllOccurrences(pattern)
		fmt.Printf("Pattern '%s': Found=%v, Occurrences=%v, Count=%d\n", 
			pattern, found, occurrences, len(occurrences))
	}
	
	// Find repeated substrings
	fmt.Println("\n3. Repeated substrings:")
	fmt.Printf("Longest repeated substring: '%s'\n", trie1.FindLongestRepeatedSubstring())
	repeated := trie1.FindRepeatedSubstrings()
	fmt.Printf("All repeated substrings: %v\n", repeated)
	
	// Test Case 2: More complex text
	text2 := "abcabcabc"
	fmt.Printf("\n4. Building suffix trie for: %s\n", text2)
	trie2 := NewSuffixTrie(text2)
	
	fmt.Printf("Longest repeated substring: '%s'\n", trie2.FindLongestRepeatedSubstring())
	fmt.Printf("All repeated substrings: %v\n", trie2.FindRepeatedSubstrings())
	
	// Test Case 3: DNA sequence analysis
	dna := "ATCGATCGATCG"
	fmt.Printf("\n5. DNA sequence analysis: %s\n", dna)
	dnaTrie := NewSuffixTrie(dna)
	
	// Find common DNA patterns
	dnaPatterns := []string{"AT", "CG", "ATCG", "TCGA", "GATC"}
	for _, pattern := range dnaPatterns {
		occurrences := dnaTrie.FindAllOccurrences(pattern)
		fmt.Printf("DNA pattern '%s': %v (count: %d)\n", pattern, occurrences, len(occurrences))
	}
	
	// Test Case 4: Longest common substring
	fmt.Println("\n6. Longest common substring:")
	str1 := "abcdefgh"
	str2 := "xyzabcpqr"
	lcs := FindLongestCommonSubstring(str1, str2)
	fmt.Printf("LCS of '%s' and '%s': '%s'\n", str1, str2, lcs)
	
	// Test Case 5: Palindromes
	palindromeText := "racecar"
	fmt.Printf("\n7. Palindromic substrings in '%s':\n", palindromeText)
	palindromeTrie := NewSuffixTrie(palindromeText)
	palindromes := palindromeTrie.FindPalindromicSubstrings()
	fmt.Printf("Palindromes: %v\n", palindromes)
	
	// Test Case 6: Shortest unique substring
	fmt.Println("\n8. Shortest unique substrings:")
	uniqueText := "abcabc"
	uniqueTrie := NewSuffixTrie(uniqueText)
	for _, char := range "abc" {
		shortest := uniqueTrie.FindShortestUniqueSubstring(rune(char))
		fmt.Printf("Shortest unique substring containing '%c': '%s'\n", char, shortest)
	}
	
	// Test Case 7: Performance test
	fmt.Println("\n9. Performance test:")
	longText := strings.Repeat("abcdefghijklmnopqrstuvwxyz", 10)
	fmt.Printf("Building suffix trie for text of length: %d\n", len(longText))
	longTrie := NewSuffixTrie(longText)
	
	testPattern := "xyz"
	occurrences := longTrie.FindAllOccurrences(testPattern)
	fmt.Printf("Pattern '%s' found %d times\n", testPattern, len(occurrences))
	
	// Print trie structure for small example
	fmt.Println("\n10. Trie structure for 'abc':")
	smallTrie := NewSuffixTrie("abc")
	smallTrie.PrintTrie()
}