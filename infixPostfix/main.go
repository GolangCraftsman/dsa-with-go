package main

import "fmt"

type Stack struct {
	Items []string
	Top   int
}

func main() {
	expression := "x-y/z-k*d"
	fmt.Println(getPostfix(expression))
}

func getPostfix(expression string) string {
	stack := &Stack{Top: -1}
	postfix := ""
	i := 0
	for i < len(expression) {
		char := expression[i]
		c := string(char)
		if isOperand(c) {
			postfix += c
			i++
		} else {
			if !stack.isEmpty() {
				if getPrecedence(c) > getPrecedence(stack.Items[stack.Top]) {
					stack.Push(c)
					i++
				} else {
					postfix += stack.Pop()
				}
			} else {
				stack.Push(c)
				i++
			}
		}
	}
	for !stack.isEmpty() {
		postfix += stack.Pop()
	}
	return postfix
}

func isOperand(char string) bool {
	if char == "/" || char == "*" || char == "+" || char == "-" {
		return false
	} else {
		return true
	}
}

func getPrecedence(c string) int {
	if c == "/" || c == "*" {
		return 2
	}
	if c == "+" || c == "-" {
		return 1
	}
	return 0
}

func (s *Stack) isEmpty() bool {
	if len(s.Items) == 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack) Push(val string) {
	s.Items = append(s.Items, val)
	s.Top++
}

func (s *Stack) Pop() string {
	num := s.Items[s.Top]
	s.Items = s.Items[:s.Top]
	s.Top--
	return num
}
