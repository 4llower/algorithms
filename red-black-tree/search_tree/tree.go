package search_tree

type SearchTree struct {
	Size int
	root *Node
}

func CreateSearchTree(base []int) *SearchTree {
	tree := &SearchTree{}

	for _, value := range base {
		tree.Insert(value)
	}

	return tree
}
