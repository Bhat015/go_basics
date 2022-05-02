package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	person1 := person{
		firstName: "James",
		lastName:  "Bond",
	}
	fmt.Println(person1.firstName)

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Printf("%s %s\n", alex.firstName, alex.lastName)

	fmt.Printf("%+v", alex) // {firstName:"Alex" lastName:"Anderson"}

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@email.com",
			zipCode: 94000,
		},
	}

	fmt.Println(jim.contact.email)

	jimPointer := &jim // jimPointer is a pointer to jim
	jimPointer.updateName("Jimmy")
	jim.print()

	//shortcut
	//go will allow to call receiver pointer or pointer variable
	jim.updateName("Jimmy")

	jim.print()

}

func (pointerPerson *person) updateName(newFirstName string) {
	(*pointerPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
