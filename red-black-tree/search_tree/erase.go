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
	var current *Node
	for node != tree.root && node.color == Black {
		if isLeft(node) {
			current = getParent(node).right
			if current.color == Red {
				current.color = Black
				node.parent.color = Red
				rotateLeft(getParent(node), tree)
				current = getParent(node).right
			}
			if current.left.color == Black && current.right.color == Black {
				current.color = Red
				node = getParent(node)
			} else {
				if current.right.color == Black {
					current.left.color = Black
					current.color = Red
					rotateRight(node, tree)
					current = node.parent.right
				}
				current.color = getParent(node).color
				node.parent.color = Black
				current.right.color = Black
				rotateLeft(getParent(node), tree)
				node = tree.root
			}
		} else {
			current = getParent(node).left
			if current.color == Red {
				current.color = Black
				node.parent.color = Red
				rotateRight(getParent(node), tree)
				current = getParent(node).left
			}
			if current.right.color == Black && current.left.color == Black {
				current.color = Red
				node = getParent(node)
			} else {
				if current.left.color == Black {
					current.right.color = Black
					current.color = Red
					rotateLeft(node, tree)
					current = getParent(node).left
				}
				current.color = getParent(node).color
				node.parent.color = Black
				current.left.color = Black
				rotateRight(getParent(node), tree)
				node = tree.root
			}
		}
	}
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
