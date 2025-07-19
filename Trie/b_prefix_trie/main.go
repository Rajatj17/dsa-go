package main

import "fmt"

// Problem: Prefix Trie (Auto-complete, Word Suggestions)
// Implement a trie with prefix search functionality
// Time Complexity: O(m) for insert, search, and startsWith where m is key length
// Space Complexity: O(ALPHABET_SIZE * N * M) where N is number of keys, M is average length

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	word     string // Store the complete word for easier retrieval
}

type PrefixTrie struct {
	root *TrieNode
}

func NewPrefixTrie() *PrefixTrie {
	return &PrefixTrie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

// Insert a word into the trie
func (t *PrefixTrie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		node = node.children[char]
	}
	node.isEnd = true
	node.word = word
}

// Search for a complete word in the trie
func (t *PrefixTrie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

// Check if there's any word that starts with the given prefix
func (t *PrefixTrie) StartsWith(prefix string) bool {
	node := t.root
	for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return true
}

// Get all words that start with the given prefix
func (t *PrefixTrie) GetWordsWithPrefix(prefix string) []string {
	var result []string
	node := t.root
	
	// Navigate to the prefix node
	for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return result // Empty result if prefix doesn't exist
		}
		node = node.children[char]
	}
	
	// DFS to collect all words from this node
	t.dfs(node, prefix, &result)
	return result
}

// Helper function for DFS traversal
func (t *PrefixTrie) dfs(node *TrieNode, currentWord string, result *[]string) {
	if node.isEnd {
		*result = append(*result, currentWord)
	}
	
	for char, child := range node.children {
		t.dfs(child, currentWord+string(char), result)
	}
}

// Get all words that start with the given prefix (alternative approach using stored words)
func (t *PrefixTrie) GetWordsWithPrefixOptimized(prefix string) []string {
	var result []string
	node := t.root
	
	// Navigate to the prefix node
	for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return result
		}
		node = node.children[char]
	}
	
	// DFS to collect all words from this node
	t.dfsOptimized(node, &result)
	return result
}

// Helper function for optimized DFS (using stored words)
func (t *PrefixTrie) dfsOptimized(node *TrieNode, result *[]string) {
	if node.isEnd {
		*result = append(*result, node.word)
	}
	
	for _, child := range node.children {
		t.dfsOptimized(child, result)
	}
}

// Delete a word from the trie
func (t *PrefixTrie) Delete(word string) bool {
	return t.deleteHelper(t.root, word, 0)
}

func (t *PrefixTrie) deleteHelper(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		// We've reached the end of the word
		if !node.isEnd {
			return false // Word doesn't exist
		}
		node.isEnd = false
		node.word = ""
		// Return true if current node has no children (can be deleted)
		return len(node.children) == 0
	}
	
	char := rune(word[index])
	child, exists := node.children[char]
	if !exists {
		return false // Word doesn't exist
	}
	
	shouldDeleteChild := t.deleteHelper(child, word, index+1)
	
	if shouldDeleteChild {
		delete(node.children, char)
		// Return true if current node has no children and is not end of another word
		return len(node.children) == 0 && !node.isEnd
	}
	
	return false
}

// Get the longest common prefix of all words in the trie
func (t *PrefixTrie) GetLongestCommonPrefix() string {
	var prefix string
	node := t.root
	
	for len(node.children) == 1 && !node.isEnd {
		for char, child := range node.children {
			prefix += string(char)
			node = child
		}
	}
	
	return prefix
}

// Count total number of words in the trie
func (t *PrefixTrie) CountWords() int {
	return t.countWordsHelper(t.root)
}

func (t *PrefixTrie) countWordsHelper(node *TrieNode) int {
	count := 0
	if node.isEnd {
		count = 1
	}
	
	for _, child := range node.children {
		count += t.countWordsHelper(child)
	}
	
	return count
}

// Get all words in the trie
func (t *PrefixTrie) GetAllWords() []string {
	var result []string
	t.dfs(t.root, "", &result)
	return result
}

// Print the trie structure
func (t *PrefixTrie) PrintTrie() {
	fmt.Println("Trie Structure:")
	t.printHelper(t.root, "", "")
}

func (t *PrefixTrie) printHelper(node *TrieNode, prefix, indent string) {
	if node.isEnd {
		fmt.Printf("%s%s [END: %s]\n", indent, prefix, node.word)
	} else if prefix != "" {
		fmt.Printf("%s%s\n", indent, prefix)
	}
	
	for char, child := range node.children {
		t.printHelper(child, string(char), indent+"  ")
	}
}

func main() {
	fmt.Println("=== Prefix Trie Implementation ===")
	
	// Create a new trie
	trie := NewPrefixTrie()
	
	// Test data
	words := []string{
		"apple", "app", "apricot", "application", "apply",
		"banana", "band", "bandana", "can", "cat", "car", "card", "care", "careful",
	}
	
	// Insert words
	fmt.Println("\n1. Inserting words:")
	for _, word := range words {
		trie.Insert(word)
		fmt.Printf("Inserted: %s\n", word)
	}
	
	// Search for words
	fmt.Println("\n2. Searching for words:")
	searchWords := []string{"app", "apple", "application", "appl", "banana", "xyz"}
	for _, word := range searchWords {
		found := trie.Search(word)
		fmt.Printf("Search '%s': %v\n", word, found)
	}
	
	// Check prefixes
	fmt.Println("\n3. Checking prefixes:")
	prefixes := []string{"app", "appl", "ban", "ca", "xyz"}
	for _, prefix := range prefixes {
		exists := trie.StartsWith(prefix)
		fmt.Printf("Prefix '%s' exists: %v\n", prefix, exists)
	}
	
	// Get words with prefix
	fmt.Println("\n4. Words with prefixes:")
	testPrefixes := []string{"app", "ban", "ca", "car"}
	for _, prefix := range testPrefixes {
		words := trie.GetWordsWithPrefix(prefix)
		fmt.Printf("Words with prefix '%s': %v\n", prefix, words)
	}
	
	// Auto-complete functionality
	fmt.Println("\n5. Auto-complete suggestions:")
	autoCompletePrefixes := []string{"ap", "ca", "ba"}
	for _, prefix := range autoCompletePrefixes {
		suggestions := trie.GetWordsWithPrefixOptimized(prefix)
		fmt.Printf("Auto-complete for '%s': %v\n", prefix, suggestions)
	}
	
	// Statistics
	fmt.Println("\n6. Trie statistics:")
	fmt.Printf("Total words: %d\n", trie.CountWords())
	fmt.Printf("Longest common prefix: '%s'\n", trie.GetLongestCommonPrefix())
	fmt.Printf("All words: %v\n", trie.GetAllWords())
	
	// Delete words
	fmt.Println("\n7. Deleting words:")
	deleteWords := []string{"app", "bandana", "xyz"}
	for _, word := range deleteWords {
		deleted := trie.Delete(word)
		fmt.Printf("Delete '%s': %v\n", word, deleted)
	}
	
	// Verify deletions
	fmt.Println("\n8. Verifying deletions:")
	for _, word := range deleteWords {
		found := trie.Search(word)
		fmt.Printf("Search '%s' after deletion: %v\n", word, found)
	}
	
	// Final statistics
	fmt.Println("\n9. Final statistics:")
	fmt.Printf("Total words after deletion: %d\n", trie.CountWords())
	fmt.Printf("All words after deletion: %v\n", trie.GetAllWords())
	
	// Print trie structure
	fmt.Println("\n10. Trie structure:")
	trie.PrintTrie()
	
	// Practical use case: Search suggestions
	fmt.Println("\n11. Practical use case - Search suggestions:")
	searchTrie := NewPrefixTrie()
	searchTerms := []string{
		"javascript", "java", "python", "go", "golang", "react", "angular", "vue",
		"nodejs", "express", "mongodb", "mysql", "postgresql", "redis",
	}
	
	for _, term := range searchTerms {
		searchTrie.Insert(term)
	}
	
	userInputs := []string{"ja", "go", "py", "re", "no"}
	for _, input := range userInputs {
		suggestions := searchTrie.GetWordsWithPrefix(input)
		fmt.Printf("User typed '%s' -> Suggestions: %v\n", input, suggestions)
	}
}