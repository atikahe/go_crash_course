package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName, lastName, city, gender string
	age int
}

// Value receiver (doesnt change data), like C void
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// Pointer receiver, changes data
func (p *Person) birthdayPassed() {
	p.age++
}

func (p * Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person 
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person2 := Person{firstName: "Bob", lastName: "Johnson", city: "New York", gender: "m", age: 30}

	fmt.Println(person1)
	person1.birthdayPassed()
	fmt.Println(person1.age)
	person1.getMarried("Williams")
	fmt.Println(person1.greet())

	fmt.Println(person2)
	person2.getMarried("Smith")
	fmt.Println(person2.greet())	
}