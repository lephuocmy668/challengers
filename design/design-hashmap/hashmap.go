package hashmap

// Assignment:
// https://leetcode.com/problems/design-hashmap/submissions/

// Simple implementation of hashmap, for more effective solution
// Refer: http://www.algolist.net/Data_structures/Hash_table
// http://www.algolist.net/Data_structures/Hash_table/Chaining
const (
	ArraySize = 100000
)

func hashFunc(key int) int {
	return key % ArraySize
}

// Node ...
type Node struct {
	Key   int
	Value int
	Next  *Node
}

// LinkedList ...
type LinkedList struct {
	Head *Node
	Tail *Node
}

// NewLinkedList returns new linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// IsEmpty ...
func (lk *LinkedList) IsEmpty() bool {
	return lk.Head == nil
}

// FindByKey ...
func (lk *LinkedList) FindByKey(key int) *Node {
	if lk == nil || lk.Head == nil {
		return nil
	}

	node := lk.Head
	for node != nil {
		if node.Key == key {
			return node
		}
		node = node.Next
	}
	return nil
}

// DeleteByKey ...
func (lk *LinkedList) DeleteByKey(key int) {
	if lk == nil || lk.Head == nil {
		return
	}

	var prev *Node
	curr := lk.Head
	for curr != nil {
		if curr.Key == key {
			if prev != nil {
				prev.Next = curr.Next
			} else {
				lk.Head = curr.Next
			}
		}
		prev = curr
		curr = curr.Next
	}
}

// Append ...
func (lk *LinkedList) Append(newNode *Node) {
	if lk.Head == nil {
		lk.Head = newNode
		lk.Tail = newNode
		return
	}

	lk.Tail.Next = newNode
	lk.Tail = newNode
}

type MyHashMap struct {
	Clusters []*LinkedList
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		Clusters: make([]*LinkedList, ArraySize),
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	index := hashFunc(key)
	newNode := &Node{
		Key:   key,
		Value: value,
	}

	if this.Clusters[index] == nil {
		this.Clusters[index] = NewLinkedList()
		this.Clusters[index].Append(newNode)
		return
	}

	node := this.Clusters[index].FindByKey(key)
	if node == nil {
		this.Clusters[index].Append(newNode)
		return
	}

	node.Value = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	node := this.Clusters[hashFunc(key)].FindByKey(key)
	if node != nil {
		return node.Value
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.Clusters[hashFunc(key)].DeleteByKey(key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
