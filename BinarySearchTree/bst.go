package BinarySearchTree

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// Insert add a node to the tree
// the key to be added should not already be in the tree
func (n *Node) Insert(key int) {
	if key < n.Data {
		// Move Right
		if n.Right == nil {
			n.Right = &Node{Data: key}
		} else {
			n.Right.Insert(key)
		}
	} else if key > n.Data {
		// Move Left
		if n.Left == nil {
			n.Left = &Node{Data: key}
		} else {
			n.Left.Insert(key)
		}
	} else {
		n.Left.Insert(key)
	}
}

// Search
// Returns True if the given key exist in the tree
func (n *Node) Search(key int) bool {
	if key < n.Data {
		n.Left.Search(key)
	} else if key > n.Data {
		n.Right.Search(key)
	}
}

func BST() {
	tree := &Node{Data: 100}
	tree.Insert(200)
	tree.Insert(300)
	fmt.Println(tree)
}
