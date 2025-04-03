package main

import "fmt"

//Struct - It is a data structure which is a collection of properties that are related together.

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// contactInfo // Another way of embedding struct
}

func main() {
	// alice := person{firstName: "Alice", lastName: "Anderson"}
	// fmt.Println(alice)

	var alice person
	alice.firstName = "Alice"
	fmt.Println(alice)
	fmt.Printf("%+v\n", alice)

	bob := person{
		firstName: "Bob",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "bob@gmail.com",
			zipCode: 93234,
		},
	}

	// bob.print()
	// bobPointer := &bob
	// bobPointer.updateName("James")
	// (*bobPointer).print()

	// or

	bob.updateName("James")
	bob.print()

}

func (p person) print() {
	fmt.Printf("Name: %v %v\n", p.firstName, p.lastName)
	fmt.Printf("Email: %v\n", p.contact.email)
	fmt.Printf("Zip: %v\n", p.contact.zipCode)
}

func (p *person) updateName(firstName string) {
	p.firstName = firstName
}
