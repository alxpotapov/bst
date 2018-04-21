package bst

import "testing"

func TestTree_Iterator(t *testing.T) {
	tree := &Tree{}

	tree.Insert("2", 2)
	tree.Insert("1", 1)
	tree.Insert("7", 7)
	tree.Insert("6", 6)
	tree.Insert("8", 8)
	tree.Insert("3", 3)

	it := tree.Iterator()
	for node := range it{
		t.Log(node.key, ":", node.value)
	}
}
