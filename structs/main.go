package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string

	// You don't have to specify the variable name if it already exists
	contactInfo
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// &jim is a reference to the place in memory (Pointer)
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")

	jim.updateName("Jimmy")
	jim.print()
}

// *person is a Type of Pointer that points at a person
// Whenever there is a *thing where a Type would be it is a type description
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointerToPerson is an operator that turns variable into the value of a pointer
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
