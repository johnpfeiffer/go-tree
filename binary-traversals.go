package gotree

import (
	"bytes"
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

// TraversePreOrderRecursive shows the node data (in pre-order) and continues recursively https://en.wikipedia.org/wiki/Tree_traversal#Pre-order
func TraversePreOrderRecursive(n *Node) string {
	var s string
	var b bytes.Buffer
	if n == nil {
		return ""
	}
	b.WriteString(fmt.Sprintf("%d ", n.Data))
	if n.Left != nil {
		s += TraversePreOrderRecursive(n.Left)
	}
	if n.Right != nil {
		s += TraversePreOrderRecursive(n.Right)
	}
	return b.String() + s
}

// TraversePreOrder shows the node data (in pre-order) and continues iteratively
func TraversePreOrder(n *Node) string {
	var b bytes.Buffer
	if n == nil {
		return ""
	}
	stack := []*Node{n}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		b.WriteString(fmt.Sprintf("%d ", current.Data))
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}
	return b.String()
}

// TraverseInOrderRecursive shows the node data ("in-order"), in a BST this ouputs the data in sorted order
func TraverseInOrderRecursive(n *Node) string {
	var s string
	if n == nil {
		return ""
	}
	if n.Left != nil {
		s = s + " " + TraverseInOrderRecursive(n.Left)
	}
	s = s + " " + strconv.Itoa(n.Data)
	if n.Right != nil {
		s = s + " " + TraverseInOrder(n.Right)
	}
	return strings.TrimSpace(s)
}

// TraverseInOrder iteratively shows the Left most nodes, then the parent, then the Right nodes (in a BST this ouputs the data in sorted order)
func TraverseInOrder(root *Node) string {
	var s string
	if root == nil {
		return ""
	}
	stack := []*Node{}
	stack = append(stack, root)
	current := root.Left
	for len(stack) > 0 || current != nil {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1] // pop
		stack = stack[:len(stack)-1]
		s = s + " " + strconv.Itoa(current.Data)
		current = current.Right
	}
	return strings.TrimSpace(s)
}

// TraverseLevelOrder is a breadth first traversal https://en.wikipedia.org/wiki/Tree_traversal#Breadth-first_search
func TraverseLevelOrder(n *Node) string {
	var s string
	if n == nil {
		return ""
	}
	q := list.New()
	q.PushBack(n)
	for q.Len() > 0 {
		temp := q.Front()
		q.Remove(temp)
		current := temp.Value.(*Node)
		s = s + strconv.Itoa(current.Data) + " "
		if current.Left != nil {
			q.PushBack(current.Left)
		}
		if current.Right != nil {
			q.PushBack(current.Right)
		}
	}
	return strings.TrimSpace(s)
}
