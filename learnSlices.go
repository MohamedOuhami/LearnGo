package main

import "fmt"

func learnSlices() {
	// Slices are an even more important way of representing data than slices

	// Here, we define an uninitialized slice. It is equal to nil, and has a length of 0
	var s []string
	fmt.Println("uninit: ", s, s == nil, len(s) == 0)

	// To create a slice of non-zero length, we use the make function
	sl := make([]string, 3)
	fmt.Println("emp: ", sl, "len :", len(sl), "cap: ", cap(sl))

	// Now, we can set the values in the slices and get them
	sl[0] = "a"
	sl[1] = "b"
	sl[2] = "c"

	fmt.Println("Set : ", sl)
	fmt.Println("get :", sl[1])

	// Now, In addition to the basic operations, we can even append to the slices presented using the append func
	sl = append(sl, "d")
	sl = append(sl, "e")

	fmt.Println("Set : ", sl)

	// Slices also can be copied using the copy method
	c := make([]string, len(sl))
	copy(c, sl)

	fmt.Println("Set of copy : ", sl)

	// We also have the option to slice through the values of the slice
	sl1 := sl[2:]
	sl2 := sl[1:4]
	sl3 := sl[:5]

	fmt.Println("Set of slice : ", sl1)
	fmt.Println("Set of slice : ", sl2)
	fmt.Println("Set of slice : ", sl3)

}
