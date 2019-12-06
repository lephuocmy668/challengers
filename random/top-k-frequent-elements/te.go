package te

// Runtime: 20 ms, faster than 26.41% of Go online submissions for Top K Frequent Elements.
// https://leetcode.com/problems/top-k-frequent-elements/submissions/
// TODO: solve DRY
func topKFrequent(nums []int, k int) []int {
	numberMap := make(map[int]*Element)
	nodeMap := make(map[int]*Node)
	lk := NewLinkedList()

	for _, num := range nums {
		if numberElement, ok := numberMap[num]; ok {
			node := nodeMap[numberElement.count]
			delete(node.Elements, numberElement)

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
