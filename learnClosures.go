package main

import "fmt"

func learnClosures() {

	// Go supports anonymous functions which can form Closures
	// When we do intSeq, we set i to 0, and then assign the function to nextInt
	nextInt := intSeq()

	// Here, we're executing the anonymous function inside the intSeq while remembering the i value
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Here, we setuping another function with another counter
	newIntSeq := intSeq()

	fmt.Println(newIntSeq())
	fmt.Println(newIntSeq())
}

func intSeq() func() int {
	// Setup the counter
	i := 0
	// The function to return
	return func() int {
		// Incrementing by 1
		i++
		// Return the value
		return i
	}
}
