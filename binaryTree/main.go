package main

import "fmt"

type BinaryTree struct {
	leftChild  *BinaryTree
	rightChild *BinaryTree
	data       int
}

func main() {
	/*
			5
		 1	   3
		0 2   7 9
	*/
	bt := &BinaryTree{}

	bt.createNode(5)
	bt.createLeftNode(1)
	bt.createRightNode(3)

	bt.leftChild.createLeftNode(0)
	bt.leftChild.createRightNode(2)

	bt.rightChild.createLeftNode(7)
	bt.rightChild.createRightNode(9)

	bt.preOrderTraversal()
	fmt.Println()
	bt.postOrderTraversal()
	fmt.Println()
	bt.inOrderTraversal()
	fmt.Println()
}

func (bt *BinaryTree) preOrderTraversal() {
	if bt != nil {
		fmt.Printf("%d ", bt.data)
		bt.leftChild.preOrderTraversal()
		bt.rightChild.preOrderTraversal()
	}
}

func (bt *BinaryTree) postOrderTraversal() {
	if bt != nil {
		bt.leftChild.postOrderTraversal()
		bt.rightChild.postOrderTraversal()
		fmt.Printf("%d ", bt.data)
	}
}

func (bt *BinaryTree) inOrderTraversal() {
	if bt != nil {
		bt.leftChild.inOrderTraversal()
		fmt.Printf("%d ", bt.data)
		bt.rightChild.inOrderTraversal()
	}
}

func (bt *BinaryTree) createNode(val int) {
	bt.data = val
	bt.leftChild = nil
	bt.rightChild = nil
}

func (bt *BinaryTree) createLeftNode(val int) {
	node := &BinaryTree{leftChild: nil, rightChild: nil, data: val}
	bt.leftChild = node
}

func (bt *BinaryTree) createRightNode(val int) {
	node := &BinaryTree{leftChild: nil, rightChild: nil, data: val}
	bt.rightChild = node
}
