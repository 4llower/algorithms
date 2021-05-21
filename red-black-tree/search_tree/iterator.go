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
	return &Iterator{tree, nil}
}

func (iterator *Iterator) Next() *Node {
	if iterator.Node == nil {
		return getMinimum(iterator.Tree.root)
	}
	return getNext(iterator.Node)
}

func (iterator *Iterator) Prev() {

}
