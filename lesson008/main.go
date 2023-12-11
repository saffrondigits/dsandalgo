package main

import (
	"fmt"
	"log"
)

type Queue struct {
	items  []int
	length int
}

func (q *Queue) Push(item int) {
	q.items = append(q.items, item)
	q.length++
}

func (q *Queue) Pop() {
	if len(q.items) == 0 {
		log.Println("Queue is empty!")
		return
	}

	q.items = q.items[1:]
	q.length--
}

func (q *Queue) Display() {
	fmt.Printf("The stack has %v items\n", q.length)

	for _, v := range q.items {
		fmt.Printf("%v\t", v)
	}

	fmt.Println()
}

func main() {
	stack := &Queue{}

	stack.Push(5)
	stack.Push(3)
	stack.Push(6)
	stack.Push(9)
	stack.Push(3)

	stack.Display()

	stack.Pop()
	stack.Pop()

	stack.Display()

	stack.Pop()
	stack.Display()

	stack.Pop()
	stack.Pop()
	stack.Display()

	stack.Pop()

}
