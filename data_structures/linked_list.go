package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (List *LinkedList) PushFront(num int) {
	newNode := &Node{
		Value: num,
		Next:  List.Head,
	}
	List.Head = newNode
}

func main() {
	list := LinkedList{}

	list.PushFront(19)
	list.PushFront(20)
	list.PushFront(21)
	list.PushFront(22)

	current := list.Head

	for current != nil {
		fmt.Println(current)
		current = current.Next
	}

}
