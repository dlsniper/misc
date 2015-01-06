package main

import (
	"fmt"

	"github.com/dlsniper/misc/hashtables"
)

func main() {
	demo := hashtables.NewHashtable(10)
	demo.Insert(1, "demo1")
	demo.Insert(2, "demo2")
	demo.Insert(3, "demo3")
	demo.Insert(4, "demo4")
	demo.Insert(5, "demo5")
	demo.Insert(6, "demo6")
	demo.Insert(12, "demo7")

	fmt.Printf("%s\n", demo)

	for _, key := range []int{6, 10} {
		k, v := demo.Find(key)
		fmt.Printf("%d\t-> %d\t%q\n", key, k, v)
	}
}
