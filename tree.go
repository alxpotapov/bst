package bst


type Tree struct {
	root *Node
}

func(t *Tree)Insert(key string, value interface{}) {
	if t.root == nil {
		t.root = &Node{key:key, value:value}
		return
	}
	t.root.Insert(key, value)
}

func (t *Tree)Find(key string)(interface{}, bool) {
	if t.root == nil {
		return nil, false
	}
	return t.root.Find(key)
}


func(t *Tree)Delete(key string) {
	if t.root == nil {
		return
	}
	t.root = t.root.Delete(key)
}


func(t *Tree)inOrder(node *Node, it chan <- *Node) {
	if node.left != nil {
		t.inOrder(node.left, it)
	}
	it <- node
	if node.right != nil {
		t.inOrder(node.right, it)
	}
}

func (t *Tree) Iterator() <- chan *Node {
	it := make(chan *Node)
	go func(){
		t.inOrder(t.root, it)
		close(it)
	}()
	return it
}
