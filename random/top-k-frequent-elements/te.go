package te

// Runtime: 24 ms, faster than 16.90% of Go online submissions for Top K Frequent Elements.
// https://leetcode.com/problems/top-k-frequent-elements/submissions/
func topKFrequent(nums []int, k int) []int {
	kmap := make(map[int]*Element)
	nodeMap := make(map[int]*Node)
	lk := NewLinkedList()

	for _, num := range nums {
		if k, ok := kmap[num]; ok {
			node := nodeMap[k.count]
			element, ok := node.Elements[k]
			if !ok {
			}
			delete(node.Elements, k)

			k.count++
			if node, ok := nodeMap[k.count]; ok {
				node.Elements[k] = k
			} else {
				node := &Node{
					Elements: map[*Element]*Element{
						element: element,
					},
				}
				nodeMap[k.count] = node
				lk.Append(node)
			}

		} else {
			count := 1
			k := &Element{
				k:     num,
				count: count,
			}
			kmap[num] = k

			if node, ok := nodeMap[count]; ok {
				node.Elements[k] = k
			} else {
				node := &Node{
					Elements: map[*Element]*Element{
						k: k,
					},
				}
				nodeMap[count] = node
				lk.Prepend(node)
			}
		}
	}

	return lk.ToResult(k)
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
