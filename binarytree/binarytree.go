package binarytree

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeStack struct {
	nodes []*TreeNode
}

func NewTreeNodeStack() *TreeNodeStack {
	return &TreeNodeStack{
		nodes: []*TreeNode{},
	}
}

func (s *TreeNodeStack) IsEmpty() bool {
	return len(s.nodes) == 0
}

func (s *TreeNodeStack) Push(node *TreeNode) {
	s.nodes = append(s.nodes, node)
}

func (s *TreeNodeStack) Pop() *TreeNode {
	result := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return result
}

func (s *TreeNodeStack) Peek() *TreeNode {
	return s.nodes[len(s.nodes)-1]
}
