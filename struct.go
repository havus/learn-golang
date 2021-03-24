package main

import "fmt"

type Person struct {
	FirstName, LastName, Address string
	Age int
}

// Struct Method
func (orang Person) sayHaiTo(name string) {
	fmt.Println("Holla", name, "my Name is", orang.FirstName, orang.LastName)
}
func (orang Person) GetFullName() string {
	return orang.FirstName + " " + orang.LastName
}

// Interface
type DataPerson interface {
	GetFullName() string
}

func sayGoodbye(data DataPerson) {
	fmt.Println("Goodbye", data.GetFullName())
}

func main() {
	// var john Person = Person{}
	john := Person{
		FirstName: 	"John",
		LastName:		"Doe",
		Address:		"Balikpapan",
		Age:				20,
	}

	fmt.Println(john.FirstName)

	john.FirstName = "Johnson"
	fmt.Println(john.FirstName)

	john.sayHaiTo("Orang Kota")

	sayGoodbye(john)
}
