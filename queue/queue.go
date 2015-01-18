package queue

import "fmt"

type (
	Queue interface {
		Enqueue(elem interface{})
		Dequeue() (interface{}, error)
		Size() int
	}

	MyQueue struct {
		queue []interface{}
	}
)

func (queue *MyQueue) Enqueue(elem interface{}) {
	queue.queue = append(queue.queue, elem)
}

func (queue *MyQueue) Dequeue() (interface{}, error) {
	if queue.Size() == 0 {
		return nil, fmt.Errorf("no more elements to dequeue")
	}

	var elem []interface{}
	elem, queue.queue = queue.queue[:1], queue.queue[1:]
	return elem[0], nil
}

func (queue *MyQueue) Size() int {
	return len(queue.queue)
}

func NewQueue() Queue {
	result := &MyQueue{}
	result.queue = []interface{}{}

	return result
}
