package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		asciiValue := i
		character := rune(asciiValue)
		s := fmt.Sprintf("%c", character)
		fmt.Printf("%v\n",s)
	}
}
