package main

import searchTree "./search_tree"

func main() {
	a := searchTree.CreateSearchTree([]int{})
	a.Insert(5)
	a.Insert(2)
	a.Insert(5)
	print(a.GetMin().Value)
}
