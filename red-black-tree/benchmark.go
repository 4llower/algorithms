package main

import "math/rand"
import "container/list"
import searchTree "./search_tree"

const MAXN = 1000000

type ArrayValue struct {
	Value int
}

func main() {
	tree := searchTree.CreateSearchTree([]int{})
	sourceArray := list.New()

	for i := 0; i < MAXN; i++ {
		newValue := rand.Int()
		sourceArray.PushBack(newValue)
		tree.Insert(newValue)
		//println(newValue)
	}

	//println("---- end ----")

	for it := sourceArray.Front(); it != nil; it = it.Next() {
		//println((*it).Value.(int))
		if tree.Find((*it).Value.(int)) == nil {
			print("NOT PASSED!")
			return
		}
	}

	print("PASSED")
}
