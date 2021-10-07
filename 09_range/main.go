package main

import "fmt"

func main() {
	ids := []int{33, 27, 98, 55, 72, 46, 88}

	// Loop
	for i, id := range ids {
		fmt.Printf("%d - ID %d\n", i, id)
	}

	// Withot index
	for _, id := range ids {
		fmt.Printf("ID %d\n", id)
	}

	// Sum
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range w map
	emails := map[string]string{"Soobin": "soobin@txt.com", "Arin": "arin@omg.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}
}