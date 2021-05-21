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

func (tree *SearchTree) Begin() *Iterator {
	return &Iterator{tree, getMinimum(tree.root)}
}

func (iterator *Iterator) Next() *Iterator {
	nextNode := getNext(iterator.Node)

	if nextNode == nil {
		return nil
	}

	return &Iterator{iterator.Tree, nextNode}
}

func (iterator *Iterator) Prev() {

}
