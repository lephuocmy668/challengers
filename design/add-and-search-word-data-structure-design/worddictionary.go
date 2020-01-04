package worddictionary

// https://leetcode.com/problems/add-and-search-word-data-structure-design
// Runtime: 104 ms, faster than 39.58% of Go online submissions for Add and Search Word - Data structure design.
// Memory Usage: 22.8 MB, less than 50.00% of Go online submissions for Add and Search Word - Data structure design.
type WordDictionary struct {
	root *TrieNode
}

type TrieNode struct {
	isEnd    bool
	char     rune
	children [26]*TrieNode
}

func NewTrieNode(isEnd bool, char rune, children [26]*TrieNode) TrieNode {
	return TrieNode{
		isEnd:    isEnd,
		char:     char,
		children: children,
	}
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	root := NewTrieNode(false, 0, [26]*TrieNode{})
	return WordDictionary{
		root: &root,
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, c := range word {
		childNode := NewTrieNode(false, c, [26]*TrieNode{})
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &childNode
		}
		node = node.children[c-'a']
	}
	node.isEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return search(word, 0, this.root)
}

func search(word string, index int, node *TrieNode) bool {
	if index == len(word) {
		return node.isEnd
	}

	if word[index] == '.' {
		for _, child := range node.children {
			if child != nil && search(word, index+1, child) {
				return true
			}
		}
		return false
	} else {
		next := node.children[word[index]-'a']
		return next != nil && search(word, index+1, next)
	}
}
