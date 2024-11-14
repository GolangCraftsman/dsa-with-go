package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func main() {
	head := &Node{1, nil}
	fmt.Println("before Addition")
	head.Traverse()
	head = head.AddNodeAtFirst(2)
	head = head.AddNodeInBetween(2, 2)
	head = head.AddNodeAtLast(3)
	head = head.AddNodeAtLast(4)
	fmt.Println("After Addition")
	head.Traverse()
	head = head.DeleteNodeByValue(1)
	head = head.DeleteNodeByValue(2)
	fmt.Println("After Deletion")
	head.Traverse()
	head = head.AddNodeAtFirst(1)
	head = head.AddNodeAtLast(5)
	// head = head.DeleteNodeByValue(6)
	head = head.DeleteNodeByIndex(2)
	fmt.Println("After Updating")
	head.Traverse()
}

func (head *Node) Traverse() {

	for {
		fmt.Println(head.data)
		if head.next == nil {
			break
		}
		head = head.next
	}
}

func (head *Node) AddNodeAtFirst(value int) *Node {
	newNode := &Node{value, nil}
	newNode.next = head
	return newNode
}

func (head *Node) AddNodeInBetween(value, index int) *Node {
	newNode := &Node{value, nil}

	p := head

	for i := 0; i < index-1; i++ {
		p = p.next
	}

	newNode.next = p.next
	p.next = newNode

	return head
}

func (head *Node) AddNodeAtLast(value int) *Node {
	newNode := &Node{value, nil}

	if head == nil {
		head = newNode
		return head
	}

	p := head

	for p.next != nil {
		p = p.next
	}
	p.next = newNode
	return head
}

func (head *Node) DeleteNodeByValue(value int) *Node {
	if head == nil {
		fmt.Println("List is empty")
		return head
	}

	if head.data == value {
		head = head.next
		return head
	}

	p := head
	prevNode := &Node{}
	for p.next != nil {
		if p.data == value {
			prevNode.next = p.next
			return head
		}
		prevNode = p
		p = p.next
	}
	if p.data == value {
		prevNode.next = nil
	}
	return head

}

func (head *Node) DeleteNodeByIndex(index int) *Node {
	if head == nil {
		fmt.Println("List is empty")
		return head
	}

	if index == 0 {
		head = head.next
		return head
	}

	p := head
	prevNode := &Node{}
	for i := 0; i < index; i++ {
		prevNode = p
		p = p.next
		if p.next == nil {
			return head
		}
	}
	if p.next != nil {
		prevNode.next = p.next
	} else {
		prevNode.next = nil
	}
	return head

}
