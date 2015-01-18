package bst

import (
	"fmt"

	"github.com/dlsniper/misc/queue"
)

type (
	BinarySearchTree interface {
		Insert(value int)
		Find(value int) bool
	}

	Node struct {
		left  *Node
		value int
		right *Node
	}

	BST struct {
		tree *Node
	}
)

func (bst *BST) Insert(value int) {
	if bst.tree == nil {
		bst.tree = &Node{value: value}
		return
	}

	node := bst.find(value)

	if node.value > value {
		node.left = &Node{value: value}
	} else if node.value < value {
		node.right = &Node{value: value}
	}
}

func (bst *BST) find(value int) *Node {
	node := bst.tree
	parentNode := node
	for node != nil {
		parentNode = node
		if node.value > value {
			node = node.left
		} else if node.value < value {
			node = node.right
		} else {
			return node
		}
	}

	return parentNode
}

func (bst *BST) Find(value int) bool {
	return bst.find(value).value != value
}

func (bst *BST) BFT() string {
	result := ""

	levelNodes := queue.NewQueue()

	levelNodes.Enqueue(bst.tree)
	for levelNodes.Size() > 0 {
		elem, err := levelNodes.Dequeue()
		if err != nil {
			break
		}
		node := elem.(*Node)

		if node.left != nil {
			levelNodes.Enqueue(node.left)
		}

		if node.right != nil {
			levelNodes.Enqueue(node.right)
		}

		result += fmt.Sprintf("%v ", node.value)
	}

	return result
}

func (bst *BST) String() string {
	result := ""

	dft := func(node *Node) {}

	dft = func(node *Node) {
		if node.left != nil {
			dft(node.left)

		}
		result += fmt.Sprintf("%v ", node.value)
		if node.right != nil {
			dft(node.right)
		}
	}
	dft(bst.tree)

	return result
}

func NewBinarySearchTree() *BST {
	var bst BST

	return &bst
}
