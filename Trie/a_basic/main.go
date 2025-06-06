package basic

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	value    interface{}
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
		value:    nil,
	}
}

type Trie struct {
	root *TrieNode
	size int
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
		size: 0,
	}
}

func (t *Trie) InsertWithValue(word string, value interface{}) {
	node := t.root

	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = NewTrieNode()
		}

		node = node.children[char]
	}

	if !node.isEnd {
		t.size++
	}

	node.isEnd = true
	node.value = value
}

func (t *Trie) Insert(word string) {
	t.InsertWithValue(word, nil)
}

func (t *Trie) findNode(str string) *TrieNode {
	node := t.root

	for _, char := range str {
		if child, exists := node.children[char]; exists {
			node = child
		} else {
			return nil
		}
	}

	return node
}

func (t *Trie) Search(word string) bool {
	node := t.findNode(word)

	return node != nil && node.isEnd
}

func (t *Trie) SearchPrefix(word string) bool {
	return t.findNode(word) != nil
}
