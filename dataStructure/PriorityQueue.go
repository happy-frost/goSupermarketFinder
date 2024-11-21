package dataStructure

import (
	"errors"
)

type node[T comparable] struct {
	item     T
	priority float64
	next     *node[T]
}

// Use a binary heap to find priority (min in this case) item
type PriorityQueue[T comparable] struct {
	front *node[T]
	back  *node[T]
	size  int
}

func NewPriorityQueue[T comparable]() PriorityQueue[T] {
	return PriorityQueue[T]{nil, nil, 0}
}

func (p *PriorityQueue[T]) Enqueue(id T, val float64) error {
	newNode := &node[T]{
		item:     id,
		priority: val,
		next:     nil,
	}

	if p.front == nil {
		p.front = newNode
		p.back = newNode
	} else if p.size == 1 {
		if newNode.priority >= p.front.priority {
			p.front.next = newNode
			p.back = newNode
		} else {
			p.back = p.front
			newNode.next = p.back
			p.front = newNode
		}
	} else {
		current := p.front
		for current != nil {
			// If we are on the last node (there is no next)
			if current.next == nil {
				current.next = newNode
				p.back = newNode
				break
			} else if newNode.priority < current.next.priority {
				newNode.next = current.next
				current.next = newNode
				break
			}
			current = current.next
		}
	}
	p.size++
	return nil
}

func (p *PriorityQueue[T]) Dequeue() (T, error) {
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

// func (p *PriorityQueue) printAllNodes() error {
// 	currentNode := p.front
// 	if currentNode == nil {
// 		fmt.Println("Queue is empty.")
// 		return nil
// 	}
// 	fmt.Printf("%+v\n", currentNode.item)
// 	for currentNode.next != nil {
// 		currentNode = currentNode.next
// 		fmt.Printf("%+v\n", currentNode.item)
// 	}

// 	return nil
// }

func (p *PriorityQueue[T]) IsEmpty() bool {
	return p.size == 0
}

func (p *PriorityQueue[T]) Size() int {
	return p.size
}
