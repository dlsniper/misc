package main

import (
	"fmt"

	"github.com/dlsniper/misc/sort/bitmap"
)

func main() {
	array := []int{12, 2, 31, 4, 1, 11, 17}

	array = bitmap.BitmapSort(array)

	fmt.Printf("%v\n", array)
}
