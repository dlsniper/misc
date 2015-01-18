package main

import (
	"fmt"

	"github.com/dlsniper/misc/sort/merge"
)

func main() {
	array := []int{12, 2, 31, -4, 1, 11, 17}

	array = merge.MergeSort(array)

	fmt.Printf("%v\n", array)
}
