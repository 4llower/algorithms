package search_tree

type Comparable = func(value1 interface{}, value2 interface{}) int

type SearchTree struct {
	Size    int
	root    *Node
	Compare Comparable
}

func CreateSearchTree(base []interface{}) *SearchTree {
	tree := &SearchTree{0, nil, defaultCompare}

	for _, value := range base {
		tree.Insert(value)
	}

	return tree
}

func CreateSearchTreeComparable(base []interface{}, compare Comparable) *SearchTree {
	tree := &SearchTree{0, nil, compare}

	for _, value := range base {
		tree.Insert(value)
	}

	return tree
}
