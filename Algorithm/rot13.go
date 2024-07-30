package main

import (
	"fmt"
	"strings"
)

func rot13(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+13)%26
	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+13)%26
	}
	return r
}

func main() {

	fmt.Println(strings.Map(rot13, "Hello, my name is Raphael. I love writing algorithms..."))

}
