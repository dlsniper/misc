package main

import (
	"fmt"

	"github.com/dlsniper/misc/sort/quick"
)

func main() {
	array := []int{12, 2, 31, -4, 1, 11, 17}

	array = quick.QuickSort(array)

	fmt.Printf("%v\n", array)
}
