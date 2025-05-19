package main

import "fmt"

func learnArrays() {
	a := [4]int{1, 2, 3, 4}

	fmt.Println(a)

	b := [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c := [...]int{1, 3: 5, 4}
	fmt.Println(c)

}
