package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	t := Tree{Root: &TreeNode{Data: 0}}
	fmt.Printf("%#v\n", t)
	t.AddValue(1)
	t.Display()
	t.AddValue(2)
	t.Display()

	fmt.Println("Binary Search Tree")
	bst := BinarySearchTree{}
	bst.InsertValue(0)
	fmt.Printf("%#v\n", bst)
	bst.InsertValue(-1)
	bst.Display()
	fmt.Printf("%#v\n", bst.Find(0))
	fmt.Printf("%#v\n", bst.Find(-1))
	fmt.Printf("%#v\n", bst.Find(-2))

	bst.InsertValue(1)
	fmt.Println("traversed in-order", bst.Display())
	fmt.Println("traversed pre-order:", TraversePreOrder(bst.Root))
	fmt.Println("should equal:", sortedIntsString([]int{0, -1, 1}))

	b := []int{2, 0, 1, -1}
	fmt.Println(b)
	bst2 := createBST(b)
	fmt.Println(" bst2", bst2.Root.Data)
	fmt.Println("   bst2", bst2.Root.left.Data)
	fmt.Println("JOHN", TraverseInOrder(bst2.Root))
	bst2.RemoveValue(0)
	fmt.Println("JOHN", TraverseInOrder(bst2.Root))
}

func createBST(a []int) BinarySearchTree {
	bst := BinarySearchTree{}
	for _, v := range a {
		bst.InsertValue(v)
	}
	return bst
}

// SortedIntsString converts a slice of ints to a string, e.g. {1, 2} becomes " 1 2" (does not modify the original slice)
func sortedIntsString(a []int) string {
	var result string
	temp := make([]int, len(a))
	copy(temp, a)
	sort.Ints(temp)
	for _, v := range temp {
		result = result + " " + strconv.Itoa(v)
	}
	return result
}

func intRemoved(target int, a []int) []int {
	var result []int
	for _, v := range a {
		if v != target {
			result = append(result, v)
		}
	}
	return result
}
