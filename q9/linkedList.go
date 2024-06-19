package main

import (
	"log"
)

type LinkedList struct {
	Head   *LinkedItem
	Lenght int
}

type LinkedItem struct {
	Value int
	Next  *LinkedItem
}

func (l *LinkedList) append(node *LinkedItem) *LinkedList {
	var currentlyNode *LinkedItem = l.Head
	if currentlyNode == nil {
		l.Head = node
		l.Lenght++
		return l
	}

	for currentlyNode.Next != nil {
		currentlyNode = currentlyNode.Next
	}

	if currentlyNode.Next == nil {
		currentlyNode.Next = node
		l.Lenght++
	}

	return l
}

func (l *LinkedList) prepend(node *LinkedItem) *LinkedList {
	node.Next = l.Head
	l.Head = node
	l.Lenght++
	return l
}

func (l *LinkedList) insert(node *LinkedItem, index int) *LinkedList {
	if l.Lenght < index || index < 0 {
		log.Fatalln("index out range.")
	}

	if l.Lenght == 0 {
		l.Head = node
		l.Lenght++
		return l
	}

	i := 0
	var selectedNode *LinkedItem = l.Head
	var beforeNode *LinkedItem
	for i < index {
		if i == index-1 {
			beforeNode = selectedNode
		}
		selectedNode = selectedNode.Next
		i++
	}

	beforeNode.Next = node
	node.Next = selectedNode
	l.Lenght++
	return l
}

func printNodes(node *LinkedItem) {
	println(node.Value)
	if node.Next != nil {
		printNodes(node.Next)
	}
}

func main() {

	linked := LinkedList{
		Head:   nil,
		Lenght: 0,
	}

	tail2 := LinkedItem{
		Value: 12,
		Next:  nil,
	}

	tail3 := LinkedItem{
		Value: 99,
		Next:  nil,
	}

	tail4 := LinkedItem{
		Value: 55,
		Next:  nil,
	}

	tail5 := LinkedItem{
		Value: 77,
		Next:  nil,
	}

	linked.append(&tail3)
	linked.append(&tail4)
	linked.insert(&tail2, 1)
	linked.prepend(&tail5)
	printNodes(linked.Head)
	println("Lenght: ", linked.Lenght)
}
