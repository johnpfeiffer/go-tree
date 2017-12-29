package main

import "fmt"

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
	fmt.Println("pre-order", bst.Display())
	fmt.Println("in order:")
	DisplayOrdered(bst.Root)
}
