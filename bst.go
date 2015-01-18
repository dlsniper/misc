package main

import (
	"fmt"

	"github.com/dlsniper/misc/bst"
)

func main() {
	bst := bst.NewBinarySearchTree()
	bst.Insert(10)
	bst.Insert(11)
	bst.Insert(9)
	bst.Insert(6)
	bst.Insert(8)
	bst.Insert(9)
	bst.Insert(10)
	bst.Insert(15)
	bst.Insert(12)
	fmt.Printf("Binary search tree DFT in-order: %s\n", bst)
	fmt.Printf("Binary search tree BFT in-order: %s\n", bst.BFT())
}
