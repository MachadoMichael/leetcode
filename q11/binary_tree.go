package main

import "fmt"

type node struct {
	Left  *node
	Value int
	Right *node
}

type tree struct {
	Root *node
}

func (t *tree) Add(v int) {
	t.Root = insertNode(t.Root, v)
}

func (t *tree) Remove(v int) {
	t.Root = removeNode(t.Root, v)
}

func insertNode(n *node, v int) *node {
	if n == nil {
		return &node{Value: v}
	} else if n.Value > v {
		n.Left = insertNode(n.Left, v)
	} else {
		n.Right = insertNode(n.Right, v)
	}
	return n
}

func removeNode(n *node, v int) *node {

	if n.Value == v {
		if n.Left == nil && n.Right == nil {
			return nil
		} else if n.Left == nil && n.Right != nil {
			return n.Right
		} else if n.Right == nil && n.Left != nil {
			return n.Left
		} else {
			n.Right.Left = n.Left
			return n.Right
		}
	} else if n.Value > v {
		n.Left = removeNode(n.Left, v)
	} else {
		n.Right = removeNode(n.Right, v)
	}

	return n
}

func (t *tree) Print() {
	printLevelsAsLists(t.Root)
}

func printLevelsAsLists(root *node) {
	if root == nil {
		return
	}

	// A queue to hold nodes at each level
	queue := []*node{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var level []int // slice to store node values at the current level

		// Process all nodes at the current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Add the current node value to the level slice
			level = append(level, node.Value)

			// Enqueue the left child if it exists
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			// Enqueue the right child if it exists
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// Print the current level as a list
		fmt.Println(level)
	}
}

func main() {
	tree := &tree{Root: &node{Value: 10}}
	tree.Add(8)
	tree.Add(12)
	tree.Add(9)
	tree.Add(5)
	tree.Add(4)
	tree.Print()
	removeNode(tree.Root, 5)
	println("------------")
	tree.Print()
}
