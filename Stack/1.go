package main

import (
	"errors"
	"fmt"
)

// Stack - our stack reperesentation
type Stack struct {
	Elements []int
}

// Push - The method which handles pushing new elements
// on top of our stack
func (s *Stack) Push(element int) {
	s.Elements = append(s.Elements, element)
}

//Pop - remove the top element from the stack
// or return an error if the stack is empty
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	} else {
		lastElementIndex := len(s.Elements) - 1
		var lastElement int
		lastElement, s.Elements = s.Elements[lastElementIndex], s.Elements[:lastElementIndex]
		return lastElement, nil
	}
}

//Peek - returns the top element of the stack without updating the stack
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	} else {
		return s.Elements[len(s.Elements)-1], nil
	}
}

// Length - will return the size of the stack
func (s *Stack) Length() int {
	return len(s.Elements)
}

// IsEmpty - will return true if the stack is 0
func (s *Stack) IsEmpty() bool {
	return len(s.Elements) == 0
}

func main() {
	fmt.Println("Stack Data Structure")
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	//stack.Push(4)

	peekOne, _ := stack.Peek()
	fmt.Println(peekOne)

	fmt.Println(stack.Length())

	elemOne, _ := stack.Pop()
	fmt.Println(elemOne)
	elemTwo, _ := stack.Pop()
	fmt.Println(elemTwo)
	elemThree, _ := stack.Pop()
	fmt.Println(elemThree)

	fmt.Println(stack.IsEmpty())
}
