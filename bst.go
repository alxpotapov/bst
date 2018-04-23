// Package bst contain a implementation of binary search tree
package bst

import (
	"strings"
	"fmt"
)
// Callback ...
type Callback func(node *Node)

// Node contains the key as string, value as interface{} and 2 pointers to his children
type Node struct {
	left, right *Node
	Key         string
	Value       interface{}
}

// Insert a key and value into a sorted tree
func (n *Node) Insert(key string, value interface{}) {
	comp := strings.Compare(n.Key, key)
	switch {
	case comp > 0:
		if n.left != nil {
			n.left.Insert(key, value)
		} else {
			n.left = &Node{Key: key, Value: value}
		}
	case comp < 0:
		if n.right != nil {
			n.right.Insert(key, value)
		} else {
			n.right = &Node{Key: key, Value: value}
		}
	default:
		n.Value = value
	}
}

// Find value by key
func (n *Node) Find(key string)(interface{},  bool) {
	comp := strings.Compare(n.Key, key)
	switch {
	case comp > 0:
		if n.left != nil {
			return n.left.Find(key)
		}
	case comp < 0:
		if n.right != nil {
			return n.right.Find(key)
		}

	default:
		return n.Value, true
	}
	return nil, false
}

// Delete a value from a sorted tree by key
func (n *Node) Delete(key string) *Node {
	comp := strings.Compare(n.Key, key)
	switch {

	case comp > 0:
		if n.left != nil {
			n.left = n.left.Delete(key)
		}
	case comp < 0:
		if n.right != nil {
			n.right = n.right.Delete(key)
		}

	default:
		// no children
		if n.left == nil && n.right == nil {
			return nil
		}
		// only right
		if n.left == nil {
			return n.right
		}
		// only left
		if n.right == nil {
			return n.left
		}
		// two children
		//find minimum value on the right side
		node := n.right.minimum()
		// and set data to current node
		n.Key = node.Key
		n.Value = node.Value
		// delete node on the right side
		n.right = n.right.Delete(node.Key)
	}
	return n
}

func (n *Node) minimum() *Node {
	if n.left != nil {
		return n.left.minimum()
	}
	return n
}


func (n *Node) String() string {
	if n == nil {
		return "()"
	}
	s := ""
	if n.left != nil {
		s += n.left.String() + " "
	}
	s += fmt.Sprintf("%s:%v", n.Key, n.Value)
	if n.right != nil {
		s += " " + n.right.String()
	}
	return "(" + s + ")"
}
