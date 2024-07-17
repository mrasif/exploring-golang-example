package main

import "fmt"

type Node struct {
	Element int
	Next    *Node
}

type Stack struct {
	Size int
	Top  *Node
}

func (stack *Stack) Push(element int) {
	newNode := Node{
		Element: element,
	}
	if stack.Size == 0 {
		stack.Top = &newNode
		stack.Size = stack.Size + 1
		return
	}

	previous := stack.Top
	newNode.Next = previous
	stack.Top = &newNode
	stack.Size = stack.Size + 1
}

func (stack *Stack) Pop() int {
	if stack.Size == 0 {
		fmt.Println("Stack is empty.")
		return 0
	}
	stack.Size = stack.Size - 1
	element := stack.Top.Element
	stack.Top = stack.Top.Next
	return element
}

func (stack *Stack) Print() {
	fmt.Print("Elements of stack is [")
	top := stack.Top
	for top != nil {
		fmt.Print(top.Element, " ")
		top = top.Next
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("Stack")
	var stack Stack

	stack.Push(10)
	stack.Push(5)
	stack.Push(14)
	stack.Print()

	fmt.Println(stack.Size)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Size)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Size)
}
