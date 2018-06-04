package bst

/*

Дерево бинарного поиска

	tree := NewTree(func(f, s interface{}) Comparison {
		intF, _ := f.(string)
		intS, _ := s.(string)
		switch {
		case intF < intS:
			return IsLesser
		case intF > intS:
			return IsGreater
		}
		return AreEqual
	})

	// Empty
	if empty := tree.Empty(); empty {
		// ...
	}

	// Insert
	tree.Insert("A", "AA")
	tree.Insert("L", "LL")
	tree.Insert("G", "GG")
	fmt.Println(tree)

	// Find
	if value, found := tree.Find("L"); found {
		// ...
	}

	// Delete
	tree.Delete("A")
	fmt.Println(tree)


	// Clear
	tree.Clear()


*/
