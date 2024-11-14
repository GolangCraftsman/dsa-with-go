package main

import "fmt"

type Stack struct {
	Items []int
	Top   int
}

func main() {
	stack := &Stack{Top: -1}

	stack.isEmpty()

	stack.Push(5)
	stack.Push(10)
	stack.Push(15)
	stack.Pop()
	stack.Peek()
}

func (s *Stack) isEmpty() bool {
	if len(s.Items) == 0 {
		fmt.Println("Stack is empty")
		return true
	} else {
		fmt.Println("Stack is not empty")
		return false
	}
}

func (s *Stack) Push(val int) {
	s.Items = append(s.Items, val)
	s.Top++
}

func (s *Stack) Pop() {
	if !s.isEmpty() {
		fmt.Println("popped from stack:", s.Items[s.Top])
		s.Items = s.Items[:s.Top]
		s.Top--
	}
}

func (s *Stack) Peek() {
	if !s.isEmpty() {
		fmt.Println("Top element is: ", s.Items[s.Top])
	}
}
