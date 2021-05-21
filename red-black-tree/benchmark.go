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

type ArrayValue struct {
	Value int
}

func hardTest() bool {
	const bigN = 1000000
	println("--- HARD TEST START ---")
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
			println("--- HARD TEST END ---\n")
			return false
		}
	}

	fmt.Printf(SuccessColor, fmt.Sprintf("PASSED at %f with %d elements\n", time.Now().Sub(start).Seconds(), bigN))
	println("--- HARD TEST END ---\n")
	return true
}

func minimumFindTest() bool {
	const N = 100
	println("--- MINIMUM TEST START ---")

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
			println("--- MINIMUM TEST END ---\n")
			return false
		}
	}
	fmt.Printf(SuccessColor, fmt.Sprintf("PASSED with %d elements\n", N))
	println("--- MINIMUM TEST END ---\n")
	return true
}

func sortedOrderTest() bool {
	const N = 10
	var sourceArray [N]int
	println("--- SORTED ORDER START ---")

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

	treeList := tree.ToSortedList()

	for it := treeList.Front(); it != nil; it = it.Next() {
		if (*it).Value.(int) != slice[index] {
			fmt.Printf(ErrorColor, "NOT PASSED!\n")
			println("--- SORTED ORDER END ---\n")
			return false
		}
		index++
	}

	fmt.Printf(SuccessColor, fmt.Sprintf("PASSED with %d elements\n", N))
	println("--- SORTED ORDER END ---\n")
	return true
}

func main() {
	passedTests := 0
	allTests := 3

	fmt.Printf(InfoColor, "START TESTING...\n\n")

	if hardTest() {
		passedTests++
	}

	if minimumFindTest() {
		passedTests++
	}

	if sortedOrderTest() {
		passedTests++
	}

	if passedTests == allTests {
		fmt.Printf(SuccessColor, fmt.Sprintf("--- PASSED %d / %d ---", passedTests, allTests))
	} else {
		fmt.Printf(WarningColor, fmt.Sprintf("--- PASSED %d / %d ---", passedTests, allTests))
	}
}
