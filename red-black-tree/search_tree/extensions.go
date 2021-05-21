package search_tree

import "container/list"

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

	if node.left == nil {
		return node
	}

	return getMinimum(node.left)
}

func preOrder(node *Node, current *list.List) {
	if node == nil {
		return
	}
	preOrder(node.left, current)
	current.PushBack(node.Value)
	preOrder(node.right, current)
}

func (tree *SearchTree) Values() *list.List {
	result := list.New()
	preOrder(tree.root, result)
	return result
}

func (tree *SearchTree) Find(value int) *Node {
	return getNodeWithSameValue(tree.root, value)
}

func (tree *SearchTree) GetMin() *Node {
	return getMinimum(tree.root)
}
