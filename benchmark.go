package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sort"
	"time"

	"github.com/dlsniper/misc/sort/histogram"
	"github.com/dlsniper/misc/sort/merge"
	"github.com/dlsniper/misc/sort/quick"
)

type Array []int

func (array Array) Len() int {
	return len(array)
}

func (array Array) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

func (array Array) Less(i, j int) bool {
	return array[i] < array[j]
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	array := []int{}

	start := time.Now()
	rand.Seed(time.Now().UTC().UnixNano())
	elems := 10000000
	for i := 0; i < elems; i++ {
		array = append(array, rand.Intn(math.MaxInt16))
	}

	fmt.Printf("Generated %d elements in %s\n", elems, time.Since(start))

	start = time.Now()
	array1, array2, array3, array4 := array, array, array, array
	fmt.Printf("Copied arrays in %s\n", time.Since(start))

	first := make(chan string)

	bms := func(ch chan string) {
		_ = histogram.HistogramSort(array1)
		ch <- "histogram"
	}

	ms := func(ch chan string) {
		_ = merge.MergeSort(array2)
		ch <- "merge"
	}

	qs := func(ch chan string) {
		_ = quick.QuickSort(array3)
		ch <- "quick"
	}

	intern := func(ch chan string) {
		sort.Sort(Array(array4))
		ch <- "internal"
	}

	//histogram.Init()
	_, _, _, _ = bms, ms, qs, intern
	start = time.Now()
	go ms(first)
	go intern(first)
	go qs(first)
	go bms(first)

	alg := <-first
	elapsed := time.Since(start)
	fmt.Printf("First to finish: %v in %s\n", alg, elapsed)
}
