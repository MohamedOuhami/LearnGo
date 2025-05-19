package main

import "fmt"

func learnMaps() {
	// Maps is the equivalent of Dicts in other languages, so It's a key-value pairing system
	// We create maps using the make() function

	m := make(map[string]int)

	m["key1"] = 1
	m["key2"] = 2

	fmt.Println(m)

	// We can delete a member of the map using its key
	delete(m, "key1")

	fmt.Println(m)

	// We can even clear the map completely using the clear function
	clear(m)

	fmt.Println(m)

	// When getting the value from a map, you also get whether it is present or not in the second return value
	_, prs := m["key2"]

	fmt.Println(prs)

	// And finally we can create a map in the same line
	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m2)

}
