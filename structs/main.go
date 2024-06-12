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
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var john person
	fmt.Println(john)
	fmt.Printf("%+v \n", john)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.print()
	// & operator = to access memory address of jim
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	jim.updateName("jimmy")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%v+ \n", p)
}

// reciver *person = this func can only be called by a (pointer to person)*person
// *pointerToPerson in func = operator * to access value the pointer is reference
func (pointerToPerson *person) updateName(newFirstName string) {
	// (*pointerToPerson) = value of pointerToPerson
	(*pointerToPerson).firstName = newFirstName
}
