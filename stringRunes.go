package main

import (
	"fmt"
	"unicode/utf8"
)

func learnStringsRunes() {

	// A GO string is a read-only slice of bytes.
	// In GO, the concept of a character is called a "rune" = An integer that represents a Unicode code point

	const s = "สวัสดี"

	fmt.Println(s)

	// Since string is the equivalent of []byte, this will return the number of bytes in the slice
	fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()

	// Here, we get the count of 6, so basically Go sees 6 characters even though we only got 4, because some characters can be encoded on multiple bytes
	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	runes := []rune(s)

	for i, r := range runes {
		fmt.Printf("Rune %d: %c \n", i, r)
	}

}
