package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) pushForward(number int) {
	newNode := &Node{
		Value: number,
	}
	if list.Head == nil {
		list.Head = newNode
		return
	}
	newNode.Next = list.Head
	list.Head = newNode
}

func main() {
	list := LinkedList{}

	list.pushForward(1)
	list.pushForward(3)
	list.pushForward(5)
	list.pushForward(7)

	current := list.Head

	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
}
