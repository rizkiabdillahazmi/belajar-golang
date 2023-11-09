package main

import "fmt"

type HasName interface {
	GetName() string
}

// interface untuk param
func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// implementasi interface
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{"Rizki"}
	SayHello(person)

	animal := Animal{"Kucing"}
	SayHello(animal)
}
