package te

// Runtime: 24 ms, faster than 16.90% of Go online submissions for Top K Frequent Elements.
// https://leetcode.com/problems/top-k-frequent-elements/submissions/
func topKFrequent(nums []int, k int) []int {
	numberMap := make(map[int]*Element)
	nodeMap := make(map[int]*Node)
	lk := NewLinkedList()

	for _, num := range nums {
		if numberElement, ok := numberMap[num]; ok {
			node := nodeMap[numberElement.count]
			element := node.Elements[numberElement]
			delete(node.Elements, numberElement)

			numberElement.count++
			if node, ok := nodeMap[numberElement.count]; ok {
				node.Elements[numberElement] = numberElement
			} else {
				node := &Node{
					Elements: map[*Element]*Element{
						element: element,
					},
				}
				nodeMap[numberElement.count] = node
				lk.Append(node)
			}

		} else {
			numberElement := &Element{
				k: num,
			}
			numberMap[num] = numberElement

			numberElement.count++
			if node, ok := nodeMap[numberElement.count]; ok {
				node.Elements[numberElement] = numberElement
			} else {
				node := &Node{
					Elements: map[*Element]*Element{
						numberElement: numberElement,
					},
				}
				nodeMap[numberElement.count] = node
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
