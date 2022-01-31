package main

import (
	"errors"
	"fmt"
)

// PriorityQueue - Our reperesentation of a queue data structure
type PriorityQueue struct {
	High []int
	Low  []int
}

//Enqueue - add an element of type int to the end of the queue
func (q *PriorityQueue) Enqueue(element int, highPriority bool) {
	if highPriority {
		q.High = append(q.High, element)
	} else {
		q.Low = append(q.Low, element)
	}
}

//Dequeue - returns the first element of the queue
func (q *PriorityQueue) Dequeue() (int, error) {
	//if the length of highPriority queue !=0
	//return the first element from highPriority queue
	// otherwise if the length of lowPriority queue !=0
	// reurn the first element from lowPriority queue

	if len(q.High) != 0 {
		var firstElement int
		firstElement, q.High = q.High[0], q.High[1:]
		return firstElement, nil
	}
	if len(q.Low) != 0 {
		var firstElement int
		firstElement, q.Low = q.Low[0], q.Low[1:]
		return firstElement, nil
	}
	return 0, errors.New("queue is empty")

}

// Length - returns the length of the queue
func (q *PriorityQueue) Length() int {
	return len(q.High) + len(q.Low)
}

// IsEmpty - returns true if the queue is empty
func (q *PriorityQueue) IsEmpty() bool {
	return q.Length() == 0
	// if len(q.Elements) == 0 {
	// 	return true
	// }
	// return false
}

// Peek - returns the first element in the queue without updating the queue
func (q *PriorityQueue) Peeks() (int, error) {
	if len(q.High) != 0 {
		return q.High[0], nil
	}
	if len(q.Low) != 0 {
		return q.Low[0], nil
	}
	return 0, errors.New("empty queue")
}

func main() {
	fmt.Println("Queues Section")

	queue := PriorityQueue{}
	fmt.Println("check the Priority queue after initiation", queue)
	queue.Enqueue(17, true)
	fmt.Println("check the Priority queue after adding the first element to the High Priority queue", queue)
	queue.Enqueue(10, false)
	fmt.Println("check the queue after adding the first element to the Low Priority queue", queue)
	fmt.Println("==========================================================================================")
	qElement, _ := queue.Dequeue()
	fmt.Println("check the queue after applying Dequeue method", qElement)
	fmt.Println("check the queue status after applying Dequeue method", queue)

	fmt.Println("==========================================================================================")
	queue.Enqueue(35, true)
	fmt.Println("check the Priority queue after adding another element to the High Priority queue", queue)
	qElement, _ = queue.Dequeue()
	fmt.Println("check the queue after applying Dequeue method", qElement)
	fmt.Println("check the queue status after applying Dequeue method", queue)
	qElement, _ = queue.Dequeue()
	fmt.Println("check the queue after applying Dequeue method", qElement)
	fmt.Println("check the queue status after applying Dequeue method", queue)

}
