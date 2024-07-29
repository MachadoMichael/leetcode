package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Tree struct {
	Content  MainContent
	Children []*Node
}

type MainContent struct {
	Base64   string
	Subtitle string
}

type Node struct {
	Value    string
	Children []*Node
}

func NewNode(value string) *Node {
	return &Node{
		Value:    value,
		Children: []*Node{},
	}
}

func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

func generateRandomComment(r *rand.Rand) string {
	comments := []string{
		"Great post!",
		"Love this!",
		"Amazing!",
		"Nice photo!",
		"Wow, awesome!",
		"Beautiful!",
		"Fantastic!",
		"Very cool!",
		"Superb!",
		"Impressive!",
		"Love it!",
		"Awesome!",
		"Nice shot!",
		"Cool!",
		"Well done!",
		"Incredible!",
		"Wow!",
		"Stunning!",
		"Bravo!",
		"Outstanding!",
		"Excellent!",
		"Terrific!",
		"Wonderful!",
		"Marvelous!",
		"Spectacular!",
		"Remarkable!",
		"Gorgeous!",
		"Splendid!",
		"Breathtaking!",
		"Stellar!",
		"Astonishing!",
		"Magnificent!",
		"Super!",
		"Phenomenal!",
		"Awesome job!",
		"Kudos!",
		"Superb work!",
		"Well crafted!",
		"Artistic!",
		"Talented!",
		"Admirable!",
		"Inspired!",
		"Great capture!",
		"Impeccable!",
		"Outstanding work!",
		"Terrific job!",
		"Charming!",
		"Exceptional!",
		"Delightful!",
		"Praiseworthy!",
		"Ravishing!",
		"Sensational!",
		"Exquisite!",
		"Flawless!",
		"Wondrous!",
		"Majestic!",
		"Enchanting!",
		"Bewitching!",
		"Dazzling!",
		"Lovely!",
		"Fabulous!",
		"Brilliant!",
		"Radiant!",
		"Ethereal!",
		"Mesmerizing!",
		"Mesmerizing!",
		"Euphoric!",
		"Graceful!",
		"Poetic!",
		"Harmonious!",
		"Soothing!",
		"Soulful!",
		"Hypnotic!",
		"Vivid!",
		"Enthralling!",
		"Dreamy!",
		"Sublime!",
		"Immaculate!",
		"Spellbinding!",
		"Profound!",
		"Elegant!",
		"Resplendent!",
		"Uplifting!",
		"Divine!",
		"Stately!",
		"Polished!",
	}
	return comments[r.Intn(len(comments))]
}

func generateTreeWithRandomNodes(r *rand.Rand, nodeCount int) *Tree {
	root := NewNode(generateRandomComment(r))
	nodes := []*Node{root}

	for i := 1; i < nodeCount; i++ {
		newNode := NewNode(generateRandomComment(r))
		parentNode := nodes[r.Intn(len(nodes))]
		parentNode.AddChild(newNode)
		nodes = append(nodes, newNode)
	}

	photo := MainContent{
		Base64:   "",
		Subtitle: "",
	}

	return &Tree{
		Content:  photo,
		Children: []*Node{root},
	}
}

func printTreeByLevels(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			fmt.Printf("%s ", node.Value)
			queue = append(queue, node.Children...)
		}
		fmt.Println()
	}
}

func (t *Tree) printTreeStructured(node *Node, level int) {
	if node == nil {
		return
	}

	indent := strings.Repeat("  ", level)
	fmt.Printf("%s- %s\n", indent, node.Value)

	for _, child := range node.Children {
		t.printTreeStructured(child, level+1)
	}
}

func (t *Tree) searchComment(node *Node, comment string, level int) bool {

	start := time.Now()
	if node == nil {
		return false
	}

	if node.Value == comment {
		elapsed := time.Since(start)
		fmt.Printf("Found comment: %s at level %d\n and searchTime: %v", node.Value, level, elapsed)
		return true
	}

	for _, child := range node.Children {
		if t.searchComment(child, comment, level+1) {
			return true
		}
	}

	return false
}

func (t *Tree) SaveToJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(t)
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	tree := generateTreeWithRandomNodes(r, 50)

	tree.printTreeStructured(tree.Children[0], 0)

	if err := tree.SaveToJSON("./tree.json"); err != nil {
		fmt.Printf("Error saving tree to JSON: %v\n", err)
	}

	b := tree.searchComment(tree.Children[0], "Amazing!", 0)
	if b != true {
		fmt.Println("Comment not find.")
	}
}
