package bst

import "fmt"

type (
	BinarySearchTree interface {
		Insert(value uint32)
		Find(value uint32) bool
	}

	Node struct {
		left  *Node
		value uint32
		right *Node
	}

	BST struct {
		tree *Node
	}
)

func (bst *BST) Insert(value uint32) {
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

func (bst *BST) find(value uint32) *Node {
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

func (bst *BST) Find(value uint32) bool {
	return bst.find(value).value != value
}

func (bst *BST) String() string {
	result := ""

	dft := func(node *Node) {}

	dft = func(node *Node) {
		if node.left != nil {
			dft(node.left)

		}
		result += fmt.Sprintf("%d ", node.value)
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
