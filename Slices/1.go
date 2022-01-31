package main

import "fmt"

func main() {
	fmt.Println("Slices")
	//the firs difference between array and slice is,
	//in slice you don't have to specify the slice size
	mySlices := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("My slice is: ", mySlices)
}
