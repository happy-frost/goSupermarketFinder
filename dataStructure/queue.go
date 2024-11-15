package dataStructure

import (
	"errors"
	"fmt"
)

type Node struct {
	item string
	next *Node
}

type queue struct {
	front *Node
	back  *Node
	size  int
}

func (p *queue) Enqueue(name string) error {
	newNode := &Node{
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

func (p *queue) Dequeue() (string, error) {
	var item string

	if p.front == nil {
		return "", errors.New("empty queue!")
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

func (p *queue) PrintAllNodes() error {
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

func (p *queue) isEmpty() bool {
	return p.size == 0
}
