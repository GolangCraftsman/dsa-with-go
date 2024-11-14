package main

import "fmt"

type Queue struct {
	front int
	rear  int
	data  []int
	size  int
}

func (q *Queue) init(s int) {
	q.data = make([]int, s, s)
	q.front = -1
	q.rear = -1
	q.size = s
}

func main() {
	queue := &Queue{}
	queue.init(5)

	queue.enQueue(5)
	queue.printQueue()
	queue.deQueue()
	queue.printQueue()
	queue.enQueue(10)
	queue.printQueue()
	queue.enQueue(15)
	queue.enQueue(20)
	queue.enQueue(25)
	queue.printQueue()
	queue.enQueue(30)

	queue.deQueue()
	queue.printQueue()

	queue.deQueue()
	queue.printQueue()

	queue.deQueue()
	queue.printQueue()

	queue.deQueue()
	queue.printQueue()

	queue.deQueue()
}

func (q *Queue) isQueueFull() bool {
	if q.rear == q.size-1 {
		fmt.Println("queue is full")
		return true
	} else {
		return false
	}
}

func (q *Queue) isQueueEmpty() bool {
	if q.rear == q.front {
		fmt.Println("queue is empty")
		return true
	} else {
		return false
	}
}

func (q *Queue) enQueue(val int) {
	if q.isQueueFull() {
		return
	} else {
		q.rear++
		q.data[q.rear] = val
	}
}

func (q *Queue) deQueue() {
	if q.isQueueEmpty() {
		return
	} else {
		q.front++
	}
}

func (q *Queue) printQueue() {
	if q.isQueueEmpty() {
		return
	}
	fmt.Printf("F: %d, R: %d\n", q.front, q.rear)

	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Print(q.data[i], " ")
	}
	fmt.Println()
}
