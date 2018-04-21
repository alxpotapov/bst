package bst

import (
	"testing"
)

func TestNode_Insert(t *testing.T) {
	tree := &Node{key: "5", value: "5"}

	tree.Insert("2", "2")
	tree.Insert("1", "1")
	tree.Insert("7", "7")
	tree.Insert("6", "6")
	tree.Insert("8", "8")
	tree.Insert("3", "3")
	tree.Insert("3", "33")

	if tree.inOrder() != "(((1:1) 2:2 (3:33)) 5:5 ((6:6) 7:7 (8:8)))" {
		t.Error("expected (((1:1) 2:2 (3:33)) 5:5 ((6:6) 7:7 (8:8))), got ", tree.inOrder())
	}
}

func TestNode_Find(t *testing.T) {
	tree := &Node{key: "5", value: "5"}

	tree.Insert("2", "2")
	tree.Insert("1", "1")
	tree.Insert("7", "7")
	tree.Insert("6", "6")
	tree.Insert("8", "8")
	tree.Insert("3", "3")

	key := "5"
	if value, ok :=tree.Find(key); !ok {
		t.Errorf("expected found %s:%v", key, value)
	}

	key = "1"
	if value, ok :=tree.Find(key); !ok {
		t.Errorf("expected found %s:%v", key, value)
	}

	key = "8"
	if value, ok :=tree.Find(key); !ok {
		t.Errorf("expected found %s:%v", key, value)
	}

	key = "2"
	if value, ok :=tree.Find(key); !ok {
		t.Errorf("expected found %s:%v", key, value)
	}

	key = "20"
	if value, ok :=tree.Find(key); ok {
		t.Errorf("expected not found %s:%v", key, value)
	}

}

func TestNode_Delete(t *testing.T) {
	tree := &Node{key: "5", value: "5"}

	tree.Insert("2", "2")
	tree.Insert("1", "1")
	tree.Insert("7", "7")
	tree.Insert("6", "6")
	tree.Insert("8", "8")
	tree.Insert("3", "3")

	key := "2"
	tree = tree.Delete(key)
	if tree.inOrder() != "(((1:1) 3:3) 5:5 ((6:6) 7:7 (8:8)))" {
		t.Error("delete ", key)
		t.Error("expected (((1:1) 3:3) 5:5 ((6:6) 7:7 (8:8))), got ", tree.inOrder())
	}

	key = "5"
	tree = tree.Delete(key)
	if tree.inOrder() != "(((1:1) 3:3) 6:6 (7:7 (8:8)))" {
		t.Error("delete ", key)
		t.Error("expected (((1:1) 3:3) 6:6 (7:7 (8:8))), got ", tree.inOrder())
	}

	key = "7"
	tree = tree.Delete(key)
	if tree.inOrder() != "(((1:1) 3:3) 6:6 (8:8))" {
		t.Error("delete ", key)
		t.Error("expected (((1:1) 3:3) 6:6 (8:8)), got ", tree.inOrder())
	}

	key = "10"
	tree = tree.Delete(key)
	if tree.inOrder() != "(((1:1) 3:3) 6:6 (8:8))" {
		t.Error("delete ", key)
		t.Error("expected (((1:1) 3:3) 6:6 (8:8)), got ", tree.inOrder())
	}

	key = "3"
	tree.Delete(key)
	if tree.inOrder() != "((1:1) 6:6 (8:8))" {
		t.Error("delete ", key)
		t.Error("expected ((1:1) 6:6 (8:8)), got ", tree.inOrder())
	}

	key = "1"
	tree = tree.Delete(key)
	if tree.inOrder() != "(6:6 (8:8))" {
		t.Error("delete ", key)
		t.Error("expected (6:6 (8:8)), got ", tree.inOrder())
	}

	key = "8"
	tree = tree.Delete(key)
	if tree.inOrder() != "(6:6)" {
		t.Error("delete ", key)
		t.Error("expected (6:6), got ", tree.inOrder())
	}

	key = "6"
	tree = tree.Delete(key)
	if tree.inOrder() != "()" {
		t.Error("delete ", key)
		t.Error("expected (), got ", tree.inOrder())
	}
}

func TestNode_InOrder(t *testing.T) {
	tree := &Node{key: "5", value: "5"}

	tree.Insert("2", "2")
	tree.Insert("1", "1")
	tree.Insert("7", "7")
	tree.Insert("6", "6")
	tree.Insert("8", "8")
	tree.Insert("3", "3")

	tree = tree.Delete("2")
	tree = tree.Delete("5")
	tree = tree.Delete("7")
	tree = tree.Delete("8")

	if tree.inOrder() != "(((1:1) 3:3) 6:6)" {
		t.Error("expected (((1:1) 3:3) 6:6), got ", tree.inOrder())
	}
}

