package main

import (
	"fmt"
)

type DLL struct {
	root *Node
	last *Node
}

type Node struct {
	left  *Node
	right *Node
	value int
}

func (dll *DLL) Inset(value int) *Node {
	newNode, root := &Node{}, dll.root

	if dll.root != nil {
		newNode.right = root
		root.left = newNode
	} else { //if list empty then last element is first element
		dll.last = newNode
	}

	newNode.value = value
	dll.root = newNode

	return newNode
}

func (dll *DLL) Remove() {
	if dll.last == nil {
		return
	}

	if dll.last.left == nil {
		dll.last = nil
		dll.root = nil
		return
	}

	dll.last = dll.last.left
	dll.last.right = nil
}

func (dll *DLL) Read(node *Node) {
	if node == dll.root {
		return
	}

	node.left.right = node.right
	node.right.left = node.left

	node.right = dll.root
	node.left = nil

	dll.root.left = node

	dll.root = node
}

func (dll *DLL) ReadAll() {
	for root := dll.root; root != nil; root = root.right {
		fmt.Println(root.value)
	}
}

func (dll *DLL) ReadBack() {
	for current := dll.last; current != nil; current = current.left {
		fmt.Println(current.value)
	}
}

func main() {
	f := new(DLL)

	node := &Node{}

	for index := range 100 {
		g := f.Inset(index)

		if index == 11 {
			node = g
		}

		if index == 45 {
			f.Read(node)
		}
	}

	for range 95 {
		f.Remove()
	}

	f.ReadAll()
	f.ReadBack()
}
