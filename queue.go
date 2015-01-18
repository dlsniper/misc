package main

import (
	"fmt"

	"github.com/dlsniper/misc/queue"
)

func main() {
	myQueue := queue.NewQueue()
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)
	myQueue.Enqueue("4")

	fmt.Printf("Queue size: %d\n", myQueue.Size())

	elem, err := myQueue.Dequeue()
	fmt.Printf("Dequeued elem: %#v\n", elem)
	fmt.Printf("Queue size: %d\n", myQueue.Size())

	elem, err = myQueue.Dequeue()
	fmt.Printf("Queue size: %d\n", myQueue.Size())

	elem, err = myQueue.Dequeue()
	fmt.Printf("Queue size: %d\n", myQueue.Size())

	elem, err = myQueue.Dequeue()
	fmt.Printf("Dequeued elem: %#v\n", elem.(string))
	fmt.Printf("Queue size: %d\n", myQueue.Size())

	elem, err = myQueue.Dequeue()
	fmt.Printf("Dequeued elem: %#v\nError: %q\n", elem, err)
	fmt.Printf("Queue size: %d\n", myQueue.Size())
}
