package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Medan", "Sumatera Utara", "Indonesia"}
	address2 := address1 // copy value (pass by value)
	address2.City = "Tebing Tinggi"

	fmt.Println(address1) // address 1 tidak berubah, karena defaultnya pass by value
	fmt.Println(address2)

	// POINTER -> pass by reference
	address3 := &address1
	//var address3 *Address = &address1
	address3.City = "Siantar"
	fmt.Println(address1)
	fmt.Println(address3)
}
