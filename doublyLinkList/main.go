package main

import "fmt"

type doublyLinkList struct {
	data int
	prev *doublyLinkList
	next *doublyLinkList
}

func main() {
	head := &doublyLinkList{data: 15, prev: nil, next: nil}

	fmt.Println("After adding at first")
	head = head.addNodeAtFirstPos(10)
	head = head.addNodeAtFirstPos(5)
	head.traverseForward()

	fmt.Println("After adding at last")
	head = head.addNodeAtLastPos(30)
	head = head.addNodeAtLastPos(35)
	head.traverseForward()

	fmt.Println("After adding at mid")
	head = head.addNodeAtMidPos(20, 3)
	head = head.addNodeAtMidPos(25, 4)
	head.traverseForward()

	fmt.Println("After deleting first node")
	head = head.deleteFirstNode()
	head.traverseForward()

	fmt.Println("After deleting last node")
	head = head.deleteLastNode()
	head.traverseForward()

	fmt.Println("After deleting mid node")
	head = head.deleteMidNode(2) // expecting valid index number
	head = head.deleteMidNode(3)
	head.traverseForward()

	head = head.deleteMidNode(2)
	head.traverseForward()
	fmt.Println("After deleting mid node last")
	head = head.deleteMidNode(1)
	head.traverseForward()

	fmt.Println("After deleting mid node last1")
	head = head.deleteMidNode(1)
	head.traverseForward()
	head = head.addNodeAtFirstPos(5)
	head.traverseForward()
	head = head.addNodeAtLastPos(10)
	head.traverseForward()
	head = head.deleteMidNode(1)
	head.traverseForward()
	head = head.deleteMidNode(2)
}

func (dll *doublyLinkList) addNodeAtFirstPos(val int) *doublyLinkList {

	newNode := &doublyLinkList{data: val, prev: nil, next: nil}
	if dll == nil {
		return newNode
	}
	dll.prev = newNode
	newNode.next = dll
	return newNode
}

func (dll *doublyLinkList) addNodeAtLastPos(val int) *doublyLinkList {

	newNode := &doublyLinkList{data: val, prev: nil, next: nil}
	head := dll
	for dll.next != nil {
		dll = dll.next
	}
	dll.next = newNode
	newNode.next = nil
	newNode.prev = dll
	return head
}

func (dll *doublyLinkList) addNodeAtMidPos(val, pos int) *doublyLinkList {

	newNode := &doublyLinkList{data: val, prev: nil, next: nil}
	if dll.next == nil {
		fmt.Println("list is empty!")
		return nil
	}
	head := dll

	for range pos - 1 {
		dll = dll.next
	}
	fmt.Println("data: ", dll.data)
	newNode.next = dll.next
	dll.next.prev = newNode
	dll.next = newNode
	newNode.prev = dll
	return head
}

func (dll *doublyLinkList) deleteFirstNode() *doublyLinkList {
	if dll == nil {
		fmt.Println("list is empty!")
		return nil
	}
	if dll.next != nil {
		dll = dll.next
		dll.prev = nil
	} else {
		return nil
	}
	return dll
}

func (dll *doublyLinkList) deleteLastNode() *doublyLinkList {
	if dll == nil {
		fmt.Println("list is empty!")
		return nil
	}

	if dll.next == nil {
		return nil
	}
	head := dll
	for dll.next != nil {
		dll = dll.next
	}
	dll = dll.prev
	dll.next = nil
	return head
}

func (dll *doublyLinkList) deleteMidNode(pos int) *doublyLinkList {

	if dll == nil {
		fmt.Println("empty list!")
		return nil
	}
	head := dll
	if pos == 1 {
		return head.deleteFirstNode()

	}

	if dll.next != nil {
		i := 1
		for range pos - 1 {
			dll = dll.next
			if dll == nil {
				fmt.Println("not a valid index")
				return head
			}
			if dll.next == nil && i == pos-1 {
				return head.deleteLastNode()
			} else if dll.next == nil && i > pos {
				fmt.Println("not a valid index")
				return head
			}
			i++
		}
		dll.next.prev = dll.prev
		dll.prev.next = dll.next
	} else if dll.next == nil && pos > 1 {
		fmt.Println("not a valid index")
	}
	return head
}

func (dll *doublyLinkList) traverseForward() {
	if dll == nil {
		fmt.Println("list is empty!")
		return
	}
	fmt.Println("Printing forward(head -> tail)")
	for {
		fmt.Print(dll.data, " ")
		if dll.next == nil {
			break
		}
		dll = dll.next
	}
	fmt.Println()
	dll.traverseBackward()
}

func (dll *doublyLinkList) traverseBackward() {
	fmt.Println("Printing backward(tail -> head)")
	if dll == nil {
		fmt.Println("list is empty!")
		return
	} else if dll.prev == nil {
		fmt.Println("It is an first node or single node with data", dll.data)
		return
	}
	for {
		fmt.Print(dll.data, " ")
		if dll.prev == nil {
			break
		}
		dll = dll.prev
	}
	fmt.Println()
}
