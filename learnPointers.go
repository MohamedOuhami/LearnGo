package main

import "fmt"

func learnPointer() {

	// Pointers allows us to pass references to values and records
	// Here, we set up the value of i to 1
	i := 1
	fmt.Println(i)

	// When we do zeroVal(i), we only give a copy of the i to the value that normally should change It to 0, but when we print i again
	// It prints 1, because only the copy changed not the actual member
	zeroVal(i)
	fmt.Println("zeroval:", i)

	// But when we use the pointer to a value. We change the actual value by giving It the address to that box, and change its value
	zeroPtr(&i)
	fmt.Println("zeroptr", i)

	fmt.Println("pointer", &i)

}

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}
