package search_tree

type Iterator struct {
	Tree *SearchTree
	Node *Node
}

func getNext(node *Node) *Node {
	if node.right != nil {
		return getMinimum(node.right)
	}
	current := getParent(node)
	for current != nil && node == current.right {
		node = current
		current = getParent(current)
	}
	return current
}

func getPrev(node *Node) *Node {
	if node.left != nil {
		return getMaximum(node.left)
	}

	current := getParent(node)

	for current != nil && node == current.left {
		node = current
		current = getParent(current)
	}

	return current
}

func (tree *SearchTree) Begin() *Iterator {
	return &Iterator{tree, getMinimum(tree.root)}
}

func (tree *SearchTree) End() *Iterator {
	return &Iterator{tree, getMaximum(tree.root)}
}

func (iterator *Iterator) Next() *Iterator {
	nextNode := getNext(iterator.Node)

	if nextNode == nil {
		return nil
	}

	return &Iterator{iterator.Tree, nextNode}
}

func (iterator *Iterator) Prev() *Iterator {
	prevNode := getPrev(iterator.Node)

	if prevNode == nil {
		return nil
	}

	return &Iterator{iterator.Tree, prevNode}
}
