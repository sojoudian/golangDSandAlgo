package main

import (
	"errors"
	"fmt"
)

// Queue - Our reperesentation of a queue data structure
type Queue struct {
	Elements []int
}

//Enqueue - add an element of type int to the end of the queue
func (q *Queue) Enqueue(element int) {
	q.Elements = append(q.Elements, element)
}

//Dequeue - returns the first element of the queue
func (q *Queue) Dequeue() (int, error) {
	// return the first element
	// update the element slice, to remove the first element from the slice
	//error checking and validate queue is not empty
	if len(q.Elements) == 0 {
		return 0, errors.New("queue is empty")
	}
	var firstElement int
	firstElement, q.Elements = q.Elements[0], q.Elements[1:]
	return firstElement, nil
}

// Length - returns the length of the queue
func (q *Queue) Length() int {
	return len(q.Elements)
}

// IsEmpty - returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
	// if len(q.Elements) == 0 {
	// 	return true
	// }
	// return false
}

// Peek - returns the first element in the queue without updating the queue
func (q *Queue) Peeks() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return q.Elements[0], nil
}

func main() {
	fmt.Println("Queues Section")

	queue := Queue{}
	fmt.Println("check the queue after initiation", queue)
	queue.Enqueue(1)
	fmt.Println("check the queue after adding the first element", queue)
	queue.Enqueue(2)
	fmt.Println("check the queue after adding the second element", queue)

	qElement, _ := queue.Dequeue()
	fmt.Println("check the queue after applying Dequeue method", qElement)
	fmt.Println("check the queue status after applying Dequeue method", queue)

}
