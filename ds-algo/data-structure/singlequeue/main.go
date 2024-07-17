package main

import (
	"fmt"
)

type Node struct {
	Element int
	Next    *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

func (q *Queue) Enqueue(element int) {
	newNode := Node{
		Element: element,
	}
	if q.Tail == nil {
		q.Head = &newNode
		q.Tail = q.Head
		return
	}
	q.Tail.Next = &newNode
	q.Tail = &newNode
}

func (q *Queue) Dequeue() int {
	if q.Head == nil {
		return 0
	}
	removable := q.Head
	q.Head = q.Head.Next
	return removable.Element
}

func (q *Queue) Print() {
	fmt.Print("[")
	current := q.Head
	for current != nil {
		fmt.Print(current.Element, " ")
		current = current.Next
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("Single Queue")
	var queue Queue

	queue.Enqueue(3)
	queue.Enqueue(23)
	queue.Enqueue(2)
	queue.Print()
	fmt.Println(queue.Dequeue())
	queue.Print()

}
