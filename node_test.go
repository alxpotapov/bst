package bst

import (
	"testing"
)

var n *node
var comparer = func(f, s interface{}) Comparison {
	intF, ok := f.(string)
	if !ok {
		return IsLesser
	}
	intS, ok := s.(string)
	if !ok {
		return IsLesser
	}
	switch {
	case intF < intS:
		return IsLesser
	case intF > intS:
		return IsGreater
	}
	return AreEqual
}

func TestInsert(t *testing.T) {

	n = n.insert("A", nil, comparer)
	t.Log(n)
	n = n.insert("L", nil, comparer)
	t.Log(n)
	n = n.insert("G", nil, comparer)
	t.Log(n)
	n = n.insert("O", nil, comparer)
	t.Log(n)
	n = n.insert("R", nil, comparer)
	t.Log(n)
	n = n.insert("I", nil, comparer)
	t.Log(n)
	n = n.insert("T", nil, comparer)
	t.Log(n)
	n = n.insert("H", nil, comparer)
	t.Log(n)
	n = n.insert("M", nil, comparer)
	t.Log(n)
}

func TestFind(t *testing.T) {
	t.Log("Find A")
	if _, found := n.find("A", comparer); !found {
		t.Fatal("expected found")
	}
	t.Log("Find X")
	if _, found := n.find("X", comparer); found {
		t.Fatal("expected not found")
	}
	t.Log("Find L")
	if _, found := n.find("L", comparer); !found {
		t.Fatal("expected found")
	}
	t.Log("Find M")
	if _, found := n.find("M", comparer); !found {
		t.Fatal("expected found")
	}
	t.Log("Find W")
	if _, found := n.find("W", comparer); found {
		t.Fatal("expected not found")
	}

}

func TestDelete(t *testing.T) {
	t.Log("Delete A")
	n = n.delete("A", comparer)
	t.Log(n)
	t.Log("Delete L")
	n = n.delete("L", comparer)
	t.Log(n)
	t.Log("Delete G")
	n = n.delete("G", comparer)
	t.Log(n)
	t.Log("Delete O")
	n = n.delete("O", comparer)
	t.Log(n)
	t.Log("Delete R")
	n = n.delete("R", comparer)
	t.Log(n)
	t.Log("Delete I")
	n = n.delete("I", comparer)
	t.Log(n)
	t.Log("Delete T")
	n = n.delete("T", comparer)
	t.Log(n)
	t.Log("Delete H")
	n = n.delete("H", comparer)
	t.Log(n)
	t.Log("Delete M")
	n = n.delete("M", comparer)
	t.Log(n)
}
