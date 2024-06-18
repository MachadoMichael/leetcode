package main

import "fmt"

type LinkedList struct {
	Head   *LinkedItem
	Lenght uint
}

type LinkedItem struct {
	Node int
	Next *LinkedItem
}

func (l *LinkedList) append(node *LinkedItem) *LinkedList {
	if l.Head.Next == nil {
		l.Head.Next = node
		l.Lenght++
		return l
	} else {
		return l.append(node)
	}
}

func (l *LinkedList) prepend(node *LinkedItem) *LinkedList {
	node.Next = l.Head
	l.Head = node
	return l
}

func (l *LinkedList) insert(node *LinkedItem, index int) *LinkedList {
	i := 0
	var selectedNode *LinkedItem
	var beforeNode *LinkedItem
	for i < index {
		if i == index-1 {
			beforeNode = selectedNode
		}
		selectedNode = findNode(l.Head)
	}

	beforeNode.Next = node
	node.Next = selectedNode
	l.Lenght++
	return l
}

func (l *LinkedList) getNode(index int) *LinkedItem {
	var selectedNode *LinkedItem
	i := 0
	for i < index-1 {
		selectedNode = findNode(l.Head)
		i++
	}

	return selectedNode
}

func findNode(node *LinkedItem) *LinkedItem {
	return node.Next
}

func main() {
	fmt.Println("vim-go")

	linked := LinkedList{
		Head:   nil,
		Lenght: 0,
	}

	tail2 := LinkedItem{
		Node: 12,
		Next: nil,
	}

	tail3 := LinkedItem{
		Node: 99,
		Next: nil,
	}

	fmt.Println(linked.Head)
	linked.prepend(&tail2)
	linked.prepend(&tail2)
	linked.prepend(&tail2)
	linked.prepend(&tail2)
	linked.prepend(&tail2)
	linked.insert(&tail3, 2)
	fmt.Println(linked.getNode(2))
}
