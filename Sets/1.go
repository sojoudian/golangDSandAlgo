package main

import (
	"errors"
	"fmt"
)

//The most important characteristics of Sets is
//that every element of the set *MUST* be uniqe
//and it is NOT ordered!
//it has tw important functions, add and get

// Set - our reperesentation of a set data structure
type Set struct {
	//  Element is map of strings to a empty struct
	Elements map[string]struct{}
}

//we have to have a method to instantiate this map

func NewSet() *Set {
	var set Set
	set.Elements = make(map[string]struct{})
	return &set
}

// Add - adds an element to the set
func (s *Set) Add(element string) {
	s.Elements[element] = struct{}{}
}

// Delete - deletes an element from the set, if it exists
func (s *Set) Delete(element string) error {
	if _, exist := s.Elements[element]; !exist {
		return errors.New("Element not present in set")
	}
	delete(s.Elements, element)
	return nil
}

// Contains - check to see if the element exists within the set
func (s *Set) Contains(element string) bool {
	_, exist := s.Elements[element]
	return exist
}

// List - Print out all of the elements within the set
func (s *Set) List() {
	for k, _ := range s.Elements {
		fmt.Println(k)
	}

}

func main() {
	fmt.Println("############## Sets ##############\n")

	mySet := NewSet()

	mySet.Add("Earth")
	mySet.Add("Venus")
	mySet.Add("Mars")
	mySet.List()

	// if we add Earch again, you can see in the
	//Output that we will only have one Earth element
	//This is the uniqueness that we talked about
	fmt.Println("############## #### ##############\n")
	mySet.Add("Earth")
	mySet.Delete("Venus")
	mySet.List()

	fmt.Println(mySet.Contains("Mars"))
}
