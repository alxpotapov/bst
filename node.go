package bst

import (
	"fmt"
)

type node struct {
	key      interface{}
	value    interface{}
	children [2]*node
}

func newNode(key, value interface{}) *node {
	return &node{key: key, value: value}
}

func (n *node) insert(key, value interface{}, comparer Comparer) *node {
	// если не найден узел, то возвращаем новый
	if n == nil {
		return newNode(key, value)
	}
	var offset offset
	switch comparer(n.key, key) {
	case IsGreater:
		offset = left
	case IsLesser:
		offset = right
	case AreEqual:
		// либо заменяем value в случае set(map),
		// либо добавляем в список в случае multiset(multimap)
		return n
	}
	n.children[offset] = n.children[offset].insert(key, value, comparer)
	return n
}

func (n *node) find(key interface{}, comparer Comparer) (interface{}, bool) {
	// если не найден узел, то возвращаем новый
	if n == nil {
		return nil, false
	}
	var offset offset
	switch comparer(n.key, key) {
	case IsGreater:
		offset = left
	case IsLesser:
		offset = right
	case AreEqual:
		return n.value, true
	}
	return n.children[offset].find(key, comparer)
}

func (n *node) delete(key interface{}, comparer Comparer) *node {
	// если не найден узел, то возвращаем nil
	if n == nil {
		return nil
	}
	var offset offset
	switch comparer(n.key, key) {
	case IsGreater:
		offset = left
	case IsLesser:
		offset = right
	case AreEqual:
		return n.splice(comparer)
	}
	n.children[offset] = n.children[offset].delete(key, comparer)
	return n
}

// String - вывод на экран
func (n *node) String() string {
	if n == nil {
		return ""
	}
	s := ""
	s += n.children[left].String()
	s += fmt.Sprintf("%v", n.key)
	s += n.children[right].String()
	return "(" + s + ")"
}

func (n *node) splice(comparer Comparer) *node {
	//Удалить узел и вернуть nil
	if n.children[left] == nil && n.children[right] == nil {
		return nil
	}
	//Удалить узел и вернуть его левую подветвь
	if n.children[right] == nil {
		return n.children[left]
	}
	//Удалить узел и вернуть его правую подветвь
	if n.children[left] == nil {
		return n.children[right]
	}
	//Заменить значение текущего узла на максимум левой подветви
	//Удалить максимум левой подветви
	//Вернуть собранное значение
	tempNode := n.children[left].findMax()
	n.key = tempNode.key
	n.value = tempNode.value
	n.children[left] = n.children[left].delete(n.key, comparer)
	return n
}

// findMax - вернуть узел с максимальным значением из левой подветви
func (n *node) findMax() *node {
	if n.children[right] != nil {
		n.children[right] = n.children[right].findMax()
	}
	return n
}
