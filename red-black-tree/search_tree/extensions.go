package search_tree

import "container/list"

func getNodeWithSameValue(compare Comparable, node *Node, value interface{}) *Node {
	if node == nil {
		return nil
	}

	if compare(node.Value, value) == 0 {
		return node
	}

	if compare(value, node.Value) < 0 {
		return getNodeWithSameValue(compare, node.left, value)
	}

	return getNodeWithSameValue(compare, node.right, value)
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

func getMaximum(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.right == nil {
		return node
	}

	return getMaximum(node.right)
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

func (tree *SearchTree) Find(value interface{}) *Node {
	return getNodeWithSameValue(tree.Compare, tree.root, value)
}

func (tree *SearchTree) GetMin() *Node {
	return getMinimum(tree.root)
}

func (tree *SearchTree) GetMax() *Node {
	return getMaximum(tree.root)
}
