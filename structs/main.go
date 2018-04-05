package main

import "fmt"

// Stucts are sort of like objects in javascript
type contactInfo struct {
	email string
	zip   int
}

// Setting variable name as the same name as type
// e.g. contactInfo set as contactInfo type
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// Order matters in this way of declaring a struct
	alex := person{"Al", "Pu", contactInfo{"al@al.com", 12222}}
	// A better way to declare a person struct
	bobContact := contactInfo{email: "bobisgreat@hotmail.com", zip: 12209}
	bob := person{firstName: "Bob",
		lastName:    "Anderson",
		contactInfo: bobContact}

	// Created this way, sets the fields to zero values
	var chuck person
	// Setting is just like the javascript. I recognize this
	chuck.firstName = "Chuck"
	chuck.lastName = "Finley"
	chuck.contactInfo.email = "hero@yahoo.com"
	chuck.contactInfo.zip = 22225

	alex.updateName("Alex")
	alex.print()
	bob.print()
	chuck.print()
}

// reciever is a pointer to a person, that is what the * refers to
func (p *person) updateName(newFirstName string) {
	// address into value do *address
	// value into addres do &value
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
