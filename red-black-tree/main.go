package main

import searchTree "./search_tree"

func main() {
	a := searchTree.CreateSearchTree([]int{})
	a.Insert(5)
	a.Insert(2)
	a.Insert(5)
	a.Insert(22)
	a.Insert(1)
	a.Insert(555)
	a.Insert(3)
	a.Insert(0)
	print(a.GetMin().Value)
}
