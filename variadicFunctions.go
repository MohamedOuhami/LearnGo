package main

import "fmt"

// The variadic functions can take an arbitrary number of arguments and still work out
func learnVariadic(nums ...int) {

	fmt.Println(nums, " ")

	total := 0

	for _, num := range nums {
		total += num

	}

	fmt.Println(total)

}
