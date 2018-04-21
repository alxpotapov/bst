// Package bst contain a implementation of binary search tree
package bst

import (
	"strings"
	"fmt"
)

// Node contains the key as string, value as interface{} and 2 pointers to his children
type Node struct {
	left, right *Node
	key         string
	value       interface{}
}

// Insert a key and value into a sorted tree
func (n *Node) Insert(key string, value interface{}) {
	comp := strings.Compare(n.key, key)
	switch {
	case comp > 0:
		if n.left != nil {
			n.left.Insert(key, value)
		} else {
			n.left = &Node{key: key, value: value}
		}
	case comp < 0:
		if n.right != nil {
			n.right.Insert(key, value)
		} else {
			n.right = &Node{key: key, value: value}
		}
	default:
		n.value = value
	}
}

// Find value by key
func (n *Node) Find(key string)(interface{},  bool) {
	comp := strings.Compare(n.key, key)
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
		return n.value, true
	}
	return nil, false
}

// Delete a value from a sorted tree by key
func (n *Node) Delete(key string) *Node {
	comp := strings.Compare(n.key, key)
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
		// change data between nodes
		n.key = node.key
		n.value = node.value
		n.right = n.right.Delete(node.key)
	}
	return n
}

func (n *Node) minimum() *Node {
	if n.left != nil {
		return n.left.minimum()
	}
	return n
}

func (n *Node) inOrder() string {
	if n == nil {
		return "()"
	}
	s := ""
	if n.left != nil {
		s += n.left.inOrder() + " "
	}
	s += fmt.Sprintf("%s:%v", n.key, n.value)
	if n.right != nil {
		s += " " + n.right.inOrder()
	}
	return "(" + s + ")"
}
