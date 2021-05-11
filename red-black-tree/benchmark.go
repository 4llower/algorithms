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
		sourceArray.PushBack(ArrayValue{Value: newValue})
		tree.Insert(newValue)
	}

	for it := sourceArray.Front(); it != nil; it = it.Next() {
		if tree.Find(it.Value.(*ArrayValue).Value) == nil {
			print("NOT PASSED!")
			return
		}
	}

	print("PASSED")
}
