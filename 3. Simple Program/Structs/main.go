package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string

	// this part is embedded struct
	// which contains same as contactInfo
	contact contactInfo
}

func main() {
	var rob person
	fmt.Printf("%+v\n", rob)
	rob.firstName = "Rob"
	rob.lastName = "Pike"
	fmt.Printf("%+v\n", rob)

	adith := person{
		firstName: "Adithya Visnu",
		lastName:  "Prasetyo Putra",
	}
	fmt.Println(adith)

	bob := person{"Uncle", "Bob", contactInfo{}}
	fmt.Println(bob)

	reid := person{"Uncle", "Reid", contactInfo{"reidididi@gmail.com", 40555}}
	reid.updateFirstName("Shawn")
	reid.print()
	// why those code above is not replacing the first name?
	// see the explanation below on updateFirstName function

	// what that ampersand and stars symbol means?
	// & or ampersand: is to get an address from memory of particular value
	//                 some people called it a pointer for an address
	// * or stars: is to get an value from an particular address
	reidPointer := &reid
	fmt.Println(reidPointer)
	reidPointer.updateLastName("Hannah")
	reid.print()
}

// what is that receiver function has a stars symbol meant?!!?!
// it means like hey give me a pointer/an address to the value of the type person
// stars on function receiver and inside a body of function would be different things
func (pointerToReid *person) updateLastName(newLastName string) {
	fmt.Println("should be ", newLastName)
	// then we can use the receiver function parameter
	// into the value with stars symbol
	// or hey give me a value of that pointer/address
	(*pointerToReid).lastName = newLastName
}

// when golang use this function receiver
// it will copy the entire struct to the new memory address
// so the p on this function is different with the reid variable above
func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

// print formats function is for printing a struct recursively with its values
func (p person) Print() {
	fmt.Printf("%+v\n", p)
}
