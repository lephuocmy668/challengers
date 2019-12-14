package trie

// Runtime: 72 ms, faster than 73.72% of Go online submissions for Implement Trie (Prefix Tree).
// Memory Usage: 19 MB, less than 100.00% of Go online submissions for Implement Trie (Prefix Tree).
// https://leetcode.com/problems/implement-trie-prefix-tree/
type Trie struct {
	root *trieNode
}

type trieNode struct {
	children map[rune]*trieNode
	isLeaf   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: &trieNode{
			children: make(map[rune]*trieNode),
		},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	r := []rune(word)
	node := this.root

	for _, c := range r {
		if n, ok := node.children[c]; ok {
			node = n
			continue
		}

		newNode := &trieNode{
			children: make(map[rune]*trieNode),
		}
		node.children[c] = newNode
		node = newNode
	}

	node.isLeaf = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	r := []rune(word)
	node := this.root

	for _, c := range r {
		if n, ok := node.children[c]; ok {
			node = n
			continue
		}
		return false
	}

	return node.isLeaf
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	r := []rune(prefix)
	node := this.root

	for _, c := range r {
		if n, ok := node.children[c]; ok {
			node = n
			continue
		}
		return false
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
