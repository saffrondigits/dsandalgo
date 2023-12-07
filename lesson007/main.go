package main

import (
	"fmt"
	"log"
)

type Stack struct {
	items  []int
	length int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
	s.length++
}

func (s *Stack) Pop() {
	if len(s.items) == 0 {
		log.Println("Stack is empty!")
		return
	}

	s.items = s.items[:len(s.items)-1]
	s.length--
}

func (s *Stack) Display() {
	fmt.Printf("The stack has %v items\n", s.length)

	for _, v := range s.items {
		fmt.Printf("%v\t", v)
	}

	fmt.Println()
}

func main() {
	stack := &Stack{}

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
