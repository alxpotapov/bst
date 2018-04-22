package bst

import "testing"

func TestTree_ChannelIterator(t *testing.T) {
	tree := &Tree{}

	tree.Insert("2", 2)
	tree.Insert("1", 1)
	tree.Insert("7", 7)
	tree.Insert("6", 6)
	tree.Insert("8", 8)
	tree.Insert("3", 3)

	for node := range tree.ChannelIterator() {
		t.Log(node.Key, ":", node.Value)
	}
}

func TestTree_StatefulIterator(t *testing.T) {
	tree := &Tree{}

	tree.Insert("2", 2)
	tree.Insert("1", 1)
	tree.Insert("7", 7)
	tree.Insert("6", 6)
	tree.Insert("8", 8)
	tree.Insert("3", 3)

	for it := tree.StatefulIterator(); it.Next(); {
		t.Log(it.Node().Key, ":", it.Node().Value)
	}
}


func TestTree_ForEach(t *testing.T) {
	tree := &Tree{}

	tree.Insert("2", 2)
	tree.Insert("1", 1)
	tree.Insert("7", 7)
	tree.Insert("6", 6)
	tree.Insert("8", 8)
	tree.Insert("3", 3)

	tree.ForEach(func(node *Node){
		t.Log(node.Key, ":", node.Value)
	})

}