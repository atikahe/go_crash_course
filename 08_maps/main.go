package main

import "fmt"

func main() {
	// Define map
	// map[key_type]value_type
	// emails := make(map[string]string)

	// Assign key value
	// emails["Soobin"] = "soobin@txt.com"
	// emails["Arin"] = "arin@omg.com"

	// Shorthand
	emails := map[string]string{"Soobin": "soobin@txt.com", "Arin": "arin@omg.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Soobin"])

	// Delete from map
	delete(emails, "Soobin")
	fmt.Println(emails)
}