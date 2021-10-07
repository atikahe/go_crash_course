package main

import (
	"fmt"
)

func main() {
	// Array declaration
	// var ghostArr [2]string

	// Assign value
	// ghostArr[0] = "Jelangkung"
	// ghostArr[1] = "Poci"

	// Declare and initialize
	ghostArr := []string{"Jelangkung", "Poci", "Jerapah", "Ketan"}

	fmt.Println(ghostArr)
	// Slice from idx 1, before idx 3
	fmt.Println(ghostArr[1:3])
}