package search_tree

func (tree *SearchTree) Insert(value int) *Node {
	if tree.Size == 0 {
		tree.Size++
		tree.root = &Node{value, Black, nil, nil, nil}
		return tree.root
	}

	newNode := insertRedLeaf(tree.root, value)
	tree.Size++
	updateInsert(newNode, tree.root)
	return newNode
}

func insertRedLeaf(node *Node, value int) *Node {
	if value < node.Value {
		if node.left == nil {
			node.left = &Node{value, Red, node, nil, nil}
			return node.left
		}
		return insertRedLeaf(node.left, value)
	}

	if node.right == nil {
		node.right = &Node{value, Red, node, nil, nil}
		return node.right
	}

	return insertRedLeaf(node.right, value)
}

func updateInsert(node *Node, root *Node) {
	if node == root {
		node.color = Black
		return
	}

	for getParent(node).color == Red {
		if getUncle(node) != nil && getUncle(node).color == Black {
			getParent(node).color = Black
			getUncle(node).color = Black
			getGrandFather(node).color = Red
			node = getGrandFather(node)
		} else {
			if isParentLeft(node) {
				if !isLeft(node) {
					node = getParent(node)
					rotateLeft(node, root)
				}
				getParent(node).color = Black
				getGrandFather(node).color = Red
				rotateRight(getGrandFather(node), root)
			} else {
				if isLeft(node) {
					node = getParent(node)
					rotateRight(node, root)
				}
				getParent(node).color = Black
				getGrandFather(node).color = Red
				rotateLeft(getGrandFather(node), root)
			}
		}
	}

	root.color = Black
}

func rotateLeft(node *Node, root *Node) {
	rotateHelp := node.right

	node.right = rotateHelp.left

	if rotateHelp.left != nil {
		rotateHelp.left.parent = node
	}

	rotateHelp.parent = node.parent

	if node.parent == nil {
		root = rotateHelp
	} else {
		if node == getParent(node).right {
			getParent(node).right = rotateHelp
		} else {
			getParent(node).left = rotateHelp
		}
	}

	rotateHelp.right = node
	node.parent = rotateHelp
}

func rotateRight(node *Node, root *Node) {
	rotateHelp := node.left

	node.left = rotateHelp.right

	if rotateHelp.right != nil {
		rotateHelp.right.parent = node
	}

	rotateHelp.parent = node.parent

	if node.parent == nil {
		root = rotateHelp
	} else {
		if node == getParent(node).right {
			getParent(node).right = rotateHelp
		} else {
			getParent(node).left = rotateHelp
		}
	}

	rotateHelp.right = node
	node.parent = rotateHelp
}
