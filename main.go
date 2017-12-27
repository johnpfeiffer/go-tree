package main

import "fmt"

func main() {
	t := Tree{Root: &Node{Data: 0}}
	fmt.Printf("%#v\n", t)
	t.AddValue(1)
	t.Display()
	t.AddValue(2)
	t.Display()

}
