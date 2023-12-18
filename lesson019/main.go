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
	list.Display()

}
