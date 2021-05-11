package search_tree

func getParent(node *Node) *Node {
	return node.ancestor
}

func getGrandFather(node *Node) *Node {
	return getParent(node.ancestor)
}

func getUncle(node *Node) *Node {
	if isParentLeft(node) {
		return getGrandFather(node).right
	} else {
		return getGrandFather(node).left
	}
}

func isParentLeft(node *Node) bool {
	if getGrandFather(node).left == getParent(node) {
		return true
	}
	return false
}

func isLeft(node *Node) bool {
	if getParent(node).left == node {
		return true
	}
	return false
}
