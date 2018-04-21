package bst

// Tree - a binary search tree
type Tree struct {
	root *Node
}

// Insert node
func (t *Tree) Insert(key string, value interface{}) {
	if t.root == nil {
		t.root = &Node{Key: key, Value: value}
		return
	}
	t.root.Insert(key, value)
}

// Find node
func (t *Tree) Find(key string) (interface{}, bool) {
	if t.root == nil {
		return nil, false
	}
	return t.root.Find(key)
}

// Delete node
func (t *Tree) Delete(key string) {
	if t.root == nil {
		return
	}
	t.root = t.root.Delete(key)
}

// inOrder - in-order way walking in tree and send node to chan
func (t *Tree) inOrder(node *Node, it chan<- *Node) {
	if node.left != nil {
		t.inOrder(node.left, it)
	}
	it <- node
	if node.right != nil {
		t.inOrder(node.right, it)
	}
}

// ChannelIterator ...
func (t *Tree) ChannelIterator() <-chan *Node {
	it := make(chan *Node)
	go func() {
		defer close(it)
		t.inOrder(t.root, it)
	}()
	return it
}

func (t *Tree) StatefulIterator() *StatefulIterator {
	return NewStatefulIterator(t)
}

// StatefulIterator hold the iteration state in the iterator struct itself
type StatefulIterator struct {
	current int
	nodes   []*Node
}

// NewStatefulIterator ...
func NewStatefulIterator(tree *Tree) *StatefulIterator {
	it := &StatefulIterator{current: -1, nodes: make([]*Node, 0, 16)}
	it.push(tree.root)
	return it
}

func (s *StatefulIterator) push(node *Node) {
	if node.left != nil {
		s.push(node.left)
	}
	s.nodes = append(s.nodes, node)
	if node.right != nil {
		s.push(node.right)
	}
}

// Next advances iterator to next value. It returns false to indicate end of iteration
func (s *StatefulIterator) Next() bool {
	s.current++
	if len(s.nodes) > s.current {
		return true
	}
	return false
}

// Node to get the current pointer of node of the iterator
func (s *StatefulIterator) Node() *Node {
	return s.nodes[s.current]
}
