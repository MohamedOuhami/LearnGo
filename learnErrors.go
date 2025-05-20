package main

import (
	"errors"
	"fmt"
)

func foo(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}

	return arg + 1, nil
}

func learnError() {
	for _, i := range []int{7, 42} {

		if r, e := foo(i); e == nil {

			fmt.Printf("The number %s is cool with result %s ", i, r)
		} else {
			fmt.Println("There was an error")
		}
	}
}
