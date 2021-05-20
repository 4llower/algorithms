package main

import (
	"fmt"
	"math/rand"
	"time"
)
import "container/list"
import searchTree "./search_tree"

const MAXN = 1000000

type ArrayValue struct {
	Value int
}

func hardTest() bool {
	println("--- HARD TEST START ---")
	tree := searchTree.CreateSearchTree([]int{})
	sourceArray := list.New()

	start := time.Now()

	for i := 0; i < MAXN; i++ {
		newValue := rand.Int()
		sourceArray.PushBack(newValue)
		tree.Insert(newValue)
	}

	for it := sourceArray.Front(); it != nil; it = it.Next() {
		if tree.Find((*it).Value.(int)) == nil {
			println("NOT PASSED!")
			return false
		}
	}

	println(fmt.Sprintf("PASSED at %f with %d elements", time.Now().Sub(start).Seconds(), MAXN))
	println("--- HARD TEST END ---\n")
	return true
}

func main() {
	passedTests := 0
	allTests := 1
	println("START TESTING...\n")

	if hardTest() {
		passedTests++
	}

	println(fmt.Sprintf("--- PASSED %d / %d ---", passedTests, allTests))
}
