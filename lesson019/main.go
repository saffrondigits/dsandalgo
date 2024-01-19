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

		// i = 2, position-1 = 2
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
		fmt.Printf("%d, \t", current.Data)
		current = current.Next
	}

	fmt.Println()
}

func (l *LinkedList) DeleteFirstNode() {
	if l.Head != nil {
		l.Head = l.Head.Next
	}
}

func (l *LinkedList) DeleteLastNode() {
	if l.Head != nil {
		if l.Head.Next == nil {
			l.Head = nil
		} else {
			current := l.Head
			for current.Next.Next != nil {
				current = current.Next
			}

			current.Next = nil
		}
	}
}

func (l *LinkedList) DeleteFromSpecifiedLocation(location int) {
	if l.Head == nil {
		return
	}

	temp := l.Head

	if location == 1 {
		l.Head = temp.Next
		return
	}

	for i := 1; temp != nil && i < location-1; i++ {
		temp = temp.Next
	}

	if temp == nil || temp.Next == nil {
		return
	}

	next := temp.Next.Next

	temp.Next = next
}

func main() {
	var choice int
	var nodeValue int
	var nodeLocation int
	list := &LinkedList{}

	for {
		fmt.Println("\nEnter your choice!")
		fmt.Println("_______________________________________________________________")
		fmt.Printf("1. Insert at the beginning\t2. Insert at the end\n")
		fmt.Printf("3. Insert at the specific location\t 4. Display")
		fmt.Println()
		fmt.Printf("5. Delete the first node\t6. Delete the last node\n")
		fmt.Printf("7. Delete specific node \t8. Press -1 to Exit\n")

		fmt.Scanf("%d", &choice)

		switch choice {
		case -1:
			return
		case 1:
			fmt.Println("Enter a node value: ")
			fmt.Scanf("%d", &nodeValue)
			list.InsertAtBeginning(nodeValue)
		case 2:
			fmt.Println("Enter a node value: ")
			fmt.Scanf("%d", &nodeValue)
			list.InsertEnd(nodeValue)
		case 3:
			fmt.Println("Enter a node value: ")
			fmt.Scanf("%d", &nodeValue)
			fmt.Println("Enter a node location: ")
			fmt.Scanf("%d", &nodeLocation)
			list.InsertNodeAtTheSpecifiedLocation(nodeValue, nodeLocation)
		case 4:
			list.Display()
		case 5:
			list.DeleteFirstNode()
		case 6:
			list.DeleteLastNode()
		case 7:
			fmt.Println("Enter a node location: ")
			fmt.Scanf("%d", &nodeLocation)
			list.DeleteFromSpecifiedLocation(nodeLocation)
		case 8:
			return
		default:
			fmt.Println("Invalid input!")
		}
	}

	// list.InsertEnd(1)
	// list.InsertEnd(2)
	// list.InsertEnd(3)
	// list.InsertAtBeginning(5)
	// list.InsertNodeAtTheSpecifiedLocation(9, 3)
	// list.DeleteFirstNode()
	// list.DeleteFromSpecifiedLocation(3)
	// list.Display()

}

func MinWindowSubstring(strArr []string) string {

	// Code goes here
	return strArr[0]
}
