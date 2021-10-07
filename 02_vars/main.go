package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Var declaration
	var age int64 = 24
	const isTall = true
	var gpa float32 = 2.3

	// Shorthand
	// name := "atikah"
	// email := "atk@gmail.com"

	name, email := "atikah", "atk@gmail.com"

	fmt.Println(name, age, isTall, gpa, email)

	// Print type
	fmt.Printf("%T\n", age)
}