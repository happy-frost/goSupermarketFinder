package dataStructure

import (
	"errors"
	"fmt"
)

// type stringNode struct {
// 	item string
// 	next *stringNode
// }

type Queue[T comparable] struct {
	front *node[T]
	back  *node[T]
	size  int
}

func NewQueue[T comparable]() Queue[T] {
	return Queue[T]{nil, nil, 0}
}

func (p *Queue[T]) Enqueue(name T) error {
	newNode := &node[T]{
		item: name,
		next: nil,
	}
	if p.front == nil {
		p.front = newNode

	} else {
		p.back.next = newNode

	}
	p.back = newNode
	p.size++
	return nil
}

func (p *Queue[T]) Dequeue() (T, error) {
	var item T

	if p.front == nil {
		return item, errors.New("empty queue")
	}

	item = p.front.item
	if p.size == 1 {
		p.front = nil
		p.back = nil
	} else {
		p.front = p.front.next
	}
	p.size--
	return item, nil
}

func (p *Queue[T]) PrintAllNodes() error {
	currentNode := p.front
	if currentNode == nil {
		fmt.Println("Queue is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}

	return nil
}

func (p *Queue[T]) IsEmpty() bool {
	return p.size == 0
}
