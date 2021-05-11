package main

import searchTree "./search_tree"

func main() {
	a := searchTree.CreateSearchTree([]int{})
	a.Insert(5)
	a.Insert(3)
	a.Insert(8)
	print(a.GetMin().Value)
}
