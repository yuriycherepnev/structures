package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) PushForward(number int) {
	newNode := &Node{
		Value: number,
		Next:  list.Head,
	}

	list.Head = newNode
}

func main() {
	list := LinkedList{}

	list.PushForward(1)
	list.PushForward(3)
	list.PushForward(5)
	list.PushForward(7)

	current := list.Head

	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
}
