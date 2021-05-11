package main

import searchTree "./search_tree"

func main() {
	a := searchTree.CreateSearchTree([]int{})
	a.Insert(5)
	a.Insert(10)
	a.Insert(15)
	a.Insert(25)
	print(a.Size)
}
