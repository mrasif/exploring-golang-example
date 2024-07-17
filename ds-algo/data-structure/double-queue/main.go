package main

import "fmt"

type Node struct {
	Element int
	Next    *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

func (q *Queue) insert_front(element int) {
	newNode := Node{
		Element: element,
	}
	if q.Head == nil {
		q.Head = &newNode
		q.Tail = q.Head
		return
	}

	head := q.Head
	q.Head = &newNode
	q.Head.Next = head
}

func (q *Queue) insert_back(element int) {
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

func (q *Queue) delete_front() int {
	if q.Head == nil {
		return 0
	}
	removable := q.Head
	q.Head = q.Head.Next
	return removable.Element
}

func (q *Queue) delete_back() int {
	if q.Head == nil {
		return 0
	}
	current := q.Head
	previous := current
	for current.Next != nil {
		previous = current
		current = current.Next
	}
	if current == previous {
		q.Head = nil
		q.Tail = nil
		return current.Element
	}

	previous.Next = nil
	return current.Element
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
	fmt.Println("Double Queue")
	var queue Queue

	fmt.Println("Insert front-")
	queue.Print()
	queue.insert_front(1)
	queue.Print()
	queue.insert_front(2)
	queue.Print()
	queue.insert_front(3)
	queue.Print()

	fmt.Println("Insert back-")
	queue.insert_back(4)
	queue.Print()
	queue.insert_back(5)
	queue.Print()
	queue.insert_back(6)
	queue.Print()

	fmt.Println("Delete back-")
	queue.delete_back()
	queue.Print()
	queue.delete_back()
	queue.Print()
	queue.delete_back()
	queue.Print()

	fmt.Println("Delete front-")
	queue.delete_front()
	queue.Print()
	queue.delete_front()
	queue.Print()
	queue.delete_front()
	queue.Print()

}
