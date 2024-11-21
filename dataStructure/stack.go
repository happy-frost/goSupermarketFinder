package dataStructure

import (
	"errors"
	"fmt"
)

// type Node struct {
// 	item string
// 	next *Node
// }

type Stack[T comparable] struct {
	top  *node[T]
	size int
}

func NewStack[T comparable]() Stack[T] {
	return Stack[T]{nil, 0}
}

func (p *Stack[T]) Push(name T) error {
	newNode := &node[T]{
		item: name,
		next: nil,
	}
	if p.top == nil {
		p.top = newNode
	} else {
		newNode.next = p.top
		p.top = newNode
	}
	p.size++
	return nil
}

func (p *Stack[T]) Pop() (T, error) {
	var item T

	if p.top == nil {
		return item, errors.New("empty Stack")
	}

	item = p.top.item
	if p.size == 1 {
		p.top = nil
	} else {
		p.top = p.top.next
	}
	p.size--
	return item, nil
}

func (p *Stack[T]) PrintAllNodes() error {
	currentNode := p.top
	if currentNode == nil {
		fmt.Println("Stack is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}

	return nil
}

func (s Stack[T]) IsEmpty() bool {
	return s.size == 0
}
