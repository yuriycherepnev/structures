/*
Связный список
Элементы разбросаны по памяти и связаны указателями:

CPU очень любит последовательную память.
Когда данные лежат подряд:
процессор читает их пачками
кэш CPU работает эффективно
меньше cache miss
Поэтому slice почти всегда быстрее linked list.
*/

package main

import "fmt"

type DoubleNode struct {
	Value int
	Next  *DoubleNode
	Prev  *DoubleNode
}

type DoubleLinkedList struct {
	Head *DoubleNode
	Tail *DoubleNode
}

func (list *DoubleLinkedList) PushBack(value int) {
	newNode := &DoubleNode{
		Value: value,
	}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		return
	}
	newNode.Next = list.Tail
	list.Tail.Prev = newNode
	list.Tail = newNode
}

func (list *DoubleLinkedList) PushFront(value int) {
	newNode := &DoubleNode{
		Value: value,
	}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		return
	}
	newNode.Prev = list.Head
	list.Head.Next = newNode
	list.Head = newNode
}

func main() {
	list := DoubleLinkedList{}

	list.PushBack(40)
	list.PushBack(30)
	list.PushBack(20)
	list.PushBack(10)
	list.PushFront(50)
	list.PushFront(60)
	list.PushFront(70)
	list.PushFront(80)
	list.PushFront(90)

	current := list.Head

	for current != nil {
		fmt.Printf("%+v | %p\n", current, current)
		current = current.Prev
	}
}
