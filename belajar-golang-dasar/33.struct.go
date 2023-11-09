package main

import "fmt"

// Biasanya nama attributnya pascal case
type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var rizki Customer
	rizki.Name = "Rizki Abdillah Azmi"
	rizki.Address = "Jakarta Pusat"
	rizki.Age = 23

	fmt.Println(rizki.Name)
	fmt.Println(rizki.Address)
	fmt.Println(rizki.Age)
	fmt.Println(rizki)

	nizam := Customer{
		Name:    "Nizam Salihin",
		Address: "Jakarta barat",
		Age:     23,
	}

	fmt.Println(nizam)

	joko := Customer{"Joko", "Bandung", 27}
	fmt.Println(joko)

	rizki.sayHello("Joko")
}
