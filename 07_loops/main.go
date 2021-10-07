package main

import "fmt"

func main() {
	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}
}