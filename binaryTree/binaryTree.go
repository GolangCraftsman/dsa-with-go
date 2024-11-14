package main

import "fmt"

type BinaryTree struct {
	leftChild  *BinaryTree
	rightChild *BinaryTree
	data       int
}

func main() {
	//In a binary search tree (BST), a node's predecessor is the largest node that is smaller than the given node, and its successor is the smallest node that is larger than the given node. These nodes can be used to perform operations such as deleting a node or finding the minimum or maximum value in the tree.
	/*
			10
		 5	   15
		2 7   13 17
	*/
	bt := &BinaryTree{}

	bt.createNode(10)
	bt.insertNode(5)
	bt.insertNode(15)

	bt.insertNode(2)
	bt.insertNode(7)

	bt.insertNode(13)
	bt.insertNode(1)

	bt.preOrderTraversal()
	fmt.Println()
	bt.postOrderTraversal()
	fmt.Println()
	bt.inOrderTraversal()
	fmt.Println()

	fmt.Println("5:", bt.searchNode(5))
	bt.deleteNode(15)
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

func (bt *BinaryTree) insertNode(val int) {
	if bt == nil {
		bt.data = val
	}

	if val <= bt.data {
		if bt.leftChild == nil {
			node := &BinaryTree{leftChild: nil, rightChild: nil, data: val}
			bt.leftChild = node
			return
		} else {
			bt.leftChild.insertNode(val)
			return
		}
	} else {
		if bt.rightChild == nil {
			node := &BinaryTree{leftChild: nil, rightChild: nil, data: val}
			bt.rightChild = node
			return
		} else {
			bt.rightChild.insertNode(val)
			return
		}
	}

}

func (bt *BinaryTree) searchNode(val int) bool {
	if bt.data == val {
		return true
	}
	if val < bt.data {
		if bt.leftChild != nil {
			return bt.leftChild.searchNode(val)
		} else {
			return false
		}
	} else {
		if bt.rightChild != nil {
			return bt.rightChild.searchNode(val)
		} else {
			return false
		}
	}
}

func (bt *BinaryTree) getPredecessor(n int) *BinaryTree {
	bt = bt.leftChild
	for bt.rightChild != nil {
		bt = bt.rightChild
	}
	return bt
}

func (bt *BinaryTree) deleteNode(n int) *BinaryTree {

	if bt == nil {
		return nil
	}

	if bt.leftChild == nil && bt.rightChild == nil {
		bt = nil
		return nil
	}

	if n < bt.data {
		bt.leftChild = bt.leftChild.deleteNode(n)
	} else if n > bt.data {
		bt.rightChild = bt.rightChild.deleteNode(n)
	} else {
		ipre := bt.getPredecessor(n)
		bt.data = ipre.data
		bt.leftChild = bt.leftChild.deleteNode(ipre.data)
	}
	return bt
}
