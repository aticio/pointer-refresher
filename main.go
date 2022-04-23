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
	jim := person{
		firstName: "Özgür",
		lastName:  "Atıcı",
		contactInfo: contactInfo{
			email:   "aticiozgur@gmail.com",
			zipCode: 34743,
		},
	}

	jim.updateName("Ozu")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
