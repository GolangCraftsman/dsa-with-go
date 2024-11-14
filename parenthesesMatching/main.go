package main

import "fmt"

type Stack struct {
	Items []string
	Top   int
}

func main() {
	expression := "{a*b(c+d)/e+[f-g]}"
	if isParenthesesMatching(expression) {
		fmt.Printf("\"%s\" is balanced\n", expression)
		return
	}
	fmt.Printf("\"%s\" is not balanced\n", expression)
}

func isParenthesesMatching(expression string) bool {
	stack := &Stack{Top: -1}
	for _, char := range expression {
		if string(char) == "{" || string(char) == "(" || string(char) == "[" {
			stack.Push(string(char))
		} else if string(char) == "}" || string(char) == ")" || string(char) == "]" {
			if !stack.isEmpty() {
				if !stack.Pop(string(char)) {
					return false
				}
			} else {
				return false
			}
		}
	}
	return stack.isEmpty()
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

func (s *Stack) Pop(c string) bool {
	if s.Top != -1 {
		if s.Items[s.Top] == "{" && c == "}" {
			fmt.Println("popped from stack:", s.Items[s.Top])
			s.Items = s.Items[:s.Top]
			s.Top--
			return true
		} else if s.Items[s.Top] == "(" && c == ")" {
			fmt.Println("popped from stack:", s.Items[s.Top])
			s.Items = s.Items[:s.Top]
			s.Top--
			return true
		} else if s.Items[s.Top] == "[" && c == "]" {
			fmt.Println("popped from stack:", s.Items[s.Top])
			s.Items = s.Items[:s.Top]
			s.Top--
			return true
		} else {
			return false
		}
	}
	return false
}
