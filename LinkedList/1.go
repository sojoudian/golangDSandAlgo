package main

import "fmt"

type LinkedList struct {
	Head *Node
	Size int
}

//Node - a node repesents a link in our linked list
type Node struct {
	Value string
	Next  *Node
}

// Insert - adds a new element to the start of our Linked List
func (l *LinkedList) Insert(element string) {
	node := Node{
		Value: element,
		Next:  l.Head,
	}
	l.Head = &node
	l.Size++
}

//DeleteFirst - removes the first element from our Linked List.
func (l *LinkedList) DeleteFirst() {
	l.Head = l.Head.Next
	l.Size--
}

//List - iterates through all of the elements in our linked list
func (l *LinkedList) List() {
	current := l.Head
	for current != nil {
		fmt.Printf("%+v\n", current)
		current = current.Next
	}
}

// Search - searches through eve every element of our linked list
// returns the first element that matches the string otherwise
//it returns nil
func (l *LinkedList) Search(element string) *Node {
	current := l.Head
	for current != nil {
		if current.Value == element {
			return current
		}
		current = current.Next
	}
	return nil
}

// Delete - removes an element from our linked list if it matches
// the value
func (l *LinkedList) Delete(element string) {
	previous := l.Head
	current := l.Head
	for current != nil {
		if current.Value == element {
			previous.Next = current.Next
		}
		previous = current
		current = current.Next
	}
}
func main() {
	fmt.Println("Singly Linked List")

	var ll LinkedList
	ll.Insert("One")
	ll.Insert("Two")
	ll.Insert("Three")
	ll.Insert("Four")
	ll.Insert("Five")

	ll.Delete("Two")
	ll.List()
	fmt.Println("+++++++++++++++")

	//ll.List()

	// ll.DeleteFirst()
	// ll.List()

	if element := ll.Search("One"); element != nil {
		fmt.Printf("%+v\n", element)
	}
}
