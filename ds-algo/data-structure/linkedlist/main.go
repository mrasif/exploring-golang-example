package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Insert(data int) {
	newNode := Node{
		Data: data,
		Next: nil,
	}
	// 1. If Head is null, create node as HEAD
	if l.Head == nil {
		l.Head = &newNode
		return
	}
	// 2. If Head is not null, traverse to last Node
	currentNode := l.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	// 3. Set the node to the next of the last Node
	currentNode.Next = &newNode
}

func (l *LinkedList) Print() {
	text := ""
	// 1. If Head is empty then print "[]"
	// 2. Print untill nextNode is nil
	if l.Head != nil {
		currentNode := *l.Head
		for currentNode.Next != nil {
			text = text + fmt.Sprintf(`%d,`, currentNode.Data)
			currentNode = *currentNode.Next
		}
		text = text + fmt.Sprintf(`%d`, currentNode.Data)
	}
	fmt.Println(fmt.Sprintf(`[%v]`, text))
}

func (l *LinkedList) Reverse() {
	// 1. Initial head as nil node
	// 2. Traverse through next node of linked list, and if next is present, set current node as next of next node.
	// 3. Next node will be current node.

	if l.Head != nil {
		current := l.Head
		var previous *Node
		for current != nil {
			next := current.Next
			current.Next = previous
			previous = current
			current = next
		}
		l.Head = previous
	}
}

func main() {
	fmt.Println("Linked List")
	var linkedList LinkedList

	linkedList.Print()
	linkedList.Insert(5)
	linkedList.Insert(20)
	linkedList.Insert(30)
	linkedList.Insert(25)
	linkedList.Print()
	linkedList.Reverse()
	linkedList.Print()
}
