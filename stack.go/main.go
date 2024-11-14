package main

import "fmt"

type Stack struct {
	size int
	top  int
	arr  []int
}

func (s *Stack) initStack(size int) {
	s.arr = make([]int, size, size)
	s.top = -1
	s.size = size
}

func main() {
	stack := &Stack{}
	stack.initStack(5)
	stack.PushAndPrint(5)
	stack.PushAndPrint(10)
	stack.PushAndPrint(15)
	stack.PushAndPrint(20)
	stack.PushAndPrint(25)
	stack.PushAndPrint(30)
	stack.Peek(5)
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.GetTop()
	stack.PushAndPrint(35)
}

func (s *Stack) Pop() {
	if s.top > -1 {
		fmt.Printf("Element %d is poped from stack.\n", s.arr[s.top])
		s.top--
		if s.top > -1 {
			fmt.Printf("Now top element is: %d\n", s.arr[s.top])
		} else {
			fmt.Println("Now stack is empty...")
		}
	} else {
		fmt.Println("Stack is empty...")
	}
}
func (s *Stack) Push(elem int) {

	if s.top < s.size-1 {
		s.top++
		(s.arr)[s.top] = elem
	} else {
		fmt.Println("Stack is full...")
	}
}
func (s *Stack) GetTop() {
	if s.top == -1 {
		fmt.Println("Stack is empty...")
	} else {
		fmt.Println("Top element is:", s.arr[s.top])
	}
}
func (s *Stack) Peek(pos int) {
	if s.top == -1 {
		fmt.Println("Stack is empty...")
	} else if pos > s.top+1 || pos == 0 {
		fmt.Println("Invalid stack position.")
	} else {
		fmt.Printf("Element in stack at position %d is: %d\n", pos, s.arr[s.top-pos+1])
	}
}

func (s *Stack) PushAndPrint(el int) {
	s.Push(el)
	s.GetTop()
}
