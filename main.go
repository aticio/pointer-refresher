package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	ozu := person{
		firstName: "Özgür",
		lastName:  "Atıcı",
		contactInfo: contactInfo{
			email:   "aticiozgur@gmail.com",
			zipCode: 34743,
		},
	}

	// Go is a pass by value language. When you pass this struct to a function, Go copys that struct.
	// The struct definition and passed value will be kept seperately in the memory.

	// ozu: Actual value of struct
	// &ozu: Memory address of ozu struct
	// *ozuPointer: Value of what is being kept in that pointer address
	//ozuPointer := &ozu
	//ozuPointer.updateName("Ozu")

	// Go provides a shortcut for pointer usage with just passing the object itself as the pointer address
	ozu.updateName("Ozu")
	ozu.print()
}

// Notice how receiver changed
// When you see a * in front of a type it means that we are working with a pointer (to a person)
func (pointerToPerson *person) updateName(newFirstName string) {
	// Here the * is used as an operator to get the value of that pointer (to person)
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// Turn address into a value = *address
// Turn value into an address = &value
// Remember this structure changes when using receiver functions

//
// Reference Types vs Value Types--------------------

// Referance Types: Don't worry about pointers with these
// slices, maps, channels, pointers, functions

// Value Types: Use pointers to change these tings in a function
// int, float, string, bool, structs
