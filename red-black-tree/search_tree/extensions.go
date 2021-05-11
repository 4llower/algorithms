package search_tree

func getNodeWithSameValue(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if node.Value == value {
		return node
	}

	if node.Value > value {
		return getNodeWithSameValue(node.left, value)
	}

	return getNodeWithSameValue(node.right, value)
}

func getMinimum(node *Node) *Node {
	if node == nil {
		return nil
	}

	println(node.left, node.right)

	if node.left == nil {
		return node
	}

	return getMinimum(node.left)
}

func (tree *SearchTree) Find(value int) *Node {
	return getNodeWithSameValue(tree.root, value)
}

func (tree *SearchTree) GetMin() *Node {
	return getMinimum(tree.root)
}
