package main

import "fmt"

func main() {
	fmt.Println("Array")
	// var nameOfArray [Capacity]type
	var myArrayInt [8]int
	fmt.Println("my arrayInt is: ", myArrayInt)

	var myArrayStr [8]string
	fmt.Println("my arrayString is: ", myArrayStr)

	var myArraybyte [8]byte
	fmt.Println("my arraybyte is: ", myArraybyte)

	myArrayIntSec := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("my arrayIntSec is: ", myArrayIntSec)
	fmt.Println("lent of myArrayIntSec is: ", len(myArrayIntSec))
	fmt.Println("First element of myArrayIntSec is: ", myArrayIntSec[0])
	fmt.Println("Last element of myArrayIntSec is: ", myArrayIntSec[7])
	fmt.Println("Last element of myArrayIntSec with second method is: ", myArrayIntSec[len(myArrayIntSec)-1])
	for i, n := range myArrayIntSec {
		fmt.Printf("Index: %d\t, Value: %d\n", i, n)
	}

}
