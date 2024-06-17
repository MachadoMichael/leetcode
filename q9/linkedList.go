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

	fmt.Println(linked.Head)
	linked.prepend(&tail2)
	fmt.Println(linked.Head)
}
