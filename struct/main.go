package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "John",
		lastName:  "smith",
		contactInfo: contactInfo{
			email:   "john@gmail.com",
			zipCode: "458526",
		},
	}

	// jimPointer := &jim
	jim.updateName("Prashant")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
