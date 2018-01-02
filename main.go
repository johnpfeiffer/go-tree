package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	displayGenericTree()
	fmt.Println("Binary Search Tree")
	bst := BinarySearchTree{}
	bst.InsertValue(0)
	fmt.Printf("%#v\n", bst)
	bst.InsertValue(-1)
	bst.Display()
	fmt.Println("traversed in-order", bst.Display())
	fmt.Println("traversed pre-order", TraversePreOrder(bst.Root))
	fmt.Println("height:", bst.Height()) // should be 1
	fmt.Printf("find 0: %#v\n", bst.Find(0))
	fmt.Printf("find -1: %#v\n", bst.Find(-1))
	fmt.Printf("find -2: %v\n", bst.Find(-2))

	bst.InsertValue(1) // perfect tree
	fmt.Println("traversed in-order", bst.Display())
	fmt.Println("traversed pre-order:", TraversePreOrder(bst.Root))
	fmt.Println("should equal:", sortedIntsString([]int{0, -1, 1}))
	fmt.Println()

	b := []int{2, 0, 1, -1}
	bst2 := createBST(b)
	fmt.Println(b, "(perfect subtree) traversed in-order", TraverseInOrder(bst2.Root))
	fmt.Println("    ", bst2.Root.Data)
	fmt.Println("  ", bst2.Root.left.Data)
	fmt.Println(bst2.Root.left.left.Data, " ", bst2.Root.left.right.Data)

	bst2.RemoveValue(0)
	fmt.Println("Bug: after Removing 0:", TraverseInOrder(bst2.Root))
	fmt.Println()

	c := []int{3, 5, 2, 1, 4, 6, 7}
	bst3 := createBST(c)
	fmt.Println(c, "traversed pre-order:", TraversePreOrder(bst3.Root))
	fmt.Println("height:", bst3.Height()) // should be 3

	fmt.Println("minimum depth:", bst3.MinimumDepth())
}

func displayGenericTree() {
	t := Tree{Root: &TreeNode{Data: 0}}
	fmt.Printf("%#v\n", t)
	t.AddValue(1)
	t.Display()
	t.AddValue(2)
	t.Display()
	fmt.Println()
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
