package search_tree

func (tree *SearchTree) Erase(value interface{}) {
	foundNode := getNodeWithSameValue(tree.Compare, tree.root, value)
	if foundNode == nil {
		return
	}
	commitTransplant(tree, foundNode)
}

func commitTransplant(tree *SearchTree, node *Node) {
	current := node
	var help *Node
	if node.left == nil {
		help = node.right
		transplant(tree, node, help)
	} else if node.right == nil {
		help = node.left
		transplant(tree, node, help)
	} else {
		current = getMinimum(node.right)
		help = current.right
		if getParent(current) == node {
			help.parent = current
		} else {
			transplant(tree, current, current.right)
			current.right = node.right
			current.right.parent = current
		}
		transplant(tree, node, current)
		current.left = node.left
		current.left.parent = current
		current.color = node.color
	}
	if current.color == Black {
		updateErase(tree, help)
	}
}

func updateErase(tree *SearchTree, node *Node) {

}

func transplant(tree *SearchTree, node *Node, graft *Node) {
	if getParent(graft) == nil {
		tree.root = graft
	} else if isParentLeft(node) {
		node.parent.left = graft
	} else {
		node.parent.right = graft
	}
	graft.parent = node.parent
}
