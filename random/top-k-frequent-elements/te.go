package te

func topKFrequent(nums []int, k int) []int {
	kmap := make(map[int]*Element)
	nodeMap := make(map[int]*Node)

	return []int{}
}

// Element ...
type Element struct {
	k     int
	count int
}

// Node ...
type Node struct {
	Next     *Node
	Prev     *Node
	Elements map[*Element]*Element
}
