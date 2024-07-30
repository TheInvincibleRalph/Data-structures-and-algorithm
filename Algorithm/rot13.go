package main

import (
	"fmt"
	"strings"
)

func rot13(r rune) rune {

	// The rot13 algorithm is a simple letter substitution cipher that replaces a letter with the 13th letter after it in the alphabet.
	switch {
	//checks if the rune is a letter and if it is uppercase or lowercase
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+13)%26
	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+13)%26
	}
	return r
}

func main() {
	//Map returns a copy of the string s with all its characters modified according to the mapping function
	//rot13 is the mapping function in this case
	fmt.Println(strings.Map(rot13, "Hello, my name is Raphael. I love writing algorithms..."))

}
