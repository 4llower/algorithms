package search_tree

type Iterator struct {
	tree *SearchTree
	node *Node
}

func (tree *SearchTree) Begin() *Iterator {
	return &Iterator{tree, getMinimum(tree.root)}
}

func (iterator *Iterator) Next() {

}

func (iterator *Iterator) Prev() {

}
