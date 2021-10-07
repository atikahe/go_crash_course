package main

import "fmt"

func main() {
	x := 10
	y := 15

	if x <= y {
		fmt.Printf("%d is less or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}

	color := "orange"

	if color == "orange" {
		fmt.Println("It's a new black")
	} else if color == "black" {
		fmt.Println("It's an old orange")
	} else {
		fmt.Println("It's neither black nor orange")
	}

	// SWITCH
	switch color {
	case "orange":
		fmt.Println("It's a new black")
	case "black":
		fmt.Println("It's an old orange")
	default:
		fmt.Println("It's neither black nor orange")
	}
}