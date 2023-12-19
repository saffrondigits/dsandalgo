package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) InsertEnd(data int) {
	// Creates a new node with provide data
	node := &Node{Data: data}

	// Checks if the head of the LL is nil
	if l.Head == nil {
		l.Head = node
	} else {
		current := l.Head

		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

func (l *LinkedList) InsertAtBeginning(data int) {
	node := &Node{Data: data}

	if l.Head != nil {
		node.Next = l.Head
	}

	l.Head = node
}

func (l *LinkedList) InsertNodeAtTheSpecifiedLocation(data int, position int) {
	node := &Node{Data: data}

	if position == 1 {
		node.Next = l.Head
		l.Head = node
	} else {
		current := l.Head

		// i = 1, position-1 = 2
		for i := 1; i < position-1 && current.Next != nil; i++ {
			current = current.Next
		}

		node.Next = current.Next
		current.Next = node
	}
}

func (l *LinkedList) Display() {
	current := l.Head

	for current != nil {
		fmt.Printf("%d and %v, \n", current.Data, current.Next)
		current = current.Next
	}

	fmt.Println()
}

func main() {
	list := &LinkedList{}

	list.InsertEnd(1)
	list.InsertEnd(2)
	list.InsertEnd(3)
	list.InsertAtBeginning(5)
	list.InsertNodeAtTheSpecifiedLocation(9, 3)
	list.Display()

}
