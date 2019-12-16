package wd

type WordDictionary struct {
	root *TrieNode
}

type TrieNode struct {
	isEnd    bool
	char     rune
	children [26]*TrieNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	root := TrieNode{isEnd: false, char: 0, children: [26]*TrieNode{}}
	return WordDictionary{
		root: &root,
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &TrieNode{isEnd: false, char: c, children: [26]*TrieNode{}}
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
	}

	next := node.children[word[index]-'a']
	return next != nil && search(word, index+1, next)
}
