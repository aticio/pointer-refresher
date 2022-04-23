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
	ozu.updateName("Ozu")
	ozu.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
