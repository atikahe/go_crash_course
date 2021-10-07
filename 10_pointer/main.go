package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read val from pointer
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val from pointer
	*b = 10
	fmt.Println(a)
}