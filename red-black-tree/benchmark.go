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
	const bigN = 1000000
	tree := searchTree.CreateSearchTree([]int{})
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

	tree := searchTree.CreateSearchTree([]int{})
	sourceArray := list.New()

	for i := 0; i < N; i++ {
		newValue := rand.Int()
		sourceArray.PushBack(newValue)
		tree.Insert(newValue)

		min := math.MaxInt64

		for it := sourceArray.Front(); it != nil; it = it.Next() {
			if min > (*it).Value.(int) {
				min = (*it).Value.(int)
			}
		}

		if tree.GetMin().Value != min {
			fmt.Printf(ErrorColor, "NOT PASSED!\n")
			return false
		}
	}
	fmt.Printf(SuccessColor, fmt.Sprintf("PASSED with %d elements\n", N))
	return true
}

func sortedOrderTest() bool {
	const N = 10
	var sourceArray [N]int

	tree := searchTree.CreateSearchTree([]int{})

	for i := 0; i < N; i++ {
		newValue := rand.Intn(1000)
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
	tree := searchTree.CreateSearchTree([]int{})

	const N = 10

	for i := 0; i < N; i++ {
		newValue := rand.Intn(1000)
		tree.Insert(newValue)
	}

	for it := tree.Begin(); it != nil; it = it.Next() {
		println(it.Node.Value)
	}
	return false
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

	fmt.Printf(InfoColor, "TESTING SUCCESSFULLY COMPLETED")
}

type Function = func() bool

type Test struct {
	function Function
	name     string
}

func main() {
	setupTests([]Test{
		//{hardTest, "HARD"},
		//{minimumFindTest, "MINIMUM"},
		//{sortedOrderTest, "SORTED ORDER"},
		{iterationTest, "ITERATION"}})
}
