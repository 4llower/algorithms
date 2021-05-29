package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)
import "container/list"
import searchTree "./search_tree"

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	SuccessColor = "\033[1;32m%s\033[0m"
)

func hardTest() bool {
	const bigN = 500000
	tree := searchTree.CreateSearchTree([]interface{}{})
	sourceArray := list.New()

	start := time.Now()

	for i := 0; i < bigN; i++ {
		newValue := rand.Int()
		sourceArray.PushBack(newValue)
		tree.Insert(newValue)
	}

	for it := sourceArray.Front(); it != nil; it = it.Next() {
		if tree.Find((*it).Value.(int)) == nil {
			fmt.Printf(ErrorColor, "NOT PASSED!\n")
			return false
		}
	}

	fmt.Printf(SuccessColor, fmt.Sprintf("PASSED at %f with %d elements\n", time.Now().Sub(start).Seconds(), bigN))
	return true
}

func minimumFindTest() bool {
	const N = 100

	tree := searchTree.CreateSearchTree([]interface{}{})
	sourceArray := list.New()

	for i := 0; i < N; i++ {
		newValue := rand.Int()
		sourceArray.PushBack(newValue)
		tree.Insert(newValue)

		min := math.MaxInt64
		max := -math.MaxInt64

		for it := sourceArray.Front(); it != nil; it = it.Next() {
			if min > (*it).Value.(int) {
				min = (*it).Value.(int)
			}
			if max < (*it).Value.(int) {
				max = (*it).Value.(int)
			}
		}

		if tree.GetMin().Value != min || tree.GetMax().Value != max {
			fmt.Printf(ErrorColor, "NOT PASSED!\n")
			return false
		}
	}
	fmt.Printf(SuccessColor, fmt.Sprintf("PASSED with %d elements\n", N))
	return true
}

func sortedOrderTest() bool {
	const N = 10

	tree := searchTree.CreateSearchTree([]interface{}{})

	var sourceArray [N]int
	for i := 0; i < N; i++ {
		newValue := rand.Int()
		sourceArray[i] = newValue
		tree.Insert(newValue)
	}

	slice := sourceArray[:]

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	index := 0

	treeList := tree.Values()

	for it := treeList.Front(); it != nil; it = it.Next() {
		if (*it).Value.(int) != slice[index] {
			fmt.Printf(ErrorColor, "NOT PASSED!\n")
			return false
		}
		index++
	}

	fmt.Printf(SuccessColor, fmt.Sprintf("PASSED with %d elements\n", N))
	return true
}

func iterationTest() bool {
	tree := searchTree.CreateSearchTree([]interface{}{})

	const N = 100

	var sourceArray [N]int
	for i := 0; i < N; i++ {
		newValue := rand.Int()
		sourceArray[i] = newValue
		tree.Insert(newValue)
	}

	slice := sourceArray[:]

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	index := 0

	for it := tree.Begin(); it != nil; it = it.Next() {
		if sourceArray[index] != it.Node.Value {
			fmt.Printf(ErrorColor, "NOT PASSED IN ASC ORDER!\n")
			return false
		}
		index++
	}

	for it := tree.End(); it != nil; it = it.Prev() {
		index--
		if sourceArray[index] != it.Node.Value {
			fmt.Printf(ErrorColor, "NOT PASSED IN DESC ORDER!\n")
			return false
		}
	}

	fmt.Printf(SuccessColor, fmt.Sprintf("PASSED with %d elements\n", N))
	return true
}

func setupTests(tests []Test) {
	passedTests := 0
	allTests := len(tests)
	fmt.Printf(InfoColor, "START TESTING...\n\n")

	for i := 0; i < allTests; i++ {
		print(fmt.Sprintf("--- %s TEST START ---\n", tests[i].name))
		if tests[i].function() {
			passedTests++
		}
		println(fmt.Sprintf("--- %s TEST END ---\n", tests[i].name))
	}

	if passedTests == allTests {
		fmt.Printf(SuccessColor, fmt.Sprintf("--- PASSED %d / %d ---\n\n", passedTests, allTests))
	} else {
		fmt.Printf(WarningColor, fmt.Sprintf("--- PASSED %d / %d ---\n\n", passedTests, allTests))
	}

	fmt.Printf(InfoColor, "SUCCESSFULLY COMPLETED")
}

type BoolFunction = func() bool

type Test struct {
	function BoolFunction
	name     string
}

func main() {
	setupTests([]Test{
		{minimumFindTest, "MAXIMUM / MINIMUM"},
		{sortedOrderTest, "SORTED ORDER"},
		{iterationTest, "ITERATION"},
		{hardTest, "HARD"}})
}
