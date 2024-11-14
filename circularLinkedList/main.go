package main

import "fmt"

type circularLinkList struct {
	data int
	next *circularLinkList
}

func main() {
	head := &circularLinkList{5, nil}
	head.next = head
	head.display()
	fmt.Println("After insertion at end")
	head = head.addNodeAtEnd(10)
	head = head.addNodeAtEnd(20)
	head.display()
	fmt.Println("After insertion at beginning")
	head = head.addNodeAtFirst(0)
	head.display()
	fmt.Println("After insertion at particular position")
	head = head.addNodeAtpos(15, 3)
	head.display()
	fmt.Println("After deleting a first node")
	head = head.deleteFirstNode()
	head.display()
	fmt.Println("After deleting a last node")
	head = head.deleteLastNode()
	head.display()
	fmt.Println("After deleting a particular node")
	head = head.DeleteNodeByIndex(1)
	head.display()
	head = head.addNodeAtpos(10, 1)
	// head.Traverse() //infinite loop
}

func (c *circularLinkList) display() {
	head := c
	for {
		if c == nil {
			break
		} else if c.next == head {
			fmt.Println(c.data)
			break
		}
		fmt.Print(c.data, " ")
		c = c.next
	}
}

func (c *circularLinkList) addNodeAtEnd(val int) *circularLinkList {
	newNode := &circularLinkList{val, nil}
	if c == nil {
		newNode.next = newNode
		return newNode
	}
	head := c
	for {
		if c.next != c {
			c = c.next
		}
		c.next = newNode
		newNode.next = head
		break
	}
	return head
}

func (c *circularLinkList) addNodeAtFirst(val int) *circularLinkList {
	newNode := &circularLinkList{val, nil}
	if c == nil {
		newNode.next = newNode
		return newNode
	}
	newNode.next = c
	head := c
	for c.next != head {
		c = c.next
	}
	c.next = newNode
	return newNode
}

func (c *circularLinkList) addNodeAtpos(val, pos int) *circularLinkList {
	newNode := &circularLinkList{val, nil}
	if c == nil {
		return newNode
	}
	if pos == 0 {
		head := c.addNodeAtFirst(val)
		return head
	}
	head := c
	for i := 0; i < pos-1; i++ {
		c = c.next
	}
	newNode.next = c.next
	c.next = newNode
	return head
}

func (c *circularLinkList) deleteFirstNode() *circularLinkList {

	if c == nil {
		return c
	}

	oldHead := c
	newHead := c.next

	for c.next != oldHead {
		c = c.next
	}
	c.next = newHead
	return newHead
}

func (c *circularLinkList) deleteLastNode() *circularLinkList {

	if c != nil {
		prevNode := c
		head := c
		for c.next != head {
			prevNode = c
			c = c.next
		}
		prevNode.next = head
		return head
	}
	return c
}

func (head *circularLinkList) DeleteNodeByValue(value int) *circularLinkList {
	if head == nil {
		fmt.Println("List is empty")
		return head
	}

	if head.data == value {
		head = head.next
		return head
	}

	p := head
	prevNode := &circularLinkList{}
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

func (head *circularLinkList) DeleteNodeByIndex(index int) *circularLinkList {
	if head == nil {
		fmt.Println("List is empty")
		return head
	}

	if index == 0 {
		return head.deleteFirstNode()
	}

	p := head
	prevNode := &circularLinkList{}
	for i := 1; i <= index; i++ {
		prevNode = p
		p = p.next
		if p.next == head && i < index {
			fmt.Println("out of index")
			return head
		}
	}
	prevNode.next = p.next
	return head

}

func (c *circularLinkList) Traverse() {
	for {
		if c == nil {
			break
		}
		fmt.Println(c.data)
		c = c.next
	}
}
